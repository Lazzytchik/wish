package app

import (
	"encoding/json"
	"net/http"
)

func JsonAnswer(w http.ResponseWriter, r *http.Request, value interface{}) {
	w.WriteHeader(200)

	v, err := json.Marshal(value)
	if err != nil {
		http.Error(w, "Presenting error", 500)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(v)
}
