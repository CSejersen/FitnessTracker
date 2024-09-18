package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Csejersen/fitnessTracker/models"
	"github.com/Csejersen/fitnessTracker/storage"
	"github.com/Csejersen/fitnessTracker/utils"
)

type WorkoutHandler struct {
	Store storage.WorkoutStore
}

func (h *WorkoutHandler) HandleWorkout(w http.ResponseWriter, r *http.Request) error {
	switch method := r.Method; method {
	case "GET":
		return h.HandleGetWorkoutByUserID(w, r)

	case "POST":
		return h.HandleCreateWorkout(w, r)

	case "DELETE":
		return h.HandleDeleteWorkout(w, r)

	default:
		return fmt.Errorf("Method not allowed: %h", method)
	}
}

func (h *WorkoutHandler) HandleGetWorkoutByUserID(w http.ResponseWriter, r *http.Request) error {
	userID, err := utils.GetUserID(r)
	if err != nil {
		return err
	}

	workouts, err := h.Store.GetAllWorkoutsByUserID(*userID)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, workouts)
}

type CreateWorkoutRequest struct {
	Name string
}

func (h *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) error {
	userID, err := utils.GetUserID(r)
	if err != nil {
		return err
	}

	var req CreateWorkoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("Failed to decode request body: %v", err)
	}

	workout := &models.Workout{
		Name:      req.Name,
		UserID:    *userID,
		Exercises: []models.Exercise{},
	}

	if err := h.Store.CreateWorkout(workout); err != nil {
		return err
	}
	return nil
}

type DeleteWorkoutRequest struct {
	ID int
}

func (h *WorkoutHandler) HandleDeleteWorkout(w http.ResponseWriter, r *http.Request) error {
	var req DeleteWorkoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("Failed to decode request body: %v", err)
	}

	if err := h.Store.DeleteWorkoutByID(req.ID); err != nil {
		return err
	}
	return nil
}
