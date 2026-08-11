package chargepoint

import (
	"context"

	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
)

type Repository interface {
	// CreateOne creates one charge point
	CreateOne(context.Context, entity.NewChargePoint) (*entity.ChargePoint, error)

	// GetByID Gets a charge point by ID
	GetByID(context.Context, string) (*entity.ChargePoint, error)
}
