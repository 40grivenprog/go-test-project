package interfaces

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"golang.org/x/crypto/bcrypt"
)

// UserPgRepository belongs to interfaces layer
type UserPgRepository struct {
	SQLHandler SQLHandler
}

// Save creates new record of User 
func (ur *UserPgRepository) Save(user domain.User) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	user.Password = string(hashedPassword)

	const query = `
	INSERT INTO
			users(email, password)
		VALUES
			($1, $2)
  `

	_, err = ur.SQLHandler.Exec(query, user.Email, user.Password)

	if err != nil {
		return
	}

	return
}

// FindByEmail returns user by email
func (ur *UserPgRepository) FindByEmail(email string) (user domain.User, err error) {
	const query = `
	SELECT
	  id,
		email,
		password
	FROM
	  users
	WHERE
	  email = $1
  `

	row, err := ur.SQLHandler.Query(query, email)

	if err != nil {
		return
	}

	if !row.Next() {
		err = NewInvalidCredentialsError()
		return
	}

	defer row.Close()

	var id int
	var fetchedEmail, password string

	if err = row.Scan(&id, &fetchedEmail, &password); err != nil {
		return
	}

	user = domain.User{
		ID:       id,
		Email:    fetchedEmail,
		Password: password,
	}

	return
}
