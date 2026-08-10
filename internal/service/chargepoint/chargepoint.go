package chargepoint

import (
	"context"

	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
	repository "github.com/successgilli/api-app/internal/storage/repositories/chargepoint"
)

// Service is a charge point service
type Service struct {
	repo repository.Repository
}

// NewService initializes a new charge point service
func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

// Create crates a new charge point
func (s *Service) Create(ctx context.Context, cp entity.NewChargePoint) (*entity.ChargePoint, error) {
	createdChargePoint, err := s.repo.CreateOne(ctx, cp)
	if err != nil {
		return nil, err
	}

	return createdChargePoint, nil
}
