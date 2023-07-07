package usecases

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
)

// MockUserRepository is a mock implementation of the UserRepository interface.
type MockUserRepository struct {
	mock.Mock
}

// Save stubs UserRepository Save method
func (m *MockUserRepository) Save(u domain.User) (err error) {
	args := m.Called(u)
	err = args.Error(0)

	return
}

func TestUserInteractorStore(t *testing.T) {
	assert := assert.New(t)
	mockUserRepository := new(MockUserRepository)

	user := domain.User{Email: "test@user.com", Password: "Password!"}
	emptyUser := domain.User{}
	mockUserRepository.On("Save", user).Return(nil)
	mockUserRepository.On("Save", emptyUser).Return(errors.New("Invalid value provided"))

	interactor := UserInteractor{
		UserRepository: mockUserRepository,
	}

	err := interactor.Store(emptyUser)

	assert.Error(err)

	err = interactor.Store(user)

	assert.NoError(err)
}

// FindByEmail stubs UserRepository FindByEmail method
func (m *MockUserRepository) FindByEmail(email string) (domain.User, error) {
	args := m.Called(email)
	user := args.Get(0)

	resultUser, ok := user.(domain.User)

	if !ok {
		return domain.User{}, args.Error(1)
	}

	return resultUser, args.Error(1)
}

func TestUserInteractorLoginCheck(t *testing.T) {
	assert := assert.New(t)
	mockUserRepository := new(MockUserRepository)
	os.Setenv("TOKEN_HOUR_LIFESPAN", "3600")

	userEmail := "test@user.com"
	var emptyToken string

	password := "Password"
	userWithValidCredentials := domain.User{ID: "1", Email: userEmail, Password: password}
	userWithInvalidCredentials := domain.User{ID: "1", Email: userEmail, Password: "InvalidPassword"}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userWithValidCredentials.Password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Fail to create hashedPassword. %s", err.Error())
	}
	userFromDb := domain.User{ID: "1", Email: userEmail, Password: string(hashedPassword)}

	emptyUser := domain.User{}
	mockUserRepository.On("FindByEmail", emptyUser.Email).Return(nil, errors.New("Record with email: '' not found"))
	mockUserRepository.On("FindByEmail", userWithValidCredentials.Email).Return(userFromDb, nil)

	interactor := UserInteractor{
		UserRepository: mockUserRepository,
	}

	token, err := interactor.LoginCheck(emptyUser)

	assert.Error(err)
	assert.Equal(emptyToken, token)

	token, err = interactor.LoginCheck(userWithValidCredentials)

	assert.NoError(err)

	token, err = interactor.LoginCheck(userWithInvalidCredentials)

	assert.Error(err)
	assert.Equal(emptyToken, token)
}
