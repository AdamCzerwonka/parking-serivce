package handlers

import (
	"encoding/json"
	"net/http"
)

func errorResponse(w http.ResponseWriter, errors []string, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	result := map[string]any{
		"code":   statusCode,
		"errors": errors,
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		return err
	}

	w.Write(jsonResult)
	return nil
}

func jsonResponse(w http.ResponseWriter, data any, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonResult, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Write(jsonResult)
	return nil
}
