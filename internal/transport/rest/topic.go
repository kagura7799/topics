package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Konil-Startup/go-backend/internal/models"
	helpers "github.com/Konil-Startup/go-backend/pkg/webhelpers"
	"github.com/gorilla/mux"
)

func (h *RestHandler) CreateTopic(w http.ResponseWriter, r *http.Request) {
	const op = "rest.CreateTopic"

	var topic struct {
		Name 		string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&topic); err != nil {
		helpers.HttpBadRequest(w)
		return
	}

	if len(topic.Name) <= 1  {
		helpers.HttpError(w, 400, helpers.JSON{
            "error": "name is empty",
        })
        return
	}

	if len(topic.Description) <= 255 {
		helpers.HttpError(w, 400, helpers.JSON{
            "error": "description must be at least 255 characters long",
        })
        return
	}

	Topic := &models.Topic{
		Name: topic.Name,
        Description: topic.Description,
    }

	w.WriteHeader(201)
	if err := h.Service.CreateTopic(r.Context(), Topic); err != nil {
		helpers.HttpInternalError(w)
        return
	}
}

func (h *RestHandler) DeleteTopicByID(w http.ResponseWriter, r *http.Request) {
	const op = "rest.DeleteTopicByID"

    vars := mux.Vars(r)
    topicID, err := strconv.Atoi(vars["topic_id"])
    if err != nil || topicID < 1 {
        helpers.HttpBadRequest(w)
        return
    }

    if err := h.Service.DeleteTopicByID(r.Context(), topicID); err != nil {
        helpers.HttpInternalError(w)
        return
    }

}