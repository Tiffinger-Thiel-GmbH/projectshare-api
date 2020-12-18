package aws

import (
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/constant"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

// Errors which can be returned by the DocumentRepository.
var (
	// ErrWhileUploadingFile Make sure you have the correct S3 credentials configured.
	ErrWhileUploadingFile = errors.New("error while uploading file")

	ErrWhileDownloadingFile = errors.New("error while downloading file")

	ErrWhileRetrievingFileHead = errors.New("error while retrieving file head object")
)

// DocumentRepository implements the DocumentRepository interface by using AWS S3 as backend.
// Instantiate it always using NewDocumentRepository.
type DocumentRepository struct {
	sess       *session.Session
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

// NewDocumentRepository creates a new connection to the AWS S3 server
// and returns a DocumentRepository which can be used to access it.
func NewDocumentRepository(region string) DocumentRepository {
	r := DocumentRepository{
		sess: createSession(region),
	}

	// Create an uploader with the session and default options
	r.uploader = s3manager.NewUploader(r.sess)

	// Create a downloader with the session and default options
	r.downloader = s3manager.NewDownloader(r.sess)

	return r
}

// getMetadata loads the metadata of a document using S3.
func (r *DocumentRepository) getMetadata(location string, documentID string) (model.Metadata, error) {
	headObject, err := s3.New(r.sess).HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(location),
		Key:    aws.String(documentID),
	})

	if err != nil {
		return model.Metadata{}, fmt.Errorf("%w of %v, %v:\n%v", ErrWhileRetrievingFileHead, location, documentID, err)
	}

	// Check if the values are really set and use default values if not.
	filename, ok := headObject.Metadata["Filename"]
	if !ok {
		defaultValue := "NO_FILENAME"
		filename = &defaultValue
	}

	mimeType, ok := headObject.Metadata["Mimetype"]
	if !ok {
		defaultValue := "text/plain"
		mimeType = &defaultValue
	}

	// Return the finished Metadata.
	return model.Metadata{
		ID:       uuid.MustParse(documentID),
		Location: location,
		Name:     *filename,
		Size:     *headObject.ContentLength,
		MimeType: *mimeType,
	}, nil
}

// GetDocumentsMetadata returns a list of all documents
// inside of the given location as an Metadata slice.
// Note that files with the name of constant.MetadataUUID get ignored as they are the metadata of a project.
func (r *DocumentRepository) GetDocumentsMetadata(location string) ([]model.Metadata, error) {
	// Get all objects.
	objects, err := s3.New(r.sess).ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(location),
	})

	if err != nil {
		return nil, fmt.Errorf("%w of %v:\n%v", ErrWhileRetrievingFileHead, location, err)
	}

	// Load the metadata for all objects and do this in a async way.

	// I use a mutex for synchronisation instead of channels to avoid the channel overhead
	// and it seems more simple at this place.
	// Note that go promotes channels very much, but they should not be overused.
	// https://medium.com/mindorks/https-medium-com-yashishdua-synchronizing-states-using-mutex-vs-channel-in-go-25e646c83567
	// Because of their overhead the benefits of them are not that much if the amount of work to do is not very much.

	// Here I could get a difference of ~100 ms with / vs without concurrency (using the mutex) for just 4 metadata requests.

	var result = make([]model.Metadata, 0)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// errs will contain all errors which happen while acquiring the metadata.
	errs := make(chan error, len(objects.Contents))

	// Register the amount of workers.
	wg.Add(len(objects.Contents))

	for _, object := range objects.Contents {
		key := *object.Key

		if key == constant.MetadataUUID {
			wg.Done()
			continue
		}

		go func() {
			// when finishing the goroutine, mark the worker as done.
			defer wg.Done()

			documentID := key
			metadata, err := r.getMetadata(location, documentID)
			if err != nil {
				errs <- fmt.Errorf("location: %v, documentId: %v\n%w", location, documentID, err)
				return
			}

			// Lock the mutex, write the data and unlock it.
			mu.Lock()
			result = append(result, metadata)
			mu.Unlock()
		}()
	}

	// Wait for finishing the work and close the error channel after that.
	wg.Wait()
	close(errs)

	// For now just check and return the first error.
	if err, open := <-errs; open && err != nil {
		return nil, err
	}

	return result, nil
}

// GetDocument returns a whole document by its location and id.
func (r *DocumentRepository) GetDocument(location string, documentID string) (model.Document, error) {
	// First get the metadata.
	metadata, err := r.getMetadata(location, documentID)
	if err != nil {
		return model.Document{}, err
	}

	// Create a buffer to write to.
	// TODO: Use a temp file? As it is not good to load everything into the ram.
	var bytes = make([]byte, metadata.Size)
	writer := aws.NewWriteAtBuffer(bytes)

	// Download the file.
	_, err = r.downloader.Download(writer,
		&s3.GetObjectInput{
			Bucket: aws.String(location),
			Key:    aws.String(documentID),
		})

	if err != nil {
		return model.Document{}, fmt.Errorf("%w:\n%v", ErrWhileDownloadingFile, err)
	}

	return model.Document{
		Metadata: metadata,
		Data:     bytes,
	}, err
}

// AddDocument saves a new file as document.
// It returns the new id and an error if the upload to S3 fails.
// You have to provide the full metadata of the file.
// The field metadata.Size has to match the size of the document to be read!
// The field metadata.ID gets overwritten by this id.
func (r *DocumentRepository) AddDocument(reader io.Reader, metadata model.Metadata) (uuid.UUID, error) {
	// Upload the file.
	key := uuid.New()
	_, err := r.uploader.Upload(&s3manager.UploadInput{
		Body:   reader,
		Bucket: aws.String(metadata.Location),
		Key:    aws.String(key.String()),
		Metadata: map[string]*string{
			"Filename": &metadata.Name,
			"Mimetype": &metadata.MimeType,
		},
	})

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%w:\n%v", ErrWhileUploadingFile, err)
	}

	return key, nil
}
