package interfaces

import "github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"

// UserRepository belongs to interfaces layer
type UserRepository interface {
	Save(u domain.User) (err error)
	FindByEmail(email string) (user domain.User, err error)
}
