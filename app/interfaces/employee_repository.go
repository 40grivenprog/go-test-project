package interfaces

import "github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"

// A EmployeeRepository belong to the interfaces layer.
type EmployeeRepository interface {
	FindAllByPositionID(positionID int) (domain.Employees, error)
	FindByID(employeeID int) (domain.Employee, error)
	DeleteByID(employeeID int) error
	Save(employee domain.Employee) error
}
