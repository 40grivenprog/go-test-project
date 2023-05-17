package interfaces

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// A PositionRepository belong to the inteface layer
type PositionRepository struct {
	SQLHandler SQLHandler
}

// FindAll is returns the number of entities.
func (pr *PositionRepository) FindAll() (positions domain.Positions, err error) {
	const query = `
		SELECT
			id,
			name,
			salary
		FROM
			positions
	`
	rows, err := pr.SQLHandler.Query(query)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var salary int
		if err = rows.Scan(&id, &name, &salary); err != nil {
			return
		}
		position := domain.Position{
			ID:     id,
			Name:   name,
			Salary: salary,
		}
		positions = append(positions, position)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

// FindByID is returns the entity identified by the given id.
func (pr *PositionRepository) FindByID(positionID int) (position domain.Position, err error) {
	const query = `
		SELECT
			id,
			name,
			salary
		FROM
			positions
		WHERE
			id = $1
	`

	row, err := pr.SQLHandler.Query(query, positionID)

	if err != nil {
		return
	}

	defer row.Close()

	var id int
	var name string
	var salary int
	
	row.Next()

	if err = row.Scan(&id, &name, &salary); err != nil {
		return
	}

	position = domain.Position{
		ID:     id,
		Name:   name,
		Salary: salary,
	}

	return
}
