package responses

// ServerResponse represents a server response
//
// swagger:model ServerResponse
type ServerResponse struct {
	// ID unique identifier of the server
	// example: 1
	ID uint `json:"id"`
	// Name of the server
	// example: ESGI
	Name string `json:"name"`
	// Type of the server
	// example: SMALL | MEDIUM | LARGE
	Type string `json:"type"`
	// Status of the server
	// example: STARTED | RUNNING | STOPPING | STOPPED
	Status string `json:"status"`
}
