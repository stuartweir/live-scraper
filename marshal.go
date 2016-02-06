package main

import (
	"encoding/json"
	"net/http"
)

func marshalJSON(w http.ResponseWriter, amznID Amazon_ID) {
	js, err := json.Marshal(amznID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
