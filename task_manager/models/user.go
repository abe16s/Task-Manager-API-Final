package models

import "github.com/google/uuid"

// A user struct with id, username and password with json and bson tags
type User struct {
	ID       uuid.UUID 	`json:"id" bson:"_id"`
	Username string     `json:"username" bson:"username" binding:"required"`
	Password string     `json:"password" bson:"password" binding:"required"`
	IsAdmin  bool       `json:"is_admin" bson:"is_admin"`
}

