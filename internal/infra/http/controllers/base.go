package controllers

import (
	"encoding/json"
	"net/http"
)

func success(w http.ResponseWriter, body interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(body)
}

func delete(w http.ResponseWriter, body interface{}) error {
	w.Header().Set("Content-Type", "text")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(body)
}

func created(w http.ResponseWriter, body interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	return json.NewEncoder(w).Encode(body)
}

func notFound(w http.ResponseWriter) error {
	(w).Header().Set("Content-Type", "text")
	(w).WriteHeader(http.StatusNotFound)

	return json.NewEncoder(w).Encode("wrong id")
}

func internalServerError(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	return json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}

