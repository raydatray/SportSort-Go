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
	router.HandleFunc("GET /users/filter", handler.getUsersWithFilter)

	router.HandleFunc("POST /customers", handler.createCustomer)
	router.HandleFunc("POST /instructors", handler.createInstructor)
	router.HandleFunc("POST /owners", handler.createOwner)

	router.HandleFunc("PUT /customers/{id}", handler.updateCustomer)
	router.HandleFunc("PUT /", handler.)

	router.HandleFunc("DELETE /customers", handler.softDeleteCustomer)
	router.HandleFunc("DELETE /admin/{id}", handler.softDeleteAdmin)
	router.HandleFunc("DELETE /users/{id}", handler.confirmDelete)
}

type User struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Type     db.UserType `json:"userType"`
	Center   *int        `json:"centerId, omitempty"`
}

//Endpoint to allow customers to create new accounts
func (h *Handler) createCustomer(w http.ResponseWriter, r *http.Request) {
	var user db.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Type != db.UserTypeCustomer {
		http.Error(w, "Invalid user type for customer creation", http.StatusBadRequest)
		return
	}

	_, err := h.queries.CreateUser(ctx, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	log.Println("Received request to create a custmer")
	w.Write([]byte("Customer created"))
}

func (h *Handler) createInstructor(w http.ResponseWriter, r *http.Request) {
	var user db.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Type != db.UserTypeInstructor {
		http.Error(w, "Invalid user type for instructor creation", http.StatusBadRequest)
		return
	}

	_, err := h.queries.CreateUser(ctx, user)

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

///Endpoint to list all users regardless of type
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

//Endpoint to list all users with filters for admin
func (h *Handler) getUsersWithFilter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filterParams, err := ParseFilterUsersParams(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbParams := db.FilterUsersParams{
		Column1: filterParams.Types,
		Column2: filterParams.SportCenterIDs,
	}

	if filterParams.Deleted != nil {
		dbParams.Column3 = *filterParams.Deleted
	}

	users, err := h.queries.FilterUsers(ctx, dbParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

//Endpoint that allows a customer to update their own account
func (h *Handler) updateCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := h.queries.UpdateUser(ctx, db.)


}

//Endpoint that allows admin or owner users to update other users
func (h *Handler) updateOther(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := h.queries.UpdateUser(ctx, a)
}

//Endpoint that allows a user to delete their own account
//
func (h *Handler) softDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

}

//Endpoint that allows admin and owner users to delete other users
//
func (h *Handler) softDeleteAdmin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

}

//Endpoint that allows for the hard deletion of accounts by admins
func(h *Handler) confirmDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

}
