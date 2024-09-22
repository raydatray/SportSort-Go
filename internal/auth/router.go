package auth

import (
	"net/http"

	"github.com/raydatray/sportsort-go/db"
)

type Handler struct {
	queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
	return &Handler{queries: queries}
}

func loadAuthRoutes(router *http.ServeMux, queries *db.Queries) {
	handler := NewHandler(queries)

	router.HandleFunc("POST /login", handler)
	router.HandleFunc("POST /logout", handler)

}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {

}
