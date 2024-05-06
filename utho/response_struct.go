package utho

type BasicResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type CreateResponse struct {
	ID      string `json:"firewallid"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UpdateResponse struct {
	ID      string `json:"firewallid"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type DeleteResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
