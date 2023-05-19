package domain

import "time"

// A Users belong to the domain layer.
type Employees []Employee

// A User belong to the domain layer.
type Employee struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	PositionID int       `json:"position_id"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}
