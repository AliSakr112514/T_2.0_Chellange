package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	Trans "github.com/AliSakr112514/T_2.0_Chellange/Repo"
	"github.com/go-chi/chi/v5"
)

func get(w http.ResponseWriter, r *http.Request) {
	Id := chi.URLParam(r, "Id")
	//check if the id is not null
	if Id == "" {
		http.Error(w, fmt.Sprintf("No Id was Inserted"), 404)
		return
	}

	//get transaction
	transaction := &Trans.Transaction{}
	err, transaction := Trans.GetSingleTrans(Id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Id: %s is not found", Id), 404)
		return
	}
	json.NewEncoder(w).Encode(transaction)
	w.WriteHeader(200)

}
func addTrans(w http.ResponseWriter, r *http.Request) {
	transaction := &Trans.Transaction{}
	json.NewDecoder(r.Body).Decode(transaction)
	transaction.CreatedAt = time.Time.String(time.Now().UTC())
	transaction.Id = Trans.Generate_uuid()
	Trans.AddTransaction(transaction)
	w.Write([]byte("Transaction was added successfully"))
	w.WriteHeader(201)

}
func getAll(w http.ResponseWriter, r *http.Request) {
	transactions := Trans.GetAllTrans()
	json.NewEncoder(w).Encode(transactions)
	w.WriteHeader(200)
}

func mux() http.Handler {
	r := chi.NewRouter()
	r.Route("/transactions", func(r chi.Router) {
		r.Get("/", getAll)      //Get all transaction
		r.Get("/?Id={Id}", get) //Get transaction
		r.Post("/", addTrans)   //Post transaction
	})
	return r
}

func main() {
	r := mux()
	log.Fatal(http.ListenAndServe(":8080", r))
}
