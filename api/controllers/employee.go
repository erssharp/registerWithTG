package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"registerWithTG/db/entities"
)

var log *zap.SugaredLogger

func CreateEmployeeRouter(r *mux.Router, logger *zap.SugaredLogger) {
	s := r.PathPrefix("/employee").Subrouter()
	log = logger
	s.HandleFunc("/add", addEmployee)
	s.HandleFunc("/edit", editEmployee)
	s.HandleFunc("/delete", deleteEmployee)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	var e entities.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		log.Error("Error while decoding json", err)
	}

}

func editEmployee(w http.ResponseWriter, r *http.Request) {

}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {

}
