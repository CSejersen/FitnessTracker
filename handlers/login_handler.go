package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Csejersen/fitnessTracker/auth"
	"github.com/Csejersen/fitnessTracker/config"
	"github.com/Csejersen/fitnessTracker/storage"
	"github.com/Csejersen/fitnessTracker/utils"
)

type LoginHandler struct {
	Store storage.UserStore
	Cfg   config.Config
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("Error decoding request body %v", err)
	}
	user, err := h.Store.GetUserByUsername(req.Username)
	if err != nil || !utils.CheckPasswordSimple(req.Password, user.Password) {
		return fmt.Errorf("Invalid username or password")
	}
	log.Printf("username: %s", user.Username)

	token, err := auth.GenerateJWT(user.ID, user.Username, &h.Cfg)
	if err != nil {
		return fmt.Errorf("Failed to generate token %v", err)
	}

	response := map[string]string{"token": token}
	utils.WriteJSON(w, http.StatusOK, response)
	return nil
}
