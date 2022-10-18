package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func create(r *mux.Router) {
	s := r.PathPrefix("/employee").Subrouter()
	s.HandleFunc("/add", addEmployee)
	s.HandleFunc("/edit", editEmployee)
	s.HandleFunc("/delete", deleteEmployee)
}

func addEmployee(w http.ResponseWriter, r *http.Request){

}

func editEmployee(w http.ResponseWriter, r *http.Request){

}

func deleteEmployee(w http.ResponseWriter, r *http.Request){

}

