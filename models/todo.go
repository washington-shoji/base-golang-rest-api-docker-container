package models

import (
	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID `json:"id"`
	Label     string    `json:"label"`
	Completed bool      `json:"completed"`
}
