package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Konil-Startup/go-backend/internal/models"
	helpers "github.com/Konil-Startup/go-backend/pkg/webhelpers"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func (h *RestHandler) UserByID(w http.ResponseWriter, r *http.Request) {
	const op = "rest.UserByID"
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil || userID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
		return
	}

	user, err := h.Service.UserByID(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
		return
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("something went wrong"))
	}
}

func (h *RestHandler) UserByEmail(w http.ResponseWriter, r *http.Request) {
	const op = "rest.UserByEmail"

	vars := mux.Vars(r)
	email := vars["email"]
	if email == "" {
		helpers.HttpBadRequest(w)
		return
	}

	user, err := h.Service.UserByEmail(r.Context(), email)
	if err != nil {
		helpers.HttpNotFound(w)
		return
	}

	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		helpers.HttpInternalError(w)
	}
}

func (h *RestHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	const op = "rest.SaveUser"

	var user struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// should use some package for validation
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.HttpBadRequest(w)
		return
	}
	if user.Email == "" {
		helpers.HttpError(w, 400, helpers.JSON{
			"error": "email is empty",
		})
		return
	}
	if user.Name == "" {
		helpers.HttpError(w, 400, helpers.JSON{
			"error": "name is empty",
		})
		return
	}
	if len(user.Password) < 8 {
		helpers.HttpError(w, 400, helpers.JSON{
			"error": "password must be at least 8 characters long",
		})
		return
	}
	helpers.HttpInternalError(w)
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		h.l.Error("Error generating password hash")
		helpers.HttpInternalError(w)
		return
	}

	User := &models.User{
		Name:  user.Name,
		Email: user.Email,
		Hash:  string(password),
	}

	w.WriteHeader(201)
	if err := h.Service.SaveUser(r.Context(), User); err != nil {
		helpers.HttpInternalError(w)
		return
	}
}
