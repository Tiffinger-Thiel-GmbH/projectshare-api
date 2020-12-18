package aws

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/constant"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

// Errors which can be returned by the ProjectRepository.
var (
	ErrWhileRetrievingBuckets      = errors.New("error while retrieving buckets")
	ErrCouldNotLoadBucketAsProject = errors.New("error while trying to load a bucket as a project")
	ErrInvalidUUID                 = errors.New("error while parsing bucket name as uuid")
	ErrMarshallingMetadata         = errors.New("error while marshalling the project metadata to json")
	ErrCreatingProject             = errors.New("could not create the project")
	ErrCreatingProjectMetadata     = errors.New("could not create the project metadata")
)

// ProjectRepository implements the ProjectRepository. This implementation uses the S3 buckets as project,
// the name is the projectId. Each bucket has to contain a file with the key constant.MetadataUUID
// which contains a simple json file with the metadata such as the project name.
type ProjectRepository struct {
	DocumentRepository
}
type bucketMetaData struct {
	Name string `json:"name"`
}

// GetProjects returns all existing projects.
func (r ProjectRepository) GetProjects() ([]model.Project, error) {
	// First get all buckets.
	buckets, err := s3.New(r.sess).ListBuckets(&s3.ListBucketsInput{})

	if err != nil {
		return nil, fmt.Errorf("%w:\n%v", ErrWhileRetrievingBuckets, err)
	}

	// Then load the metadata document for each of them
	// TODO: Concurrency, would be much faster?
	var projects []model.Project
	for _, bucket := range buckets.Buckets {
		// The metadata document has always the same key.
		metadataDocument, err := r.GetDocument(*bucket.Name, constant.MetadataUUID)
		if err != nil {
			// Ignore if the metadata document does not exist.
			// It seems that aws sometimes also returns deleted buckets, which fail here.
			// This is because it's a distributed service which needs time to refresh itself.
			continue
		}

		// Parse the json of the metadata.
		bucketMetadata := bucketMetaData{}
		err = json.Unmarshal(metadataDocument.Data, &bucketMetadata)
		if err != nil {
			return nil, fmt.Errorf("%w bucket: %v:\n%v", ErrCouldNotLoadBucketAsProject, *bucket.Name, err)
		}

		// Use the bucket name as ID.
		id, err := uuid.Parse(*bucket.Name)
		if err != nil {
			return nil, fmt.Errorf("%w bucket: %v:\n%v", ErrInvalidUUID, *bucket.Name, err)
		}

		// Add it to the projects.
		projects = append(projects, model.Project{
			ID:   id,
			Name: bucketMetadata.Name,
		})
	}

	return projects, nil
}

// AddProject creates a new project in S3. This implementation just uses buckets as
// projects as stated in aws.ProjectRepository documentation.
func (r ProjectRepository) AddProject(name string) (model.Project, error) {
	// Generate a new projectId and create a new bucket using it as name.
	projectID := uuid.New()
	_, err := s3.New(r.sess).CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(projectID.String()),
	})
	if err != nil {
		return model.Project{}, fmt.Errorf("%w:\n%v", ErrCreatingProject, err)
	}

	// Create the metadata json saved as a document in the bucket.
	metadata, err := json.Marshal(bucketMetaData{
		Name: name,
	})
	if err != nil {
		return model.Project{}, fmt.Errorf("%w:\n%v", ErrMarshallingMetadata, err)
	}

	_, err = r.uploader.Upload(&s3manager.UploadInput{
		Body:   bytes.NewBuffer(metadata),
		Bucket: aws.String(projectID.String()),
		Key:    aws.String(constant.MetadataUUID),
	})
	if err != nil {
		return model.Project{}, fmt.Errorf("%w:\n%v", ErrCreatingProjectMetadata, err)
	}

	// Return the final model.
	return model.Project{
		ID:   projectID,
		Name: name,
	}, nil
}
