package main

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
)

type Friend struct {
	ID     string  `json:"id"`
	Pessoa *Pessoa `json:"pessoa"`
	Status string  `json:"status`
}

type Pessoa struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome`
}

var friends []Friend

func getFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}

func getFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range friends {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Friend{})
}

func createFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var friend Friend
	json.NewDecoder(r.Body).Decode(&friend)

	friends = append(friends, friend)
	json.NewEncoder(w).Encode(friend)
}

func updateFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func deleteFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range friends {
		if item.ID == params["id"] {
			friends = append(friends[:index], friends[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(friends)
}

func main() {
	r := mux.NewRouter()

	friends = append(friends, Friend{ID: "1", Pessoa: &Pessoa{Nome: "William", Sobrenome: "Barbosa"}, Status: "Muito Proximo"})
	friends = append(friends, Friend{ID: "2", Pessoa: &Pessoa{Nome: "Nikolas", Sobrenome: "Carneiro"}, Status: "Muito Proximo"})

	r.HandleFunc("/api/friends", getFriends).Methods("GET")
	r.HandleFunc("/api/friends/{id}", getFriend).Methods("GET")
	r.HandleFunc("/api/friends", createFriends).Methods("POST")
	r.HandleFunc("/api/friends/{id}", updateFriends).Methods("PUT")
	r.HandleFunc("/api/friends/{id}", deleteFriends).Methods("DELETE")

	http.ListenAndServe(":7000", r)
}
