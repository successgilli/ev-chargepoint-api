// Package response defines response shapes for the charge point controller package
package response

import (
	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
)

// ChargePointResponse is a Charge point response.
type ChargePointResponse struct {
	Data ChargePointResponseItem `json:"data"`
}

// ChargePointsResponse is a Charge point response.
type ChargePointsResponse struct {
	Data []ChargePointResponseItem `json:"data"`
}

// ChargePointResponseItem represents a single charge point response item.
type ChargePointResponseItem struct {
	ID     string                   `json:"id"`
	Name   string                   `json:"name"`
	Lat    float64                  `json:"lat"`
	Long   float64                  `json:"long"`
	Status entity.ChargePointStatus `json:"status"`
}
