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

// NewController creates a new charge point controller
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
// @Success 201 {object} response.ChargePointResponse
// @Failure 400 {object} ctrResponse.ErrorResponse
// @Failure 500 {object} ctrResponse.ErrorResponse
// @Router /chargepoints [post]
func (cpc *Controller) Create(c *gin.Context) {
	req, err := request.NewCreateChargePointRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ctrResponse.ErrorResponse{Error: err.Error()})
		return
	}

	chargePoint, err := cpc.service.Create(c.Request.Context(), req.ToCreateChargePointCommand())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ctrResponse.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ChargePointResponse{
		Data: response.ChargePointResponseItem{
			ID:     chargePoint.ID,
			Name:   chargePoint.Name,
			Lat:    chargePoint.Lat,
			Long:   chargePoint.Long,
			Status: chargePoint.Status,
		},
	})
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

	c.JSON(http.StatusOK, response.ChargePointResponse{
		Data: response.ChargePointResponseItem{
			ID:     chargePoint.ID,
			Name:   chargePoint.Name,
			Lat:    chargePoint.Lat,
			Long:   chargePoint.Long,
			Status: chargePoint.Status,
		},
	})
}

// GetAllNearBy finds all charge points within a given radius of a location.
// @Summary Find charge points near a location
// @Description Retrieves all charge points within the given radius (in km) of the specified latitude and longitude
// @Tags chargepoints
// @Accept json
// @Produce json
// @Param lat query number true "Latitude of the search center"
// @Param lon query number true "Longitude of the search center"
// @Param radius query number true "Search radius in kilometers"
// @Success 200 {object} response.ChargePointsResponse
// @Failure 400 {object} ctrResponse.ErrorResponse
// @Failure 500 {object} ctrResponse.ErrorResponse
// @Router /chargepoints/nearby [get]
func (cpc *Controller) GetAllNearBy(c *gin.Context) {
	req, err := request.NewFindNearbyRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ctrResponse.ErrorResponse{Error: err.Error()})
		return
	}

	results, err := cpc.service.GetAllNearBy(c.Request.Context(), req.ToFindNearByCommand())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ctrResponse.ErrorResponse{Error: err.Error()})
		return
	}

	chargePoints := []response.ChargePointResponseItem{}

	for i := range results {
		chargePoints = append(chargePoints, response.ChargePointResponseItem{
			ID:     results[i].ID,
			Name:   results[i].Name,
			Lat:    results[i].Lat,
			Long:   results[i].Long,
			Status: results[i].Status,
		})
	}

	c.JSON(http.StatusOK, response.ChargePointsResponse{Data: chargePoints})
}
