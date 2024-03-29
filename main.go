package main

import (
	"container/list"
	. "encoding/json"
	"errors"
	"github.com/flosch/pongo2"
	"github.com/gorilla/mux"
	skiplist "github.com/sean-public/fast-skiplist"
	"log"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	txs "transactions"
)

var History = skiplist.New() // dict of ids -> txs
var IdList = list.New()      // sequence of txs pointers, for easy iteration
var Balance uint64 = 0       // current balance, could be calculated from txs, but for fast access
var Generator = txs.NewRandomGenerator()
var Tpl, _ = pongo2.FromFile("index.html")

func credit(balance *uint64, amount uint64) error {
	atomic.AddUint64(balance, amount)
	return nil
}

func debit(balance *uint64, amount uint64) error {
	var mutex = &sync.Mutex{}
	var err error
	mutex.Lock()
	eval := (int64(*balance) - int64(amount)) < 0
	mutex.Unlock()
	if eval {
		err = errors.New("Invalid debit transaction")
	} else {
		atomic.AddUint64(balance, ^(amount - 1))
	}
	return err
}

func Transact(w http.ResponseWriter, r *http.Request) {
	var tx txs.TransactionDTO
	w.Header().Set("Content-type", "application/json")
	err := NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	var operation func(*uint64, uint64) error
	if tx.Type == "debit" {
		operation = debit
	} else if tx.Type == "credit" {
		operation = credit
	} else {
		e, _ := Marshal("Invalid type of transaction")
		http.Error(w, string(e), http.StatusPreconditionFailed)
		return
	}
	if operr := operation(&Balance, tx.Amount); operr == nil {
		dt := time.Now()
		tx.EmissionDate = dt.String()
		tx.Id = Generator.Generate()
		History.Set(float64(tx.Id), tx)
		IdList.PushBack(&tx)
		NewEncoder(w).Encode(tx)
	} else {
		e, _ := Marshal(operr.Error())
		http.Error(w, string(e), http.StatusBadRequest)
	}
}

func ListToSlice(l *list.List) *[]interface{} {
	res := make([]interface{}, l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		res[i] = e.Value
		i++
	}
	return &res
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	txlist := ListToSlice(IdList)
	err := NewEncoder(w).Encode(txlist)
	if err != nil {
		log.Println("Could not encode")
	}
}
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	value := History.Get(float64(id))
	if value != nil {
		NewEncoder(w).Encode(value.Value())
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func Templating(w http.ResponseWriter, r *http.Request) {
	// Execute the template per HTTP request
	txlist := ListToSlice(IdList)
	err := Tpl.ExecuteWriter(pongo2.Context{"txlist": txlist, "balance": Balance}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/transactions", Transact).Methods("POST")
	router.HandleFunc("/transactions", GetTransactions).Methods("GET")
	router.HandleFunc("/transactions/{id}", GetTransaction).Methods("GET")
	router.HandleFunc("/", Templating)
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
