package chargepoint

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	request "github.com/successgilli/api-app/internal/controller/chargepoint/request"
	response "github.com/successgilli/api-app/internal/controller/chargepoint/response"
	ctrResponse "github.com/successgilli/api-app/internal/controller/response"
	"github.com/successgilli/api-app/internal/pkg/apperrors"
	chargePointSvc "github.com/successgilli/api-app/internal/service/chargepoint"
)

type Controller struct {
	service *chargePointSvc.Service
}

// NewController creates a new chargepoint controller
func NewController(svc *chargePointSvc.Service) Controller {
	return Controller{
		service: svc,
	}
}

// Create creates a new Charge point.
// @Summary Create a charge point
// @Description Creates a new charge point with the given name and location
// @Tags chargepoints
// @Accept json
// @Produce json
// @Param request body request.CreateChargePointRequest true "Charge point payload"
// @Success 201 {object} response.CreateChargePointResponse
// @Failure 400 {object} ctrResponse.ErrorResponse
// @Failure 500 {object} ctrResponse.ErrorResponse
// @Router /chargepoints [post]
func (cpc *Controller) Create(c *gin.Context) {
	req, err := request.NewCreateChargePointRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ctrResponse.ErrorResponse{Error: err.Error()})
		return
	}

	chargePoint, err := cpc.service.Create(c.Request.Context(), req.ToEntity())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ctrResponse.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response.CreateChargePointResponse{
		ID:     chargePoint.ID,
		Name:   chargePoint.Name,
		Lat:    chargePoint.Lat,
		Long:   chargePoint.Long,
		Status: chargePoint.Status,
	}})
}

// GetByID retrieves a single charge point by its ID.
// @Summary Get a charge point by ID
// @Description Retrieves a single charge point using its unique identifier
// @Tags chargepoints
// @Accept json
// @Produce json
// @Param id path string true "Charge point ID"
// @Success 200 {object} response.ChargePointResponse
// @Failure 400 {object} ctrResponse.ErrorResponse
// @Failure 404 {object} ctrResponse.ErrorResponse
// @Failure 500 {object} ctrResponse.ErrorResponse
// @Router /chargepoints/{id} [get]
func (cpc *Controller) GetByID(c *gin.Context) {
	chargePointID, err := request.NewGetByIDRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ctrResponse.ErrorResponse{Error: err.Error()})
		return
	}

	chargePoint, err := cpc.service.GetByID(c.Request.Context(), chargePointID)
	if err != nil {
		code := http.StatusInternalServerError
		if errors.Is(err, apperrors.ErrNotFound) {
			code = http.StatusNotFound
		}

		c.JSON(code, ctrResponse.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response.CreateChargePointResponse{
		ID:     chargePoint.ID,
		Name:   chargePoint.Name,
		Lat:    chargePoint.Lat,
		Long:   chargePoint.Long,
		Status: chargePoint.Status,
	}})
}
