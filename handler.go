package main

import (
	"encoding/json"
	"net/http"
)

//Handler holds all the items and acts as a wrapper class for cleanly handling http requests
type Handler struct {
	AllItems  map[string]*Item
	FireStore *FireStore
}

func NewHandler(fs *FireStore) *Handler {
	return &Handler{AllItems: fs.LoadAllItemsFromDB(), FireStore: fs}
}

//Default redirect
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Kirk Zimmer, Cloud Computing, Go Backend")
	w.WriteHeader(http.StatusOK)
}
