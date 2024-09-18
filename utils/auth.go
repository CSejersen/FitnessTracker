package utils

import (
	"fmt"
	"net/http"

	"github.com/Csejersen/fitnessTracker/auth"
	"github.com/Csejersen/fitnessTracker/config"
	"github.com/golang-jwt/jwt/v5"
)

func CheckPassword(reqPassword string, userPassword string) bool {
	return reqPassword == userPassword
}

func GetUserID(r *http.Request) (*int, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, fmt.Errorf("Missing token %v", err)
	}

	tokenString := cookie.Value
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	token, err := auth.ValidateJWT(tokenString, cfg)
	if err != nil {
		return nil, err
	}

	var userID int
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userID = int(claims["userID"].(float64))
	} else {
		return nil, fmt.Errorf("userID not found in claims")
	}
	return &userID, nil
}
