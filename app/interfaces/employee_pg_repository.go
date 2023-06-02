package interfaces

import (
	"time"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// An EmployeePgRepository belong to the inteface layer
type EmployeePgRepository struct {
	SQLHandler SQLHandler
}

// FindAllByPositionID returns all entities by position id.
func (er *EmployeePgRepository) FindAllByPositionID(positionID int) (employees domain.Employees, err error) {
	const query = `
	SELECT
		id,
		first_name,
		last_name,
		position_id,
		updated_at,
		created_at
	FROM
		employees
	WHERE
	  position_id = $1
	`

	rows, err := er.SQLHandler.Query(query, positionID)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		var positionID int
		var updatedAt time.Time
		var createdAt time.Time

		if err = rows.Scan(&id, &firstName, &lastName, &positionID, &updatedAt, &createdAt); err != nil {
			return
		}

		employee := domain.Employee{
			ID:         id,
			FirstName:  firstName,
			LastName:   lastName,
			PositionID: positionID,
			UpdatedAt:  updatedAt,
			CreatedAt:  createdAt,
		}

		employees = append(employees, employee)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

// FindByID returns the entity identified by the given id.
func (er *EmployeePgRepository) FindByID(employeeID int) (employee domain.Employee, err error) {
	const query = `
	SELECT
		id,
		first_name,
		last_name,
		position_id,
		updated_at,
		created_at
	FROM
		employees
	WHERE
		id = $1
	`
	row, err := er.SQLHandler.Query(query, employeeID)
	if err != nil {
		return
	}

	defer row.Close()

	var id int
	var firstName string
	var lastName string
	var positionID int
	var updatedAt time.Time
	var createdAt time.Time

	row.Next()

	if err = row.Scan(&id, &firstName, &lastName, &positionID, &updatedAt, &createdAt); err != nil {
		return
	}

	employee = domain.Employee{
		ID:         id,
		FirstName:  firstName,
		LastName:   lastName,
		PositionID: positionID,
		UpdatedAt:  updatedAt,
		CreatedAt:  createdAt,
	}

	return
}

// DeleteByID is deletes the entity identified by the given id.
func (er *EmployeePgRepository) DeleteByID(employeeID int) (err error) {
	const query = `
	DELETE
	FROM
	  employees
	WHERE
	  id = $1
	`
	_, err = er.SQLHandler.Exec(query, employeeID)

	if err != nil {
		return
	}

	return
}

// Save is saves the given entity
func (er *EmployeePgRepository) Save(employee domain.Employee) (err error) {
	const query = `
	INSERT INTO
		employees(first_name, last_name, position_id)
	VALUES
		($1, $2, $3)
	`
	_, err = er.SQLHandler.Exec(query, employee.FirstName, employee.LastName, employee.PositionID)

	if err != nil {
		return
	}

	return
}
