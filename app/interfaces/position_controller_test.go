package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPositionRepository struct {
	mock.Mock
}

func (m *MockPositionRepository) FindAll() (domain.Positions, error) {
	args := m.Called()
	positions := args.Get(0)

	resultPositions, ok := positions.(domain.Positions)
	if !ok {
		return domain.Positions{}, args.Error(1)
	}

	return resultPositions, args.Error(1)
}

func (m *MockPositionRepository) FindByID(positionID string) (domain.Position, error) {
	args := m.Called(positionID)
	position := args.Get(0)

	resultPosition, ok := position.(domain.Position)

	if !ok {
		return domain.Position{}, args.Error(1)
	}

	return resultPosition, args.Error(1)
}

func (m *MockPositionRepository) Save(position domain.Position) error {
	args := m.Called(position)

	return args.Error(0)
}

func (m *MockPositionRepository) DeleteByID(positionID string) error {
	args := m.Called(positionID)

	return args.Error(0)
}

func TestPositionControllerIndex(t *testing.T) {
	t.Parallel()
	mockPositionRepository := new(MockPositionRepository)
	positionController := setupTestPositionController(mockPositionRepository)
	router := setupTestRouter(positionController)

	expectedPositions := domain.Positions{
		{ID: 1, Name: "Ruby", Salary: 5},
		{ID: 2, Name: "Java", Salary: 3},
	}
	mockPositionRepository.On("FindAll").Return(expectedPositions, nil)

	// when authorized user performs request
	w := performAdminRequest(router, "GET", "/api/admin/positions", nil)
	assert.Equal(t, http.StatusOK, w.Code)

	// when unauthorized user performs request
	w = performRequest(router, "GET", "/api/admin/positions", nil)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestPositionControllerShow(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	mockPositionRepository := new(MockPositionRepository)
	positionController := setupTestPositionController(mockPositionRepository)
	router := setupTestRouter(positionController)

	expectedPostion := domain.Position{ID: 1, Name: "Ruby", Salary: 5}
	correctID := "1"
	invalidFormatID := "string"
	recordNotFoundID := "3"

	mockPositionRepository.On("FindByID", correctID).Return(expectedPostion, nil)
	mockPositionRepository.On("FindByID", invalidFormatID).Return(nil, NewBadRequestError("position id", invalidFormatID))
	mockPositionRepository.On("FindByID", recordNotFoundID).Return(nil, NewRecordNotFoundError(recordNotFoundID))

	// when authorized user performs correct request
	path := fmt.Sprintf("/api/admin/positions/%s", correctID)
	w := performAdminRequest(router, "GET", path, nil)
	assert.Equal(http.StatusOK, w.Code)

	// when unauthorized user performs correct request
	path = fmt.Sprintf("/api/admin/positions/%s", correctID)
	w = performRequest(router, "GET", path, nil)
	assert.Equal(http.StatusUnauthorized, w.Code)

	// when authorized user performs incorrect requests
	path = fmt.Sprintf("/api/admin/positions/%s", invalidFormatID)
	w = performAdminRequest(router, "GET", path, nil)
	assert.Equal(http.StatusBadRequest, w.Code)

	path = fmt.Sprintf("/api/admin/positions/%s", recordNotFoundID)
	w = performAdminRequest(router, "GET", path, nil)
	assert.Equal(http.StatusNotFound, w.Code)

}

func TestPositionControllerStore(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	mockPositionRepository := new(MockPositionRepository)
	positionController := setupTestPositionController(mockPositionRepository)
	router := setupTestRouter(positionController)

	correctRequestBody := `{"name": "Ruby", "salary": 20}`
	incorrectRequestBody := `{"name": "Ruby"}`

	var expectedPostion domain.Position
	err := json.Unmarshal([]byte(correctRequestBody), &expectedPostion)
	if err != nil {
		panic(err)
	}

	mockPositionRepository.On("Save", expectedPostion).Return(nil)

	path := "/api/admin/positions"

	// when authorized user performs correct request
	w := performAdminRequest(router, "POST", path, bytes.NewBuffer([]byte(correctRequestBody)))
	assert.Equal(http.StatusSeeOther, w.Code)

	// when unauthorized user performs correct request
	w = performRequest(router, "POST", path, bytes.NewBuffer([]byte(correctRequestBody)))
	assert.Equal(http.StatusUnauthorized, w.Code)

	// when authorized user performs incorrect request
	w = performAdminRequest(router, "POST", path, bytes.NewBuffer([]byte(incorrectRequestBody)))
	assert.Equal(http.StatusBadRequest, w.Code)
}

func TestPositionControllerDestroy(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	mockPositionRepository := new(MockPositionRepository)
	positionController := setupTestPositionController(mockPositionRepository)
	router := setupTestRouter(positionController)

	correctID := "1"
	invalidFormatID := "string"
	recordNotFoundID := "3"

	mockPositionRepository.On("DeleteByID", correctID).Return(nil)
	mockPositionRepository.On("DeleteByID", invalidFormatID).Return(NewBadRequestError("position id", invalidFormatID))
	mockPositionRepository.On("DeleteByID", recordNotFoundID).Return(NewRecordNotFoundError(recordNotFoundID))

	// when authorized user performs correct request
	path := fmt.Sprintf("/api/admin/positions/%s", correctID)
	w := performAdminRequest(router, "DELETE", path, nil)
	assert.Equal(http.StatusSeeOther, w.Code)

	// when unauthorized user performs correct request
	w = performRequest(router, "DELETE", path, nil)
	assert.Equal(http.StatusUnauthorized, w.Code)

	// when authorized user performs incorrect request
	path = fmt.Sprintf("/api/admin/positions/%s", invalidFormatID)
	w = performAdminRequest(router, "DELETE", path, nil)
	assert.Equal(http.StatusBadRequest, w.Code)

	// when authorized user performs incorrect request
	path = fmt.Sprintf("/api/admin/positions/%s", recordNotFoundID)
	w = performAdminRequest(router, "DELETE", path, nil)
	assert.Equal(http.StatusNotFound, w.Code)
}
