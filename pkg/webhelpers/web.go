package helpers

import (
	"encoding/json"
	"net/http"
)

// json implementation for one way sending
type JSON map[string]interface{}

func (h *JSON) MarshalJSON() []byte {
	bytes, _ := json.Marshal(h)
	return bytes
}

// writes status to w in json format
func HttpError(w http.ResponseWriter, status int, responseContent JSON) {
	w.WriteHeader(400)
	w.Header().Add("Content-Type", "application/json")
	a := JSON{
		"error": http.StatusText(http.StatusBadRequest),
	}
	w.Write(a.MarshalJSON())
}

// writes 404 status to w in json format
func HttpNotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	w.Header().Add("Content-Type", "application/json")
	a := JSON{
		"error": http.StatusText(http.StatusNotFound),
	}
	w.Write(a.MarshalJSON())
}

// writes 400 status to w in json format
func HttpBadRequest(w http.ResponseWriter) {
	w.WriteHeader(400)
	w.Header().Add("Content-Type", "application/json")
	a := JSON{
		"error": http.StatusText(http.StatusBadRequest),
	}
	w.Write(a.MarshalJSON())
}

// writes 500 status to w in json format
func HttpInternalError(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Header().Add("Content-Type", "application/json")
	a := JSON{
		"error": http.StatusText(http.StatusInternalServerError),
	}
	w.Write(a.MarshalJSON())
}
