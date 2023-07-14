package domain

type CreateServer struct {
	Name   string       `json:"name"`
	Type   ServerType   `json:"type"`
	Status ServerStatus `json:"status"`
}
