// Package request defines request shapes for the charge point controller package
package request

import (
	"github.com/gin-gonic/gin"
	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
	svc "github.com/successgilli/api-app/internal/service/chargepoint"
)

// NewCreateChargePointRequest creates a new charge point creation request
func NewCreateChargePointRequest(ctx *gin.Context) (*CreateChargePointRequest, error) {
	var chargePoint CreateChargePointRequest

	err := ctx.ShouldBindJSON(&chargePoint)
	if err != nil {
		return nil, err
	}

	return &chargePoint, nil
}

type CreateChargePointRequest struct {
	Name   string                   `json:"name" binding:"required"`
	Lat    float64                  `json:"lat" binding:"required,latitude"`
	Long   float64                  `json:"long" binding:"required,longitude"`
	Status entity.ChargePointStatus `json:"status" binding:"required,chargepoint_status"`
}

// ToCreateChargePointCommand converts the request to a create-charge-point command.
func (r *CreateChargePointRequest) ToCreateChargePointCommand() svc.CreateChargePointCommand {
	return svc.CreateChargePointCommand{
		Name:   r.Name,
		Lat:    r.Lat,
		Long:   r.Long,
		Status: r.Status,
	}
}

