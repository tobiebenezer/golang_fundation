package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"social/update/internal/domain"
	"social/update/internal/middleware"
)

type postHandler struct {
	app *application
}

func (h *postHandler) Create(w http.ResponseWriter, r *http.Request) {
	h.app.CreatePost(w, r)
}

func (h *postHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.app.GetPost(w, r)
}

func (h *postHandler) Update(w http.ResponseWriter, r *http.Request) {
	h.app.UpdatePost(w, r)
}

func (h *postHandler) Delete(w http.ResponseWriter, r *http.Request) {
	h.app.DeletePost(w, r)
}

func (h *postHandler) List(w http.ResponseWriter, r *http.Request) {
	h.app.ListPosts(w, r)
}

type createPostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	UserID  int      `json:"user_id"`
	Tags    []string `json:"tags"`
}

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post with a title, content, user ID, and optional tags
// @Tags Posts
// @Accept json
// @Produce json
// @Param payload body createPostPayload true "Create Post Payload"
// @Success 201 {object} domain.Post
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/posts [post]
func (a *application) CreatePost(w http.ResponseWriter, r *http.Request) {
	var payload createPostPayload
	if err := a.readJSON(w, r, &payload); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	post := &domain.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  payload.UserID,
		Tags:    payload.Tags,
	}

	ctx := r.Context()
	if err := a.postService.Create(ctx, post); err != nil {
		a.internalServerErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusCreated, post); err != nil {
		a.internalServerErrorResponse(w, r, err)
	}
}

// GetPost godoc
// @Summary Get a post by ID
// @Description Get post details by their database ID
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} domain.Post
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/posts/{id} [get]
func (a *application) GetPost(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		a.notFoundResponse(w, r)
		return
	}

	ctx := r.Context()
	post, err := a.postService.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			a.notFoundResponse(w, r)
			return
		}
		a.internalServerErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, post); err != nil {
		a.internalServerErrorResponse(w, r, err)
	}
}

type updatePostPayload struct {
	Title   *string  `json:"title"`
	Content *string  `json:"content"`
	Tags    []string `json:"tags"`
}

// UpdatePost godoc
// @Summary Update post details
// @Description Update the title, content, or tags of a post
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param payload body updatePostPayload true "Update Post Payload"
// @Success 200 {object} domain.Post
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/posts/{id} [patch]
func (a *application) UpdatePost(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		a.notFoundResponse(w, r)
		return
	}

	var payload updatePostPayload
	if err := a.readJSON(w, r, &payload); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	post, err := a.postService.Update(ctx, id, payload.Title, payload.Content, payload.Tags)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			a.notFoundResponse(w, r)
			return
		}
		a.internalServerErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, post); err != nil {
		a.internalServerErrorResponse(w, r, err)
	}
}

// DeletePost godoc
// @Summary Soft-delete post
// @Description Marks post as deleted (soft delete)
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/posts/{id} [delete]
func (a *application) DeletePost(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		a.notFoundResponse(w, r)
		return
	}

	ctx := r.Context()
	if err := a.postService.Delete(ctx, id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			a.notFoundResponse(w, r)
			return
		}
		a.internalServerErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListPosts godoc
// @Summary List posts with pagination
// @Description List posts with support for pagination (limit, page)
// @Tags Posts
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param page query int false "Page" default(1)
// @Success 200 {array} domain.Post
// @Failure 500 {object} map[string]string
// @Router /v1/posts [get]
func (a *application) ListPosts(w http.ResponseWriter, r *http.Request) {
	params := middleware.GetPagination(r)

	ctx := r.Context()
	posts, err := a.postService.List(ctx, params.Limit, params.Offset)
	if err != nil {
		a.internalServerErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, posts); err != nil {
		a.internalServerErrorResponse(w, r, err)
	}
}
