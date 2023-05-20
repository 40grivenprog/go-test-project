package usecases

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// An EmployeeInteractor belong to the usecases layer.
type EmployeeInteractor struct {
	EmployeeRepository EmployeeRepository
}

// Index is display a listing of the resource.
func (ei *EmployeeInteractor) Index(positionID int) (employees domain.Employees, err error) {
	employees, err = ei.EmployeeRepository.FindAllByPositionID(positionID)

	return
}

// Store is store a newly created resource in storage.
func (ei *EmployeeInteractor) Store(employee domain.Employee) (err error) {
	err = ei.EmployeeRepository.Save(employee)

	return
}

// Show is display the specified resource.
func (ei *EmployeeInteractor) Show(employeeID int) (employee domain.Employee, err error) {
	employee, err = ei.EmployeeRepository.FindByID(employeeID)

	return
}

// Destroy is remove the specified resource from storage.
func (ei *EmployeeInteractor) Destroy(employeeID int) (err error) {
	err = ei.EmployeeRepository.DeleteByID(employeeID)

	return
}
