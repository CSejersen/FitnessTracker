package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Csejersen/fitnessTracker/models"
	"github.com/Csejersen/fitnessTracker/storage"
	"github.com/Csejersen/fitnessTracker/utils"
	"github.com/gorilla/mux"
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

func (h *ExerciseHandler) HandleGetExerciseByID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	IDstr, ok := vars["id"]
	if !ok {
		return fmt.Errorf("ID not found in request")
	}
	id, err := strconv.Atoi(IDstr)
	if err != nil {
		return fmt.Errorf("Invalid ID in request")
	}
	exercise, err := h.Store.GetExerciseByID(id)
	if err != nil {
		return err
	}
	log.Printf("Retrived exercise %s from db", exercise.Name)
	return utils.WriteJSON(w, http.StatusOK, exercise)
}

func (h *ExerciseHandler) handleGetExercise(w http.ResponseWriter, r *http.Request) error {
	exercises, err := h.Store.GetAllExercises()
	if err != nil {
		return err
	}
	log.Printf("Retrieved all exercises from db")
	return utils.WriteJSON(w, http.StatusOK, exercises)
}

type createExerciseRequest struct {
	Name string `json:"name"`
}

func (h *ExerciseHandler) handleCreateExercise(w http.ResponseWriter, r *http.Request) error {
	var req createExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("Failed to decode request body: %v", err)
	}

	if req.Name == "" {
		return fmt.Errorf("exercise_name cannot be empty")
	}

	exercise := &models.Exercise{
		Name: req.Name,
	}
	if err := h.Store.CreateExercise(exercise); err != nil {
		return err
	}

	log.Printf("Created exercise %s", exercise.Name)
	resp := "Created Exercise: " + exercise.Name
	utils.WriteJSON(w, http.StatusOK, resp)
	return nil
}

type deleteExerciseRequest struct {
	ID string `json:"id"`
}

func (h *ExerciseHandler) handleDeleteExercise(w http.ResponseWriter, r *http.Request) error {
	var req deleteExerciseRequest
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

	log.Printf("Deleted exercise with id %s", req.ID)
	return utils.WriteJSON(w, http.StatusOK, "Deleted exercise")
}
