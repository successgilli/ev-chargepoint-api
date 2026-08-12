package chargepoint

import (
	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
)

// FindNearByCommand command to find all charge points nearby
type FindNearByCommand struct {
	Latitude  float64
	Longitude float64

	// RadiusKM is the radius from lat/long point
	RadiusKM float64
}

// CreateChargePointCommand command to create a charge point
type CreateChargePointCommand struct {
	Name   string
	Lat    float64
	Long   float64
	Status entity.ChargePointStatus
}
