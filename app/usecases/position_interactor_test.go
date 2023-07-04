package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// MockEmployeeRepository is a mock implementation of the EmployeeRepository interface.
type MockPositionRepository struct {
	mock.Mock
}

// FindAll stubs PositionRepository FindAll method
func (m *MockPositionRepository) FindAll() (domain.Positions, error) {
	args := m.Called()
	positions := args.Get(0)

	resultPositions, ok := positions.(domain.Positions)
	if !ok {
		return domain.Positions{}, args.Error(1)
	}

	return resultPositions, args.Error(1)
}

// TestPositionInteractorIndex performs a unit test for the PositionInteractor Index method
func TestPositionInteractorIndex(t *testing.T) {
	mockPositionRepository := new(MockPositionRepository)

	expectedPositions := domain.Positions{
		{ID: 1, Name: "Ruby", Salary: 5},
		{ID: 2, Name: "Java", Salary: 3},
	}

	mockPositionRepository.On("FindAll").Return(expectedPositions, nil)

	interactor := PositionInteractor{
		PositionRepository: mockPositionRepository,
	}

	positions, err := interactor.Index()

	assert.NoError(t, err)
	assert.Equal(t, expectedPositions, positions)
}

// FindByID stubs PositionRepository FindByID method
func (m *MockPositionRepository) FindByID(positionID string) (domain.Position, error) {
	args := m.Called(positionID)
	position := args.Get(0)

	resultPosition, ok := position.(domain.Position)

	if !ok {
		return domain.Position{}, args.Error(1)
	}

	return resultPosition, args.Error(1)
}

// TestPositionInteractorShow performs a unit test for the PositionInteractor Show method
func TestPositionInteractorShow(t *testing.T) {
	assert := assert.New(t)
	mockPositionRepository := new(MockPositionRepository)

	emptyPosition := domain.Position{}
	positionID := "1"
	expectedPosition := domain.Position{ID: positionID, Name: "Ruby", Salary: 5}
	mockPositionRepository.On("FindByID", positionID).Return(expectedPosition, nil)
	mockPositionRepository.On("FindByID", "invalid value").Return(emptyPosition, errors.New("Invalid value for ID: invalid value"))
	mockPositionRepository.On("FindByID", "1000").Return(emptyPosition, errors.New("Record with id: 1000 not found"))

	interactor := PositionInteractor{
		PositionRepository: mockPositionRepository,
	}

	position, err := interactor.Show(positionID)

	assert.NoError(err)
	assert.Equal(expectedPosition, position)

	position, err = interactor.Show("invalid value")

	assert.Error(err)
	assert.Equal(emptyPosition, position)

	position, err = interactor.Show("1000")

	assert.Error(err)
	assert.Equal(emptyPosition, position)
}

// Save stubs PositionRepository Save method
func (m *MockPositionRepository) Save(position domain.Position) error {
	args := m.Called(position)

	return args.Error(0)
}

// TestPositionInteractorStore performs a unit test for the PositionInteractor Store method
func TestPositionInteractorStore(t *testing.T) {
	assert := assert.New(t)
	mockPositionRepository := new(MockPositionRepository)

	emptyPosition := domain.Position{}
	position := domain.Position{Name: "Ruby", Salary: 5}
	mockPositionRepository.On("Save", position).Return(nil)
	mockPositionRepository.On("Save", emptyPosition).Return(errors.New("Invalid value provided"))

	interactor := PositionInteractor{
		PositionRepository: mockPositionRepository,
	}

	err := interactor.Store(position)

	assert.NoError(err)

	err = interactor.Store(emptyPosition)

	assert.Error(err)
}

// DeleteByID stubs PositionRepository DeleteByID method
func (m *MockPositionRepository) DeleteByID(positionID string) error {
	args := m.Called(positionID)

	return args.Error(0)
}

// TestPositionInteractorDestroy performs a unit test for the PositionInteractor Destroy method
func TestPositionInteractorDestroy(t *testing.T) {
	assert := assert.New(t)
	mockPositionRepository := new(MockPositionRepository)
	positionID := "1"
	mockPositionRepository.On("DeleteByID", positionID).Return(nil)
	mockPositionRepository.On("DeleteByID", "invalid value").Return(errors.New("Invalid value for ID: invalid value"))
	mockPositionRepository.On("DeleteByID", "1000").Return(errors.New("Record with id: 1000 not found"))

	interactor := PositionInteractor{
		PositionRepository: mockPositionRepository,
	}

	err := interactor.Destroy(positionID)

	assert.NoError(err)

	err = interactor.Destroy("invalid value")

	assert.Error(err)

	err = interactor.Destroy("1000")

	assert.Error(err)
}
