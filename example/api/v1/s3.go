package v1

// CreateMultipartUploadData ...
type CreateMultipartUploadData struct {
	Some     int    `json:"some"`
	UploadID string `json:"upload_id"`
}

// AdditionalErrors ...
type AdditionalErrors struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}
