package chargepoint

import (
	"context"

	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
)

// FindNearByParam represents the input necessary to find all charge points nearby
type FindNearByParam struct {
	Latitude  float64
	Longitude float64
	RadiusKM  float64
}

// CreateChargePointParam input to create a charge point
type CreateChargePointParam struct {
	Name   string
	Lat    float64
	Long   float64
	Status entity.ChargePointStatus
}

type Repository interface {
	// CreateOne creates one charge point
	CreateOne(context.Context, CreateChargePointParam) (*entity.ChargePoint, error)

	// GetByID Gets a charge point by ID
	GetByID(context.Context, string) (*entity.ChargePoint, error)

	// GetAllNearBy Gets all charge points near user
	GetAllNearBy(context.Context, FindNearByParam) ([]entity.ChargePoint, error)
}
