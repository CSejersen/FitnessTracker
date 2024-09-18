package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
	if err != nil || !utils.CheckPassword(req.Password, user.EncryptedPassword) {
		return fmt.Errorf("Invalid username or password")
	}
	log.Printf("username: %s", user.Username)

	tokenString, err := auth.GenerateJWT(user.ID, user.Username, &h.Cfg)
	if err != nil {
		return fmt.Errorf("Failed to generate token %v", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false, // Ensure this is true in production (use HTTPS)
		SameSite: http.SameSiteLaxMode,
	})

	utils.WriteJSON(w, http.StatusOK, "log in succesful")
	return nil
}
