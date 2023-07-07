package interfaces

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// A PositionRepository belong to the usecases layer.
type PositionRepository interface {
	FindAll() (domain.Positions, error)
	FindByID(positionID string) (domain.Position, error)
	Save(p domain.Position) error
	DeleteByID(positionID string) error
}
