package chargepoint

import (
	"context"

	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
)

type Repository interface {
	// CreateOne creates one charge point
	CreateOne(ctx context.Context, cp entity.NewChargePoint) (*entity.ChargePoint, error)
}