package handlers

import (
	"fmt"
	"github.com/Csejersen/fitnessTracker/storage"
	"github.com/Csejersen/fitnessTracker/types"
	"github.com/Csejersen/fitnessTracker/utils"
	"net/http"
)

type ExerciseHandler struct {
	Store storage.ExerciseStore
}

func (h *ExerciseHandler) HandleExercise(w http.ResponseWriter, r *http.Request) error {
	switch method := r.Method; method {
	case "GET":
		return h.handleGetExercise(w, r)

	case "POST":
		return h.handleCreateExercise(w, r)

	case "DELETE":
		return h.handleDeleteExercise(w, r)

	default:
		return fmt.Errorf("Method not allowed: %h", method)
	}
}

func (h *ExerciseHandler) handleGetExercise(w http.ResponseWriter, r *http.Request) error {
	exercise := types.NewExercise("Bench Press", "Compound", "Chest")
	return utils.WriteJSON(w, http.StatusOK, exercise)
}

func (h *ExerciseHandler) handleCreateExercise(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *ExerciseHandler) handleDeleteExercise(w http.ResponseWriter, r *http.Request) error {
	return nil
}
