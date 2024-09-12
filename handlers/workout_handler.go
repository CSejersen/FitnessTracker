package handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/Csejersen/fitnessTracker/models"
// 	"github.com/Csejersen/fitnessTracker/storage"
// 	"github.com/Csejersen/fitnessTracker/utils"
// 	"github.com/gorilla/mux"
// )

// type WorkoutHandler struct {
// 	Store storage.ExerciseStore
// }

// func (h *ExerciseHandler) HandleWorkout(w http.ResponseWriter, r *http.Request) error {
// 	switch method := r.Method; method {
// 	case "GET":
// 		return h.handleGetWorkout(w, r)

// 	case "POST":
// 		return h.handleCreateWorkout(w, r)

// 	case "DELETE":
// 		return h.handleDeleteWorkout(w, r)

// 	default:
// 		return fmt.Errorf("Method not allowed: %h", method)
// 	}
// }

// func (h *ExerciseHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) error {
// 	vars := mux.Vars(r)
// 	IDstr, ok := vars["id"]
// 	if !ok {
// 		return fmt.Errorf("ID not found in request")
// 	}
// 	id, err := strconv.Atoi(IDstr)
// 	if err != nil {
// 		return fmt.Errorf("Invalid ID in request")
// 	}
// 	workout, err := h.Store.GetExerciseByID(id)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Retrived exercise %s from db", exercise.Name)
// 	return utils.WriteJSON(w, http.StatusOK, exercise)
// }
