package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func validMethod(w http.ResponseWriter, r *http.Request, allowedMethod string) bool {
	if r.Method != allowedMethod {
		log.Printf("Rejected non-%s method: %s", allowedMethod, r.Method)
		http.Error(w, fmt.Sprintf("Only %s method is allowed", allowedMethod), http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	buffer := new(bytes.Buffer)

	err := json.NewEncoder(buffer).Encode(data)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(buffer.Bytes())
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func decodeBody[T any](w http.ResponseWriter, r *http.Request, target *T) bool {
	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return false
	}

	return true
}
