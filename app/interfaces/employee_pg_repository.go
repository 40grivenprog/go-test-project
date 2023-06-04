package interfaces

import (
	"strconv"
	"time"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// An EmployeePgRepository belong to the inteface layer
type EmployeePgRepository struct {
	SQLHandler SQLHandler
}

// FindAllByPositionID returns all entities by position id.
func (er *EmployeePgRepository) FindAllByPositionID(positionID string) (employees domain.Employees, err error) {
	positionIDInt, err := strconv.Atoi(positionID)

	if err != nil {
		err = NewBadRequestError("position id", positionID)
		return
	}

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

	rows, err := er.SQLHandler.Query(query, positionIDInt)

	if err != nil {
		return
	}

	if !rows.Next() {
		err = NewRecordNotFoundError(positionID)
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
func (er *EmployeePgRepository) FindByID(employeeID string) (employee domain.Employee, err error) {
	employeeIDInt, err := strconv.Atoi(employeeID)

	if err != nil {
		err = NewBadRequestError("employee id", employeeID)
		return
	}

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
	row, err := er.SQLHandler.Query(query, employeeIDInt)
	if err != nil {
		return
	}

	if !row.Next() {
		err = NewRecordNotFoundError(employeeID)
		return
	}

	defer row.Close()

	var id int
	var firstName string
	var lastName string
	var positionID int
	var updatedAt time.Time
	var createdAt time.Time

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
func (er *EmployeePgRepository) DeleteByID(employeeID string) (err error) {
	employeeIDInt, err := strconv.Atoi(employeeID)

	if err != nil {
		err = NewBadRequestError("employee id", employeeID)
		return
	}

	const query = `
	DELETE
	FROM
	  employees
	WHERE
	  id = $1
	`
	result, err := er.SQLHandler.Exec(query, employeeIDInt)

	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 || err != nil {
		return NewRecordNotFoundError(employeeID)
	}

	return
}

// Save is saves the given entity
func (er *EmployeePgRepository) Save(e domain.Employee) (err error) {
	const query = `
	INSERT INTO
		employees(first_name, last_name, position_id)
	VALUES
		($1, $2, $3)
	`
	_, err = er.SQLHandler.Exec(query, e.FirstName, e.LastName, e.PositionID)

	if err != nil {
		return
	}

	return
}
