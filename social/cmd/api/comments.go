package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"social/update/internal/domain"
	"social/update/internal/middleware"
)

type CommentHandler struct {
	app *application
}

func NewCommentHandler(app *application) *CommentHandler {
	return &CommentHandler{app: app}
}

// CreateComment godoc
// @Summary Create a new comment
// @Description Create a new comment with a body and post ID
// @Tags Comments
// @Accept json
// @Produce json
// @Param payload body domain.Comment true "Create Comment Payload"
// @Success 201 {object} domain.Comment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/comments [post]
func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment domain.Comment
	if err := h.app.readJSON(w, r, &comment); err != nil {
		h.app.badRequestResponse(w, r, err)
		return
	}


	if err := h.app.commentService.CreateComment(r.Context(), &comment); err != nil {
		h.app.internalServerErrorResponse(w, r, err)
		return
	}

	h.app.writeJSON(w, http.StatusCreated, comment)
}


// GetCommentByID godoc
// @Summary Get a comment by ID
// @Description Get comment details by their database ID
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} domain.Comment
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/comments/{id} [get]
func (h *CommentHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.app.badRequestResponse(w, r, err)
		return
	}

	comment, err := h.app.commentService.GetCommentByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			h.app.notFoundResponse(w, r)
			return
		}
		h.app.internalServerErrorResponse(w, r, err)
		return
	}

	h.app.writeJSON(w, http.StatusOK, comment)
}


// UpdateComment godoc
// @Summary Update a comment
// @Description Update a comment with the given ID
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param payload body domain.Comment true "Update Comment Payload"
// @Success 200 {object} domain.Comment
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/comments/{id} [patch]
func (h *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.app.badRequestResponse(w, r, err)
		return
	}

	var comment domain.Comment
	if err := h.app.readJSON(w, r, &comment); err != nil {
		h.app.badRequestResponse(w, r, err)
		return
	}

	comment.ID = id

	updatedComment, err := h.app.commentService.UpdateComment(r.Context(), id, &comment)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			h.app.notFoundResponse(w, r)
			return
		}
		h.app.internalServerErrorResponse(w, r, err)
		return
	}

	h.app.writeJSON(w, http.StatusOK, updatedComment)
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Delete a comment with the given ID
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/comments/{id} [delete]
func (h *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.app.badRequestResponse(w, r, err)
		return
	}

	if err := h.app.commentService.DeleteComment(r.Context(), id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			h.app.notFoundResponse(w, r)
			return
		}
		h.app.internalServerErrorResponse(w, r, err)
		return
	}

	h.app.writeJSON(w, http.StatusOK, map[string]string{"message": "comment deleted successfully"})
}

// ListCommentsByPostID godoc
// @Summary List comments for a post
// @Description List comments for a specific post with pagination
// @Tags Comments
// @Accept json
// @Produce json
// @Param postID path int true "Post ID"
// @Param limit query int false "Limit" default(10)
// @Param page query int false "Page" default(1)
// @Success 200 {array} domain.Comment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/comments/post/{postID} [get]
func (h *CommentHandler) ListCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "postID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.app.badRequestResponse(w, r, err)
		return
	}

	pagination := middleware.GetPagination(r)

	comments, err := h.app.commentService.ListCommentsByPostID(r.Context(), id, pagination.Limit, pagination.Offset)
	if err != nil {
		h.app.internalServerErrorResponse(w, r, err)
		return
	}

	h.app.writeJSON(w, http.StatusOK, comments)
}

// ListCommentsByUserID godoc
// @Summary List comments by user
// @Description List comments made by a specific user with pagination
// @Tags Comments
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Param limit query int false "Limit" default(10)
// @Param page query int false "Page" default(1)
// @Success 200 {array} domain.Comment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/comments/user/{userID} [get]
func (h *CommentHandler) ListCommentsByUserID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.app.badRequestResponse(w, r, err)
		return
	}

	pagination := middleware.GetPagination(r)

	comments, err := h.app.commentService.ListCommentsByUserID(r.Context(), id, pagination.Limit, pagination.Offset)
	if err != nil {
		h.app.internalServerErrorResponse(w, r, err)
		return
	}

	h.app.writeJSON(w, http.StatusOK, comments)
}

// CountCommentsByPostID godoc
// @Summary Count comments for a post
// @Description Count the number of comments for a specific post
// @Tags Comments
// @Accept json
// @Produce json
// @Param postID path int true "Post ID"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/comments/post/{postID}/count [get]
func (h *CommentHandler) CountCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "postID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.app.badRequestResponse(w, r, err)
		return
	}

	count, err := h.app.commentService.CountCommentsByPostID(r.Context(), id)
	if err != nil {
		h.app.internalServerErrorResponse(w, r, err)
		return
	}

	h.app.writeJSON(w, http.StatusOK, map[string]int{"count": count})
}


// CountCommentsByUserID godoc
// @Summary Count comments by user
// @Description Count the number of comments made by a specific user
// @Tags Comments
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/comments/user/{userID}/count [get]
func (h *CommentHandler) CountCommentsByUserID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.app.badRequestResponse(w, r, err)
		return
	}

	count, err := h.app.commentService.CountCommentsByUserID(r.Context(), id)
	if err != nil {
		h.app.internalServerErrorResponse(w, r, err)
		return
	}

	h.app.writeJSON(w, http.StatusOK, map[string]int{"count": count})
}