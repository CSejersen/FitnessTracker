package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Csejersen/fitnessTracker/models"
	"github.com/Csejersen/fitnessTracker/storage"
	"github.com/Csejersen/fitnessTracker/utils"
)

type ExerciseHandler struct {
	Store storage.ExerciseStore
}

func (h *ExerciseHandler) HandleExercise(w http.ResponseWriter, r *http.Request) error {
	switch method := r.Method; method {
	case "GET":
		return h.HandleGetExercisesByUserID(w, r)

	case "POST":
		return h.HandleCreateExercise(w, r)

	case "DELETE":
		return h.handleDeleteExercise(w, r)

	default:
		return fmt.Errorf("Method not allowed: %s", method)
	}
}

func (h *ExerciseHandler) HandleGetExercisesByUserID(w http.ResponseWriter, r *http.Request) error {
	userID, err := utils.GetUserID(r)
	if err != nil {
		return err
	}

	exercises, err := h.Store.GetExercisesByUserID(*userID)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, exercises)
}

type CreateExerciseRequest struct {
	Name string `json:"name"`
}

func (h *ExerciseHandler) HandleCreateExercise(w http.ResponseWriter, r *http.Request) error {
	var req CreateExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("Failed to decode request body: %v", err)
	}

	if req.Name == "" {
		return fmt.Errorf("exercise_name cannot be empty")
	}

	userID, err := utils.GetUserID(r)
	if err != nil {
		return err
	}

	exercise := &models.Exercise{
		Name:   req.Name,
		UserID: *userID,
	}

	if err := h.Store.CreateExercise(exercise); err != nil {
		return err
	}

	resp := "Created Exercise: " + exercise.Name
	utils.WriteJSON(w, http.StatusOK, resp)
	return nil
}

type DeleteExerciseRequest struct {
	ID string `json:"id"`
}

func (h *ExerciseHandler) handleDeleteExercise(w http.ResponseWriter, r *http.Request) error {
	var req DeleteExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("Failed to decode request body: %v", err)
	}

	if req.ID == "" {
		return fmt.Errorf("id cannot be empty")
	}

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return fmt.Errorf("invalid ID")
	}

	if err := h.Store.DeleteExerciseByID(id); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, "Deleted exercise")
}
