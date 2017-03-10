package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	walker *Walker
}

func NewHandler(walker *Walker) *Handler {
	return &Handler{walker}
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]
	json := h.walker.Endpoints[path]

	if json == nil {
		json = h.walker.Endpoints[path+"/"]
	}

	if json == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(json)

	if err != nil {
		http.Error(w, "Unable to write response", http.StatusInternalServerError)
	}
}
