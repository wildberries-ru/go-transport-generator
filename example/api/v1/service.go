package v1

// Detail ...
type Detail struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

// Namespace ...
type Namespace struct {
	Name string `json:"name"`
}

// GetDetailsEmbedStructResponse ...
type GetDetailsEmbedStructResponse struct {
	Detail    Detail    `json:"detail"`
	Namespace Namespace `json:"namespace"`
}
