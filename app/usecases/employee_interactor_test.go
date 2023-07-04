package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// MockEmployeeRepository is a mock implementation of the EmployeeRepository interface.
type MockEmployeeRepository struct {
	mock.Mock
}

// FindAllByPositionID stubs EmployeeRepository FindAllByPositionID method
func (m *MockEmployeeRepository) FindAllByPositionID(positionID string) (domain.Employees, error) {
	args := m.Called(positionID)

	employees := args.Get(0)

	resultEmployees, ok := employees.(domain.Employees)
	if !ok {
		return domain.Employees{}, args.Error(1)
	}

	return resultEmployees, args.Error(1)

}

// TestEmployeeInteractorIndex performs a unit test for the EmployeeInteractor Index method
func TestEmployeeInteractorIndex(t *testing.T) {
	assert := assert.New(t)
	mockEmployeeRepository := new(MockEmployeeRepository)

	positionID := "1"
	expectedEmployees := domain.Employees{
		{ID: 1, FirstName: "John", LastName: "Doe", PositionID: positionID},
		{ID: 2, FirstName: "Stan", LastName: "Doe", PositionID: positionID},
	}
	emptyEmployees := domain.Employees{}

	mockEmployeeRepository.On("FindAllByPositionID", positionID).Return(expectedEmployees, nil)
	mockEmployeeRepository.On("FindAllByPositionID", "invalid value").Return(emptyEmployees, errors.New("Invalid value for ID: invalid value"))
	mockEmployeeRepository.On("FindAllByPositionID", "1000").Return(emptyEmployees, errors.New("Record with id: 1000 not found"))

	interactor := EmployeeInteractor{
		EmployeeRepository: mockEmployeeRepository,
	}

	employees, err := interactor.Index(positionID)

	assert.NoError(err)
	assert.Equal(expectedEmployees, employees)

	employees, err = interactor.Index("invalid value")

	assert.Error(err)
	assert.Equal(emptyEmployees, employees)

	employees, err = interactor.Index("1000")

	assert.Error(err)
	assert.Equal(emptyEmployees, employees)
}

// FindByID stubs EmployeeRepository FindByID method
func (m *MockEmployeeRepository) FindByID(employeeID string) (domain.Employee, error) {
	args := m.Called(employeeID)
	employee := args.Get(0)

	employeeResult, ok := employee.(domain.Employee)

	if !ok {
		return domain.Employee{}, args.Error(1)
	}

	return employeeResult, args.Error(1)
}

// TestEmployeeInteractorShow performs a unit test for the EmployeeInteractor Show method
func TestEmployeeInteractorShow(t *testing.T) {
	assert := assert.New(t)
	mockEmployeeRepository := new(MockEmployeeRepository)

	employeeID := "1"
	expectedEmployee := domain.Employee{ID: employeeID, FirstName: "John", LastName: "Doe", PositionID: "1"}
	emptyEmployee := domain.Employee{}

	mockEmployeeRepository.On("FindByID", employeeID).Return(expectedEmployee, nil)
	mockEmployeeRepository.On("FindByID", "invalid value").Return(emptyEmployee, errors.New("Invalid value for ID: invalid value"))
	mockEmployeeRepository.On("FindByID", "1000").Return(emptyEmployee, errors.New("Record with id: 1000 not found"))

	interactor := EmployeeInteractor{
		EmployeeRepository: mockEmployeeRepository,
	}

	employee, err := interactor.Show("1")

	assert.NoError(err)
	assert.Equal(expectedEmployee, employee)

	employee, err = interactor.Show("invalid value")

	assert.Error(err)
	assert.Equal(emptyEmployee, employee, "should be equal")

	employee, err = interactor.Show("1000")

	assert.Error(err)
	assert.Equal(emptyEmployee, employee)
}

// DeleteByID stubs EmployeeRepository DeleteByID method
func (m *MockEmployeeRepository) DeleteByID(employeeID string) error {
	args := m.Called(employeeID)

	return args.Error(0)
}

// TestEmployeeInteractorDestroy performs a unit test for the EmployeeInteractor Destroy method
func TestEmployeeInteractorDestroy(t *testing.T) {
	assert := assert.New(t)
	mockEmployeeRepository := new(MockEmployeeRepository)

	employeeID := "1"
	mockEmployeeRepository.On("DeleteByID", employeeID).Return(nil)
	mockEmployeeRepository.On("DeleteByID", "invalid value").Return(errors.New("Invalid value for ID: invalid value"))
	mockEmployeeRepository.On("DeleteByID", "1000").Return(errors.New("Record with id: 1000 not found"))

	interactor := EmployeeInteractor{
		EmployeeRepository: mockEmployeeRepository,
	}

	err := interactor.Destroy(employeeID)

	assert.NoError(err)

	err = interactor.Destroy("invalid value")

	assert.Error(err)

	err = interactor.Destroy("1000")

	assert.Error(err)
}

// Save stubs EmployeeRepository Save method
func (m *MockEmployeeRepository) Save(employee domain.Employee) error {
	args := m.Called(employee)

	return args.Error(0)
}

// TestEmployeeInteractorStore performs a unit test for the EmployeeInteractor Store method
func TestEmployeeInteractorStore(t *testing.T) {
	assert := assert.New(t)
	mockEmployeeRepository := new(MockEmployeeRepository)

	employee := domain.Employee{FirstName: "John", LastName: "Doe", PositionID: "1"}
	emptyEmployee := domain.Employee{}

	mockEmployeeRepository.On("Save", employee).Return(nil)
	mockEmployeeRepository.On("Save", emptyEmployee).Return(errors.New("Invalid value provided"))

	interactor := EmployeeInteractor{
		EmployeeRepository: mockEmployeeRepository,
	}

	err := interactor.Store(employee)

	assert.NoError(err)

	err = interactor.Store(emptyEmployee)

	assert.Error(err)
}
