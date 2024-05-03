package utho

type BasicResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type DeleteResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
