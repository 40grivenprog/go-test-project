package domain

import "time"

// An Employees belong to the domain layer.
type Employees []Employee

// A Employee belong to the domain layer.
type Employee struct {
	ID         string    `json:"id" bson:"_id,omitempty"`
	FirstName  string    `json:"first_name" bson:"first_name,omitempty" binding:"required"`
	LastName   string    `json:"last_name" bson:"last_name,omitempty" binding:"required"`
	PositionID string    `json:"position_id" bson:"position,omitempty" binding:"required"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at,omitempty"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at,omitempty"`
}
