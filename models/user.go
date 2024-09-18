package models

type User struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	EncryptedPassword string `json:"encrypted_password"`
}
