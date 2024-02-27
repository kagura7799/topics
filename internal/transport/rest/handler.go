package rest

import (
	"log/slog"
	"net/http"

	"github.com/Konil-Startup/go-backend/internal/service"
	"github.com/gorilla/mux"
)

type RestHandler struct {
	Service service.Service
	l       *slog.Logger
}

func New(s service.Service, l *slog.Logger) *RestHandler {
	return &RestHandler{
		Service: s,
		l:       l,
	}
}

func (h *RestHandler) Routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/user/{user_id}", h.UserByID).Methods("GET")
	r.HandleFunc("/user/email/{email}", h.UserByEmail).Methods("GET")
	r.HandleFunc("/user", h.SaveUser).Methods("POST")
	
	r.HandleFunc("/topic", h.CreateTopic).Methods("POST")
	r.HandleFunc("/topic/{topic_id}", h.DeleteTopicByID).Methods("DELETE")

	return WriteToConsole(r)
}
