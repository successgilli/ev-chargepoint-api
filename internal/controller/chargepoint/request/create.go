package request

import (
	"github.com/gin-gonic/gin"
	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
)

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
	Lat    float64                  `json:"lat" binding:"required"`
	Long   float64                  `json:"long" binding:"required"`
	Status entity.ChargePointStatus `json:"status" binding:"required"`
}

// toEntity converts the request to an entity intent item.
func (r CreateChargePointRequest) ToEntity() entity.NewChargePoint {
	return entity.NewChargePoint(r)
}
