package main

import (
	. "encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	SurName string `json:"surname,omitempty"`
	Age     byte   `json:"age,omitempty"`
}

var Users map[int]User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if _, ok := Users[user.Id]; ok {
		http.Error(w, "Already exists an user with the same id", http.StatusConflict)
		return
	} else {
		Users[user.Id] = user
	}
	log.Println("User: ", user, "added")
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	userList := make([]User, 0)
	for _, value := range Users {
		userList = append(userList, value)
	}
	NewEncoder(w).Encode(userList)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if user, ok := Users[id]; ok {
		NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if user, ok := Users[id]; ok {
		delete(Users, id)
		log.Println("User: ", user, "removed")
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func main() {
	Users = make(map[int]User)
	log.Println("Default users: ", Users)
	router := mux.NewRouter()
	router.HandleFunc("/transactions", CreateUser).Methods("POST")
	router.HandleFunc("/transactions", GetUsers).Methods("GET")
	router.HandleFunc("/transactions/{id}", GetUser).Methods("GET")
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
