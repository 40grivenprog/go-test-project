package interfaces

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// MockSQLHandler is mock for SQLHandler
type MockSQLHandler struct {
	Conn *sql.DB
	SQLHandler
}

// Query mocks Query method
func (s *MockSQLHandler) Query(query string, args ...interface{}) (Row, error) {
	rows, err := s.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	row := &MockRow{}
	row.Rows = rows

	return row, nil
}

// Exec mocks Exec method
func (s *MockSQLHandler) Exec(query string, args ...interface{}) (Result, error) {
	result, err := s.Conn.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// MockRow is mock for Row
type MockRow struct {
	Rows *sql.Rows
	Row
}

// Scan mocs Scan method
func (r MockRow) Scan(value ...interface{}) error {
	return r.Rows.Scan(value...)
}

// Next mocs Next method
func (r MockRow) Next() bool {
	return r.Rows.Next()
}

// Close mocs Closeg method
func (r MockRow) Close() error {
	return r.Rows.Close()
}

// Err mocs Err method
func (r MockRow) Err() error {
	return r.Rows.Err()
}

func createTestDatabase(t *testing.T) (*sql.DB, func()) {
	connectionString := "postgres://postgres:postgres@localhost:5432/go_test_project_test_db"
	db, dbErr := sql.Open("pgx", connectionString)
	if dbErr != nil {
		t.Fatalf("Fail to create database. %s", dbErr.Error())
	}

	rand.Seed(time.Now().UnixNano())
	schemaName := "test" + strconv.FormatInt(rand.Int63(), 10)

	_, err := db.Exec("CREATE SCHEMA " + schemaName)
	if err != nil {
		t.Fatalf("Fail to create schema. %s", err.Error())
	}

	_, err = db.Exec("SET search_path TO " + schemaName)
	if err != nil {
		t.Fatalf("Fail to set search path. %s", err.Error())
	}

	return db, func() {
		_, err := db.Exec("SET search_path TO public")
		if err != nil {
			t.Fatalf("Fail to set search path. %s", err.Error())
		}

		_, err = db.Exec("DROP SCHEMA " + schemaName + " CASCADE")
		if err != nil {
			t.Fatalf("Fail to drop schema. %s", err.Error())
		}
	}
}

func loadTestData(t *testing.T, db *sql.DB, testDataNames ...string) {
	for _, testDataName := range testDataNames {
		file, err := os.Open(fmt.Sprintf("./testdata/%s.sql", testDataName))
		if err != nil {
			t.Errorf("Failed to read file %s: %s", testDataName, err.Error())
			continue
		}
		defer file.Close()

		reader := bufio.NewReader(file)

		var query string

		for {
			line, err := reader.ReadString(';')
			if err == io.EOF {
				break
			}

			line = strings.TrimSpace(line) // Remove leading and trailing whitespace

			if line == "" {
				continue
			}

			query += line

			_, err = db.Exec(query)
			if err != nil {
				t.Errorf("Failed to exec query: %s\nERROR: %s", query, err.Error())
			}

			query = ""
		}
	}
}
