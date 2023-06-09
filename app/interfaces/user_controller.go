package interfaces

import (
	"net/http"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/gin-gonic/gin"
)

// A UserController belongs to the interface layer.
type UserController struct {
	UserInteractor usecases.UserInteractor
	Logger         usecases.Logger
}

// NewUserController creates new user controller.
func NewUserController(dbHandler interface{}, logger usecases.Logger) *UserController {
	var userRepository UserRepository

	switch dbHandler.(type) {
	case SQLHandler:
		sqlHandler, _ := dbHandler.(SQLHandler)
		userRepository = &UserPgRepository{
			SQLHandler: sqlHandler,
		}
	case MongoDBHandler:
		mongoDbHandler, _ := dbHandler.(MongoDBHandler)
		userRepository = &UserMongoRepository{
			MongoDBHandler: mongoDbHandler,
		}
	}

	return &UserController{
		UserInteractor: usecases.UserInteractor{
			UserRepository: userRepository,
		},
		Logger: logger,
	}
}

// Register is an endpoint for registration
func (uc *UserController) Register(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := uc.UserInteractor.Store(user)

	if err != nil {
		c.AbortWithError(CalculateResponseErrorStatus(err), err)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

// Login is an endpoint for login
func (uc *UserController) Login(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := uc.UserInteractor.LoginCheck(user)
	if err != nil {
		err = NewInvalidCredentialsError()
		c.AbortWithError(CalculateResponseErrorStatus(err), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
