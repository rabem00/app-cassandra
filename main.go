package main

import (
	"app-cassandra/Cassandra"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)

	router.HandleFunc("/users", Users.Get)
	router.HandleFunc("/users/new", Users.Post)
	router.HandleFunc("/users/{user_uuid}", Users.GetOne)

	router.HandleFunc("/messages", Messages.Get)
	router.HandleFunc("/messages/new", Messages.Post)
	router.HandleFunc("/messages/{message_uuid}", Messages.GetOne)

	log.Fatal(http.ListenAndServe(":8080", router))
}

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
