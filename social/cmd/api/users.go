package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"social/update/internal/domain"
	"social/update/internal/middleware"
)

type registerUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, email, and password
// @Tags Users
// @Accept json
// @Produce json
// @Param payload body registerUserPayload true "Register User Payload"
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/users [post]
func (a *application) Register(w http.ResponseWriter, r *http.Request) {
	var payload registerUserPayload
	if err := a.readJSON(w, r, &payload); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	user, err := a.userService.Register(ctx, payload.Username, payload.Email, payload.Password)
	if err != nil {
		a.internalServerErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusCreated, user); err != nil {
		a.internalServerErrorResponse(w, r, err)
	}
}

// Get godoc
// @Summary Get a user by ID
// @Description Get user details by their database ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.User
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/users/{id} [get]
func (a *application) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		a.notFoundResponse(w, r)
		return
	}

	ctx := r.Context()
	user, err := a.userService.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			a.notFoundResponse(w, r)
			return
		}
		a.internalServerErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, user); err != nil {
		a.internalServerErrorResponse(w, r, err)
	}
}

type updateUserPayload struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

// Update godoc
// @Summary Update user profile
// @Description Update username or email for a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param payload body updateUserPayload true "Update User Payload"
// @Success 200 {object} domain.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/users/{id} [patch]
func (a *application) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		a.notFoundResponse(w, r)
		return
	}

	var payload updateUserPayload
	if err := a.readJSON(w, r, &payload); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	user, err := a.userService.Update(ctx, id, payload.Username, payload.Email)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			a.notFoundResponse(w, r)
			return
		}
		a.internalServerErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, user); err != nil {
		a.internalServerErrorResponse(w, r, err)
	}
}

// Delete godoc
// @Summary Soft-delete user profile
// @Description Marks user as deleted (soft delete)
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/users/{id} [delete]
func (a *application) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		a.notFoundResponse(w, r)
		return
	}

	ctx := r.Context()
	if err := a.userService.Delete(ctx, id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			a.notFoundResponse(w, r)
			return
		}
		a.internalServerErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// List godoc
// @Summary List all users with pagination
// @Description List registered users with support for pagination (limit, page)
// @Tags Users
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param page query int false "Page" default(1)
// @Success 200 {array} domain.User
// @Failure 500 {object} map[string]string
// @Router /v1/users [get]
func (a *application) List(w http.ResponseWriter, r *http.Request) {
	params := middleware.GetPagination(r)

	ctx := r.Context()
	users, err := a.userService.List(ctx, params.Limit, params.Offset)
	if err != nil {
		a.internalServerErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, users); err != nil {
		a.internalServerErrorResponse(w, r, err)
	}
}
