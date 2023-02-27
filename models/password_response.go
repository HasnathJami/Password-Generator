package models

type Response struct {
	Password          string `json:"password"`
	GeneratedPassword string `json:"generated_password"`
	Message           string `json:"message"`
}