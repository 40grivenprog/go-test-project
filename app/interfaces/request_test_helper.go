package interfaces

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/gin-gonic/gin"
)

const testAdminToken string = "admin_token"

func performAdminRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	path = fmt.Sprintf("%s?token=%s", path, testAdminToken)
	w := performRequest(r, method, path, body)

	return w
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func setupTestRouter(positionController PositionController) *gin.Engine {
	r := gin.Default()

	protected := r.Group("/api/admin")
	protected.Use(testJwtAuthMiddleware())

	protected.GET("/positions", positionController.Index)
	protected.GET("/positions/:id", positionController.Show)
	protected.POST("/positions", positionController.Store)
	protected.DELETE("/positions/:id", positionController.Destroy)

	return r
}

func testJwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func tokenValid(c *gin.Context) error {
	token := c.Query("token")
	if token == testAdminToken {
		return nil
	}

	return errors.New("Invalid token")
}

func setupTestPositionController(positionRepository PositionRepository) (positionController PositionController) {
	positionInteractor := usecases.PositionInteractor{
		PositionRepository: positionRepository,
	}

	positionController = PositionController{
		PositionInteractor: positionInteractor,
	}

	return
}
