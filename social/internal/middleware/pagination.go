package middleware

import (
	"context"
	"net/http"
	"strconv"
)

type contextKey string

const PaginateKey contextKey = "pagination"

type PaginateParams struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Page   int    `json:"page"`
	Sort   string `json:"sort"`
	Search string `json:"search"`
}

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limit := 10
		if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
			if parsed, err := strconv.Atoi(limitStr); err == nil && parsed > 0 {
				limit = parsed
			}
		}

		page := 1
		if pageStr := r.URL.Query().Get("page"); pageStr != "" {
			if parsed, err := strconv.Atoi(pageStr); err == nil && parsed > 0 {
				page = parsed
			}
		}

		sort := "id"
		if sortStr := r.URL.Query().Get("sort"); sortStr != "" {
			sort = sortStr
		}

		search := r.URL.Query().Get("search")

		offset := (page - 1) * limit

		params := PaginateParams{
			Limit:  limit,
			Offset: offset,
			Page:   page,
			Sort:   sort,
			Search: search,
		}

		ctx := context.WithValue(r.Context(), PaginateKey, params)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetPagination(r *http.Request) PaginateParams {
	if params, ok := r.Context().Value(PaginateKey).(PaginateParams); ok {
		return params
	}
	return PaginateParams{Limit: 10, Offset: 0, Page: 1}
}
