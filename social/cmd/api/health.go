package main

import (
	"encoding/json"
	"net/http"
)

// @Summary Get health status
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /v1/health [get]
func (a *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})

}