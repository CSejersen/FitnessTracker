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

type UserHandler struct {
	Store storage.UserStore
}

func (h *UserHandler) HandleUser(w http.ResponseWriter, r *http.Request) error {
	switch method := r.Method; method {
	case "GET":
		return h.handleGetUser(w, r)

	case "POST":
		return h.handleCreateUser(w, r)

	case "DELETE":
		return h.handleDeleteUser(w, r)

	default:
		return fmt.Errorf("Method not allowed: %h", method)
	}
}

func (h *UserHandler) HandleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	IDstr, ok := vars["id"]
	if !ok {
		return fmt.Errorf("ID not found in request")
	}
	id, err := strconv.Atoi(IDstr)
	if err != nil {
		return fmt.Errorf("Invalid ID in request")
	}
	user, err := h.Store.GetUserByID(id)
	if err != nil {
		return err
	}
	log.Printf("Retrived user %s from db", user.Username)
	return utils.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	users, err := h.Store.GetAllUsers()
	if err != nil {
		return err
	}
	log.Printf("Retrieved all users from db")
	return utils.WriteJSON(w, http.StatusOK, users)
}

type createUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("Failed to decode request body: %v", err)
	}

	if req.Username == "" {
		return fmt.Errorf("username cannot be empty")
	}

	user := &models.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := h.Store.CreateUser(user); err != nil {
		return err
	}

	log.Printf("Created user %s", user.Username)
	resp := "Created User: " + user.Username
	utils.WriteJSON(w, http.StatusOK, resp)
	return nil
}

type deleteUserRequest struct {
	ID string `json:"id"`
}

func (h *UserHandler) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	var req deleteUserRequest
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

	if err := h.Store.DeleteUserByID(id); err != nil {
		return err
	}

	log.Printf("Deleted User with id %s", req.ID)
	return utils.WriteJSON(w, http.StatusOK, "Deleted user")
}
