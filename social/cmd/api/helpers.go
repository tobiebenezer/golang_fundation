package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (a *application) writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func (a *application) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	return dec.Decode(data)
}

func (a *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := map[string]any{"error": message}
	if err := a.writeJSON(w, status, env); err != nil {
		log.Printf("error writing json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (a *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	a.errorResponse(w, r, http.StatusNotFound, "the requested resource could not be found")
}

func (a *application) internalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %v", err)
	a.errorResponse(w, r, http.StatusInternalServerError, "the server encountered an error and could not complete your request")
}
