package domain

import "time"

// An Employees belong to the domain layer.
type Employees []Employee

// A Employee belong to the domain layer.
type Employee struct {
	ID         interface{} `json:"id" bson:"_id,omitempty"`
	FirstName  string      `json:"first_name" bson:"first_name,omitempty"`
	LastName   string      `json:"last_name" bson:"last_name,omitempty"`
	PositionID interface{} `json:"position_id" bson:"first_name,omitempty"`
	UpdatedAt  time.Time   `json:"updated_at" bson:"updated_at,omitempty"`
	CreatedAt  time.Time   `json:"created_at" bson:"created_at,omitempty"`
}
