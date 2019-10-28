package core

type HealthStatus struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
