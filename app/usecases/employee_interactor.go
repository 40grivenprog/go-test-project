package usecases

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// A PositionInteractor belong to the usecases layer.
type EmployeeInteractor struct {
	EmployeeRepository EmployeeRepository
}

// Index is display a listing of the resource.
func (ei *EmployeeInteractor) Index(positionId int) (employees domain.Employees, err error) {
	employees, err = ei.EmployeeRepository.FindAllByPositionID(positionId)

	return
}

func (ei *EmployeeInteractor) Store(employee domain.Employee) (err error) {
	err = ei.EmployeeRepository.Save(employee)

	return
}

func (ei *EmployeeInteractor) Show(employeeID int) (employee domain.Employee, err error) {
	employee, err = ei.EmployeeRepository.FindByID(employeeID)

	return
}

func (ei *EmployeeInteractor) Destroy(employeeID int) (err error) {
	err = ei.EmployeeRepository.DeleteByID(employeeID)

	return
}
