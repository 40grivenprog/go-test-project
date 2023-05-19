package interfaces

import (
	"encoding/json"
	"net/http"
)

func handleHttpError(w http.ResponseWriter, logger Logger, err error) {
	logger.LogError("%s", err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(err)
}

func handleHttpResponse(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
