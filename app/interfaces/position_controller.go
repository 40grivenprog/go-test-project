package interfaces

import (
	"net/http"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/gin-gonic/gin"
)

// A PositionController belong to the interface layer.
type PositionController struct {
	PositionInteractor usecases.PositionInteractor
	Logger             usecases.Logger
}

// NewPositionController returns the resource of Positions.
func NewPositionController(positionRepository PositionRepository, logger usecases.Logger) *PositionController {
	return &PositionController{
		PositionInteractor: usecases.PositionInteractor{
			PositionRepository: positionRepository,
		},
		Logger: logger,
	}
}

// Index return response which contain a listing of the resource of Positions.
func (pc *PositionController) Index(c *gin.Context) {
	positions, err := pc.PositionInteractor.Index()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, positions)
}

// Store is store a newly created resource in storage.
func (pc *PositionController) Store(c *gin.Context) {
	p := domain.Position{}

	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err := pc.PositionInteractor.Store(p)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/positions")
}

// Show returns the response containing the specified resource of a Position.
func (pc *PositionController) Show(c *gin.Context) {
	position, err := pc.PositionInteractor.Show(c.Param("id"))

	if err != nil {
		c.AbortWithError(CalculateResponseErrorStatus(err), err)
		return
	}

	c.JSON(http.StatusOK, position)
}

// Destroy is remove the specified resource from storage.
func (pc *PositionController) Destroy(c *gin.Context) {
	err := pc.PositionInteractor.Destroy(c.Param("id"))

	if err != nil {
		c.AbortWithError(CalculateResponseErrorStatus(err), err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/positions")
}
