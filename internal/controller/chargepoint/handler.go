package chargepoint

import (
	"github.com/gin-gonic/gin"
	"net/http"

	ctrResponse "github.com/successgilli/api-app/internal/controller/response"
	request "github.com/successgilli/api-app/internal/controller/chargepoint/request"
	response "github.com/successgilli/api-app/internal/controller/chargepoint/response"
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

// CreateChargePoint creates a new Charge point.
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
		c.JSON(http.StatusBadRequest, ctrResponse.ErrorResponse{ Error: err.Error()})
		return
	}

	chargePoint, err := cpc.service.Create(c.Request.Context(), req.ToEntity())

	c.JSON(http.StatusOK, gin.H{"data": response.CreateChargePointResponse{
		ID:     chargePoint.ID,
		Name:   chargePoint.Name,
		Lat:    chargePoint.Lat,
		Long:   chargePoint.Long,
		Status: chargePoint.Status,
	}})
}
