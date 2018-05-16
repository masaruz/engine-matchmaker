package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

type player struct {
	Name string
}

type session struct {
	ID      string
	Players map[string]player
}

type sessions map[string]session

func createSession() session {
	id := uuid.Must(uuid.NewV4()).String()
	return session{
		ID:      id,
		Players: make(map[string]player)}
}

func main() {
	port := 8080
	sessions := make(map[string]session)
	r := mux.NewRouter()

	r.HandleFunc("/sessions", func(w http.ResponseWriter, r *http.Request) {
		session := createSession()
		sessions[session.ID] = session
		json, _ := json.Marshal(sessions)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}).Methods("POST")

	r.HandleFunc("/sessions/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		session := sessions[id]
		json, _ := json.Marshal(session)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}).Methods("GET")

	r.HandleFunc("/sessions/reset", func(w http.ResponseWriter, r *http.Request) {
		sessions = make(map[string]session)
	}).Methods("post")

	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
