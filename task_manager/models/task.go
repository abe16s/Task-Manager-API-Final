package models

import (
	"time"
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID   `bson:"_id" json:"_id"`
	Title       string    	`bson:"title" json:"title" binding:"required"`
	Description string    	`bson:"description" json:"description" binding:"required"`
	DueDate     time.Time 	`bson:"due_date" json:"due_date" binding:"required"`
	Status      string    	`bson:"status" json:"status"`
}