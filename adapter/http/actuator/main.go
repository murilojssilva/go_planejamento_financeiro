package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthyBody have Healthy vars
type HealthyBody struct {
	Status string `json:"status"`
}

// Healthy is a status function
func Healthy(w http.ResponseWriter, r *http.Request) { // visibilidade privada por começar com letra minúscula
	if r.Method != "GET" {
		w.WriteHeader(http.StatusHTTPVersionNotSupported)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var status = HealthyBody{"alive"}

	_ = json.NewEncoder(w).Encode(status)
}
