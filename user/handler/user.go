package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grazierShahid/microserive-blogspot/user-service/models"
	"github.com/grazierShahid/microserive-blogspot/user-service/service"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(r *mux.Router, svc service.UserService) {
	h := &UserHandler{svc}

	r.HandleFunc("/register", h.Register).Methods(http.MethodPost)
	r.HandleFunc("/login", h.Login).Methods(http.MethodPost)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
_:
	json.NewDecoder(r.Body).Decode(&user)

_:
	h.svc.Register(&user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp := h.svc.Login(&creds)

	statusCode := http.StatusOK
	if resp.Status != "success" {
		statusCode = http.StatusUnauthorized
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
