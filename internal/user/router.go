package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/raydatray/sportsort-go/db"
)

type Handler struct {
	queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
	return &Handler{queries: queries}
}

func loadUserRoutes(router *http.ServeMux, queries *db.Queries) {
	handler := NewHandler(queries)

	router.HandleFunc("GET /users", handler.getAll)
	router.HandleFunc("GET /instructors", handler.getInstructors)

	router.HandleFunc("POST /customers", handler.createCustomer)
	router.HandleFunc("POST /instructors", handler.createInstructor)
	router.HandleFunc("POST /owners", handler.createOwner)
}

type User struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Type     db.UserType `json:"userType"`
	Center   *int        `json:"centerId,omitempty"`
}

func (h *Handler) createCustomer(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Type != db.UserTypeCustomer {
		http.Error(w, "Invalid user type for customer creation", http.StatusBadRequest)
		return
	}
	//Write to DB here

	w.WriteHeader(http.StatusCreated)
	log.Println("Received request to create a custmer")
	w.Write([]byte("Customer created"))
}

func (h *Handler) createInstructor(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Type != db.UserTypeInstructor {
		http.Error(w, "Invalid user type for instructor creation", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Println("Received request to create a instructor")
	w.Write([]byte("Instructor created"))
}

func (h *Handler) createOwner(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Type != db.UserTypeOwner {
		http.Error(w, "Invalid user type for owner creation", http.StatusBadRequest)
		return
	}

	if user.Center != nil {
		http.Error(w, "Sport center must be provided", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Println("Received request to create an owner")
	w.Write([]byte("Owner created"))
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.queries.ListUsers(ctx)

	if err != nil {
		log.Printf("Error fetching users: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) getInstructors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	instructors, err := h.queries.ListUserByType(ctx, db.UserTypeInstructor)

	if err != nil {
		log.Printf("Error fetching instructors: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(instructors)
}
