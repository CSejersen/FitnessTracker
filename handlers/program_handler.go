package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Csejersen/fitnessTracker/models"
	"github.com/Csejersen/fitnessTracker/storage"
	"github.com/Csejersen/fitnessTracker/utils"
	"github.com/gorilla/mux"
)

type ProgramHandler struct {
	Store storage.ProgramStore
}

func (h *ProgramHandler) HandleProgram(w http.ResponseWriter, r *http.Request) error {
	switch method := r.Method; method {
	case "GET":
		return h.HandleGetPrograms(w, r)
	case "POST":
		return h.HandleCreateProgram(w, r)
	case "DELTE":
		return h.HandleDeleteProgram(w, r)
	default:
		return fmt.Errorf("Method not allowed: %s", method)
	}
}

type AddWorkoutRequest struct {
	ID int
}

func (h *ProgramHandler) AddWorkout(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	IDstr, ok := vars["id"]
	if !ok {
		return fmt.Errorf("ID not found in request")
	}
	id, err := strconv.Atoi(IDstr)
	if err != nil {
		return fmt.Errorf("ID not valid, must be a number")
	}

	var req AddWorkoutRequest
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return fmt.Errorf("Failed to decode request body")
	}

	if err := h.Store.AddWorkout(req.ID, id); err != nil {
		return err
	}

	return nil
}

func (h *ProgramHandler) HandleGetPrograms(w http.ResponseWriter, r *http.Request) error {
	userID, err := utils.GetUserID(r)
	if err != nil {
		return err
	}

	programs, err := h.Store.GetProgramByUserID(*userID)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, programs)
}

type CreateProgramRequest struct {
	Name    string
	Split   string
	PerWeek int
}

func (h *ProgramHandler) HandleCreateProgram(w http.ResponseWriter, r *http.Request) error {
	userID, err := utils.GetUserID(r)
	if err != nil {
		return err
	}

	var req CreateProgramRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("Failed to decode request body: %v", err)
	}

	program := &models.Program{
		UserID:  *userID,
		Name:    req.Name,
		Split:   req.Split,
		PerWeek: req.PerWeek,
	}

	if err := h.Store.CreateProgram(program); err != nil {
		return err
	}
	return nil
}

type DeleteProgramRequest struct {
	ID int
}

func (h *ProgramHandler) HandleDeleteProgram(w http.ResponseWriter, r *http.Request) error {
	var req DeleteProgramRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("Failed to decode request body: %v", err)
	}
	if err := h.Store.DeleteProgram(req.ID); err != nil {
		return err
	}
	return nil
}
