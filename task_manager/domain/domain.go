package domain

import (
	"time"
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID   `bson:"_id" json:"id"`
	Title       string    	`bson:"title" json:"title" binding:"required"`
	Description string    	`bson:"description" json:"description" binding:"required"`
	DueDate     time.Time 	`bson:"due_date" json:"due_date" binding:"required"`
	Status      string    	`bson:"status" json:"status"`
}

// A user struct with id, username and password with json and bson tags
type User struct {
	ID       uuid.UUID 	`json:"id" bson:"_id"`
	Username string     `json:"username" bson:"username" binding:"required"`
	Password string     `json:"password" bson:"password" binding:"required"`
	IsAdmin  bool       `json:"is_admin" bson:"is_admin"`
}

