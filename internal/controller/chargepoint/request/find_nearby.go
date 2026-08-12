package request

import (
	"github.com/gin-gonic/gin"
	svc "github.com/successgilli/api-app/internal/service/chargepoint"
)

type FindNearbyRequest struct {
	Latitude  float64 `form:"lat" binding:"required,latitude"`
	Longitude float64 `form:"lon" binding:"required,longitude"`
	RadiusKm  float64 `form:"radius" binding:"required,gt=0"`
}

// NewFindNearbyRequest creates a new find-near-by request
func NewFindNearbyRequest(ctx *gin.Context) (*FindNearbyRequest, error) {
	var req FindNearbyRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

// ToFindNearByCommand converts the request to a find-nearby command
func (r *FindNearbyRequest) ToFindNearByCommand() svc.FindNearByCommand {
	return svc.FindNearByCommand{
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
		RadiusKM:  r.RadiusKm,
	}
}
