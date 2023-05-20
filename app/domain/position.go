package domain

import "time"

// Positions belong to the domain layer.
type Positions []Position

// A Position belong to the domain layer.
type Position struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Salary    int       `json:"salary"`
	UpdatedAt time.Time `json:updated_at`
	CreatedAt time.Time `json:created_at`
}
