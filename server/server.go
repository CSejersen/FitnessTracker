package server

import (
	"log"
	"net/http"

	"github.com/Csejersen/fitnessTracker/handlers"
	"github.com/Csejersen/fitnessTracker/utils"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr      string
	exerciseHandler *handlers.ExerciseHandler
}

func NewAPIServer(listenAddr string, exerciseHandler *handlers.ExerciseHandler) *APIServer {
	return &APIServer{
		listenAddr:      listenAddr,
		exerciseHandler: exerciseHandler,
	}
}

func (s *APIServer) newRouter() *mux.Router {
	router := mux.NewRouter()
	// Register routes
	router.HandleFunc("/exercise", wrapHandler(s.exerciseHandler.HandleExercise))

	return router
}

func (s *APIServer) Run() {
	router := s.newRouter()
	log.Println("JSON API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

type APIError struct {
	error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

// Wraps a handler that returns and error into a http.HandlerFunc
func wrapHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, APIError{error: err.Error()})
		}
	}
}
