package main

import (
	. "encoding/json"
	"github.com/gorilla/mux"
	"log"
	txs "money-accounting-system/transactions"
	"net/http"
	"strconv"
)

var Users map[int]txs.TransactionDTO
var Service = txs.NewTransactionService()

func Transact(w http.ResponseWriter, r *http.Request) {
	var tx txs.TransactionDTO
	err := NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	t, err := txs.ProcessTransactionType(tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}
	if id, err := Service.Transact(t); err != nil {
		http.Error(w, "Already exists an tx with the same id", http.StatusConflict)
	} else {
		NewEncoder(w).Encode(id)
	}
	log.Println("TransactionDTO: ", id, "added")
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	userList := make([]txs.TransactionDTO, 0)
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
		log.Println("TransactionDTO: ", user, "removed")
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func main() {
	Users = make(map[int]txs.TransactionDTO)
	log.Println("Default users: ", Users)
	router := mux.NewRouter()
	router.HandleFunc("/transactions", Transact).Methods("POST")
	router.HandleFunc("/transactions", GetUsers).Methods("GET")
	router.HandleFunc("/transactions/{id}", GetUser).Methods("GET")
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
