package server

import (
	"log"
	"net/http"

	"github.com/Csejersen/fitnessTracker/handlers"
	"github.com/Csejersen/fitnessTracker/utils"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr      string
	exerciseHandler *handlers.ExerciseHandler
	workoutHandler  *handlers.WorkoutHandler
	programHandler  *handlers.ProgramHandler
	userHandler     *handlers.UserHandler
	loginHandler    *handlers.LoginHandler
}

func NewAPIServer(
	listenAddr string,
	exerciseHandler *handlers.ExerciseHandler,
	userHandler *handlers.UserHandler,
	loginHandler *handlers.LoginHandler,
	workoutHandler *handlers.WorkoutHandler,
	programHandler *handlers.ProgramHandler) *APIServer {
	return &APIServer{
		listenAddr:      listenAddr,
		exerciseHandler: exerciseHandler,
		workoutHandler:  workoutHandler,
		programHandler:  programHandler,
		userHandler:     userHandler,
		loginHandler:    loginHandler,
	}
}

func (s *APIServer) NewRouter() *mux.Router {
	router := mux.NewRouter()
	// Register routes
	router.HandleFunc("/program", WrapHandler(s.programHandler.HandleProgram))
	router.HandleFunc("/program/{id:[0-9]+}/workout", WrapHandler(s.programHandler.AddWorkout)).Methods("POST")
	router.HandleFunc("/workout", WrapHandler(s.workoutHandler.HandleWorkout))
	router.HandleFunc("/workout/{id:[0-9]+}/exercise", WrapHandler(s.workoutHandler.AddExercise)).Methods("POST")
	router.HandleFunc("/exercise", WrapHandler(s.exerciseHandler.HandleExercise))
	router.HandleFunc("/user", WrapHandler(s.userHandler.HandleUser))
	router.HandleFunc("/user/{id:[0-9]+}", WrapHandler(s.userHandler.HandleGetUserByID)).Methods("GET")
	router.HandleFunc("/login", WrapHandler(s.loginHandler.HandleLogin)).Methods("POST")

	return router
}

func (s *APIServer) Run() error {
	router := s.NewRouter()
	log.Println("JSON API server running on port: ", s.listenAddr)
	// CORS setup
	corsHandler := gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://127.0.0.1:5173"}),
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:5173"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		gorillaHandlers.AllowCredentials(),
	)(router)

	return http.ListenAndServe(s.listenAddr, corsHandler)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

// Wraps a handler that returns and error into a http.HandlerFunc
func WrapHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, utils.APIError{Error: err.Error()})
		}
	}
}
