package domain

import "time"

// Positions belong to the domain layer.
type Positions []Position

// A Position belong to the domain layer.
type Position struct {
	ID        interface{} `json:"id" bson:"_id,omitempty"`
	Name      string      `json:"name" bson:"name,omitempty" binding:"required"`
	Salary    int         `json:"salary" bson:"salary,omitempty" binding:"required"`
	UpdatedAt time.Time   `json:"updated_at" bson:"updated_at,omitempty"`
	CreatedAt time.Time   `json:"created_at" bson:"created_at,omitempty"`
}
