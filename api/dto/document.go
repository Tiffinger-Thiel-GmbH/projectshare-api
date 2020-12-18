package dto

// DocumentMetadata defines the metadata of a document.
type DocumentMetadata struct {
	ID       string `json:"id"`
	Location string `json:"location"`
	Name     string `json:"name"`
	MimeType string `json:"mimeType"`
}
