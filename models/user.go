package models

type User struct {
	Username string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
