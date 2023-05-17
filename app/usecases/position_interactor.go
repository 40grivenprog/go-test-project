package usecases

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// A PositionInteractor belong to the usecases layer.
type PositionInteractor struct {
	PositionRepository PositionRepository
}

// Index is display a listing of the resource.
func (pi *PositionInteractor) Index() (positions domain.Positions, err error) {
	positions, err = pi.PositionRepository.FindAll()

	return
}

// Show is display the specified resource.
func (pi *PositionInteractor) Show(positionID int) (position domain.Position, err error) {
	position, err = pi.PositionRepository.FindByID(positionID)

	return
}

// Show is display the specified resource.
func (pi *PositionInteractor) Store(p domain.Position) (id int64, err error) {
	id, err = pi.PositionRepository.Save(p)

	return
}
