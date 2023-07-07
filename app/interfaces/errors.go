package interfaces

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
)

// RecordNotFoundError belongs to interfaces layer
type RecordNotFoundError struct {
	errorMessage string
}

func (err RecordNotFoundError) Error() string {
	return err.errorMessage
}

// BadRequestError belongs to interfaces layer
type BadRequestError struct {
	errorMessage string
}

func (err BadRequestError) Error() string {
	return err.errorMessage
}

// NewBadRequestError creates new BadRequestError
func NewBadRequestError(key string, value interface{}) BadRequestError {
	return BadRequestError{errorMessage: fmt.Sprintf("Invalid value for %s: %s", key, value)}
}

// NewRecordNotFoundError creates new RecordNotFoundError
func NewRecordNotFoundError(value interface{}) RecordNotFoundError {
	return RecordNotFoundError{errorMessage: fmt.Sprintf("Record with id: %s not found", value)}
}

// CalculateResponseErrorStatus calculates status depends on error type
func CalculateResponseErrorStatus(err error) (status int) {
	switch err.(type) {
	case RecordNotFoundError:
		status = http.StatusNotFound
	case BadRequestError:
		status = http.StatusBadRequest
	case *pgconn.PgError:
		status = calculatePGErrorStatus(err)
	default:
		status = http.StatusInternalServerError
	}

	return
}

func calculatePGErrorStatus(err error) (status int) {
	pgError, _ := err.(*pgconn.PgError)
	status = http.StatusInternalServerError
	
	if pgError.Code == "23503" {
		status = http.StatusBadRequest
	}

	return
}
