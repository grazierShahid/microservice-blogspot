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
