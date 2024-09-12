package auth

import (
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
		})

	return t.SignedString(key)
}

func ValidateJWT(tokenString string, cfg config.Config) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
}
