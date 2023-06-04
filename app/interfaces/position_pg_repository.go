package interfaces

import (
	"strconv"
	"time"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// A PositionPgRepository belong to the inteface layer
type PositionPgRepository struct {
	SQLHandler SQLHandler
}

// FindAll is returns the number of entities.
func (pr *PositionPgRepository) FindAll() (positions domain.Positions, err error) {
	const query = `
		SELECT
			id,
			name,
			salary,
			updated_at,
			created_at
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
		var updatedAt time.Time
		var createdAt time.Time
		if err = rows.Scan(&id, &name, &salary, &updatedAt, &createdAt); err != nil {
			return
		}
		position := domain.Position{
			ID:        id,
			Name:      name,
			Salary:    salary,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
		positions = append(positions, position)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

// FindByID is returns the entity identified by the given id.
func (pr *PositionPgRepository) FindByID(positionID string) (position domain.Position, err error) {
	positionIDint, err := strconv.Atoi(positionID)

	if err != nil {
		return
	}

	const query = `
		SELECT
			id,
			name,
			salary,
			updated_at,
			created_at
		FROM
			positions
		WHERE
			id = $1
	`

	row, err := pr.SQLHandler.Query(query, positionIDint)

	if err != nil {
		return
	}

	defer row.Close()

	var id int
	var name string
	var salary int
	var updatedAt time.Time
	var createdAt time.Time

	row.Next()

	if err = row.Scan(&id, &name, &salary, &updatedAt, &createdAt); err != nil {
		return
	}

	position = domain.Position{
		ID:        id,
		Name:      name,
		Salary:    salary,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return
}

// Save is saves the given entity
func (pr *PositionPgRepository) Save(p domain.Position) (err error) {
	const query = `
		INSERT INTO
				positions(name, salary)
			VALUES
				($1, $2)
	`

	_, err = pr.SQLHandler.Exec(query, p.Name, p.Salary)

	if err != nil {
		return
	}

	return
}

// DeleteByID is deletes the entity identified by the given id.
func (pr *PositionPgRepository) DeleteByID(positionID string) (err error) {
	positionIDint, err := strconv.Atoi(positionID)

	if err != nil {
		return
	}

	const query = `
		DELETE
		FROM
			positions
		WHERE
			id = $1
	`

	_, err = pr.SQLHandler.Exec(query, positionIDint)

	if err != nil {
		return
	}

	return
}
