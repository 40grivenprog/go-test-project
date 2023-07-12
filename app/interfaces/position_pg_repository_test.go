package interfaces

import (
	"errors"
	"testing"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestPositionPgRepositoryFindAll(t *testing.T) {
	t.Parallel()
	db, cleanup := createTestDatabase(t)
	defer cleanup()

	loadTestData(t, db, "find_all_positions")

	mockSQLHandler := MockSQLHandler{Conn: db}
	positionPgRepository := PositionPgRepository{
		SQLHandler: &mockSQLHandler,
	}

	positions, err := positionPgRepository.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, len(positions), 2)
}

func TestPositionPgRepositoryFindById(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	db, cleanup := createTestDatabase(t)
	defer cleanup()

	loadTestData(t, db, "find_by_id_position")

	mockSQLHandler := MockSQLHandler{Conn: db}
	positionPgRepository := PositionPgRepository{
		SQLHandler: &mockSQLHandler,
	}

	position, err := positionPgRepository.FindByID("1")
	assert.NoError(err)
	assert.NotEmpty(position)

	position, err = positionPgRepository.FindByID("string")
	var badRequestError BadRequestError
	if !errors.As(err, &badRequestError) { // checks if we can assign err to variable badRequestError
		t.Errorf("Invalid error type: %T. Expected: BadRequestError", err)
	}
	assert.Empty(position)

	position, err = positionPgRepository.FindByID("2")
	var recordNotFoundError RecordNotFoundError
	if !errors.As(err, &recordNotFoundError) {
		t.Errorf("Invalid error type: %T. Expected: RecordNotFoundError", err)
	}
	assert.Empty(position)
}

func TestPositionPgRepositoryDeleteByID(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	db, cleanup := createTestDatabase(t)
	defer cleanup()

	loadTestData(t, db, "delete_by_id_position")

	mockSQLHandler := MockSQLHandler{Conn: db}
	positionPgRepository := PositionPgRepository{
		SQLHandler: &mockSQLHandler,
	}

	err := positionPgRepository.DeleteByID("1")
	assert.NoError(err)

	err = positionPgRepository.DeleteByID("string")
	var badRequestError BadRequestError
	if !errors.As(err, &badRequestError) {
		t.Errorf("Invalid error type: %T. Expected: BadRequestError", err)
	}

	err = positionPgRepository.DeleteByID("2")
	var recordNotFoundError RecordNotFoundError
	if !errors.As(err, &recordNotFoundError) {
		t.Errorf("Invalid error type: %T. Expected: RecordNotFoundError", err)
	}
}

func TestPositionPgRepositorySave(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	db, cleanup := createTestDatabase(t)
	defer cleanup()

	loadTestData(t, db, "save_position")

	mockSQLHandler := MockSQLHandler{Conn: db}
	positionPgRepository := PositionPgRepository{
		SQLHandler: &mockSQLHandler,
	}

	position := domain.Position{Name: "Ruby", Salary: 200}
	err := positionPgRepository.Save(position)
	assert.NoError(err)
}
