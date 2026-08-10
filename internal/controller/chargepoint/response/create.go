package response

import (
	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
)

type ChargePointResponse struct {
	Data CreateChargePointResponse `json:"data"`
}

type CreateChargePointResponse struct {
	ID     int64                      `json:"id"`
	Name   string                   `json:"name"`
	Lat    float64                  `json:"lat"`
	Long    float64                  `json:"long"`
	Status entity.ChargePointStatus `json:"status"`
}
