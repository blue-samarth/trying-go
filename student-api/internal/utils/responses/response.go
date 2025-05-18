package responses

import (
	"net/http"
	"encoding/json"
)

func JsonResponse(w http.ResponseWriter, status string, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(map[string]interface{}{
		"status":      status,
		"status_code": statusCode,
		"data":        data,
	})
}