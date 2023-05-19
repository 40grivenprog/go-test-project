package interfaces

import (
	"time"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

type EmployeeRepository struct {
	SQLHandler SQLHandler
}

func (er *EmployeeRepository) FindAllByPositionID(positionId int) (employees domain.Employees, err error) {
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

	rows, err := er.SQLHandler.Query(query, positionId)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		var positionId int
		var updatedAt time.Time
		var createdAt time.Time

		if err = rows.Scan(&id, &firstName, &lastName, &positionId, &updatedAt, &createdAt); err != nil {
			return
		}

		employee := domain.Employee{
			ID:         id,
			FirstName:  firstName,
			LastName:   lastName,
			PositionID: positionId,
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

func (er *EmployeeRepository) FindByID(employeeID int) (employee domain.Employee, err error) {
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
	var positionId int
	var updatedAt time.Time
	var createdAt time.Time

	row.Next()

	if err = row.Scan(&id, &firstName, &lastName, &positionId, &updatedAt, &createdAt); err != nil {
		return
	}

	employee = domain.Employee{
		ID:         id,
		FirstName:  firstName,
		LastName:   lastName,
		PositionID: positionId,
		UpdatedAt:  updatedAt,
		CreatedAt:  createdAt,
	}

	return
}

func (er *EmployeeRepository) DeleteByID(employeeID int) (err error) {
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

func (er *EmployeeRepository) Save(employee domain.Employee) (err error) {
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
