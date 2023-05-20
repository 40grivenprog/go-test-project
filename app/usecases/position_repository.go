package usecases

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// A PositionRepository belong to the usecases layer.
type PositionRepository interface {
	FindAll() (domain.Positions, error)
	FindByID(int) (domain.Position, error)
	Save(domain.Position) error
	DeleteByID(int) error
}
