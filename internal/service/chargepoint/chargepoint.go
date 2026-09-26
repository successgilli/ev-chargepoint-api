package chargepoint

import (
	"context"
	"fmt"

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
func (s *Service) Create(ctx context.Context, cmd CreateChargePointCommand) (*entity.ChargePoint, error) {
	createdChargePoint, err := s.repo.CreateOne(ctx, repository.CreateChargePointParam{
		Name:   cmd.Name,
		Lat:    cmd.Lat,
		Long:   cmd.Long,
		Status: cmd.Status,
	})
	if err != nil {
		return nil, err
	}

	return createdChargePoint, nil
}

// GetByID gets a charge point by ID
func (s *Service) GetByID(ctx context.Context, id string) (*entity.ChargePoint, error) {
	createdChargePoint, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return createdChargePoint, nil
}

// GetAllNearBy gets all charge points nearby
func (s *Service) GetAllNearBy(ctx context.Context, cmd FindNearByCommand) ([]entity.ChargePoint, error) {
	cps, err := s.repo.GetAllNearBy(ctx, repository.FindNearByParam{
		Latitude:  cmd.Latitude,
		Longitude: cmd.Longitude,
		RadiusKM:  cmd.RadiusKM,
	})
	if err != nil {
		return nil, err
	}

	return cps, nil
}

// GetByID gets a charge point by ID
func (s *Service) GetByID(ctx context.Context, id string) (*entity.ChargePoint, error) {
	createdChargePoint, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	fmt.Sprintln("Armpit")

	return createdChargePoint, nil
}