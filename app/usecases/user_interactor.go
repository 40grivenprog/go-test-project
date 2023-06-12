package usecases

import (
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/utils/token"
	"golang.org/x/crypto/bcrypt"
)

// An UserInteractor belong to the usecases layer
type UserInteractor struct {
	UserRepository UserRepository
}

// Store is registering a new user
func (ui *UserInteractor) Store(u domain.User) (err error) {
	err = ui.UserRepository.Save(u)

	return
}

// LoginCheck check if user can be logged in
func (ui *UserInteractor) LoginCheck(user domain.User) (tokenString string, err error) {
	dbUser, err := ui.UserRepository.FindByEmail(user.Email)
	if err != nil {
		return
	}

	err = verifyPassword(user.Password, dbUser.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	tokenString, err = token.GenerateToken(dbUser.ID)

	if err != nil {
		return "", err
	}

	return
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
