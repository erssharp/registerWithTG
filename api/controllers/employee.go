package controllers

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"registerWithTG/db/entities"
)

type EmployeeController struct {
	Logger *zap.SugaredLogger
}

func (ec *EmployeeController) AddEmployee(w http.ResponseWriter, r *http.Request) {
	var e entities.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	ec.Logger.Info("add")
	if err != nil {
		ec.Logger.Error("Error while decoding json", err)
	}

}

func (ec *EmployeeController) GetEmployee(w http.ResponseWriter, r *http.Request) {
	ec.Logger.Info("get")
	var e entities.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		ec.Logger.Error("Error while decoding json", err)
	}

}

func (ec *EmployeeController) EditEmployee(w http.ResponseWriter, r *http.Request) {

}

func (ec *EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {

}
