package api

import (
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"registerWithTG/api/controllers"
	"strconv"
	"time"
)

type RESTAPI struct {
	server http.Server
	errors chan error
	logger *zap.SugaredLogger
}

// New returns a new instance of the REST API server.
func New(logger *zap.SugaredLogger, port int) *RESTAPI {
	router := mux.NewRouter()

	empController := controllers.EmployeeController{
		Logger: logger,
	}

	s := router.PathPrefix("/employee").Subrouter()
	s.HandleFunc("/add", empController.AddEmployee)
	s.HandleFunc("/get", empController.GetEmployee)
	s.HandleFunc("/edit", empController.EditEmployee)
	s.HandleFunc("/delete", empController.DeleteEmployee)

	return &RESTAPI{
		server: http.Server{
			Addr:    net.JoinHostPort("", strconv.Itoa(port)),
			Handler: router,
		},
		errors: make(chan error, 1),
		logger: logger,
	}
}

// Start diagnostics server.
func (rapi *RESTAPI) Start() {
	go func() {
		rapi.logger.Info("Server starting listening at", rapi.server.Addr)
		rapi.errors <- rapi.server.ListenAndServe()
		close(rapi.errors)
	}()
}

// Stop diagnostics server.
func (rapi *RESTAPI) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return rapi.server.Shutdown(ctx)
}

// Notify returns a channel to notify the caller about errors.
// If you receive an error from the channel you should stop the application.
func (rapi *RESTAPI) Notify() <-chan error {
	return rapi.errors
}
