package api

import (
	"encoding/json"
	"net/http"
)

func renderJSON(w http.ResponseWriter, v interface{}) {
	json.NewEncoder(w).Encode(v)
}
