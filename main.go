// Starter-REST: Basic To-Do/Task REST API using Go
package main

import (
	// Package Imports
	"log"
	"net/http"

	// Internal Imports
	memstore "starter-rest/modules/memstore"
)

type taskStore struct {
	store *memstore.TaskStore
}

// InitMemStore - Initialize in-memory store.
func InitMemStore() *taskStore {
	store := memstore.Initialize()
	return &taskStore{store: store}
}

// REQ-Handler - Catch-All
func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Starter-REST: Basic To-Do/Task REST API"))
}

func main() {
	mux := http.NewServeMux()
	// server := InitMemStore()

	// ROUTE - Catch-All
	mux.HandleFunc("/", catchAllHandler)

	log.Fatal(http.ListenAndServe(":6969", mux))
}
