package auth

import (
	"fmt"
	"time"

	"github.com/Csejersen/fitnessTracker/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int, userName string, cfg *config.Config) (string, error) {
	var (
		key []byte
		t   *jwt.Token
	)

	key = []byte(cfg.JWTSecret)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID":   userID,
			"userName": userName,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	return t.SignedString(key)
}

func ValidateJWT(tokenString string, cfg config.Config) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := int64(claims["exp"].(float64))
		if exp < time.Now().Unix() {
			return nil, fmt.Errorf("token has expired")
		}

		return token, nil
	}

	return nil, fmt.Errorf("invalid token")
}
