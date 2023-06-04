package interfaces

import "github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"

// A EmployeeRepository belong to the interfaces layer.
type EmployeeRepository interface {
	FindAllByPositionID(positionID string) (domain.Employees, error)
	FindByID(employeeID string) (domain.Employee, error)
	DeleteByID(employeeID string) error
	Save(e domain.Employee) error
}
