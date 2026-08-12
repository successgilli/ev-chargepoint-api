package chargepoint_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
	svc "github.com/successgilli/api-app/internal/service/chargepoint"
	repo "github.com/successgilli/api-app/internal/storage/repositories/chargepoint"
)

func TestService_Create_Success(t *testing.T) {
	mockRepo := repo.NewMockRepository(t)

	uid, err := uuid.NewV7()
	require.NoError(t, err)

	cmd := svc.CreateChargePointCommand{
		Name:   "Station A",
		Status: entity.AvailableStatus,
	}
	repoParam := repo.CreateChargePointParam{
		Name:   "Station A",
		Status: entity.AvailableStatus,
	}
	expected := &entity.ChargePoint{
		ID:     uid.String(),
		Name:   "Station A",
		Status: entity.AvailableStatus,
	}

	mockRepo.EXPECT().
		CreateOne(t.Context(), repoParam).
		Return(expected, nil)

	service := svc.NewService(mockRepo)

	result, err := service.Create(t.Context(), cmd)

	require.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestService_Create_RepositoryError(t *testing.T) {
	mockRepo := repo.NewMockRepository(t)

	cmd := svc.CreateChargePointCommand{Name: "Station A"}
	repoParam := repo.CreateChargePointParam{Name: "Station A"}
	repoErr := errors.New("db connection failed")

	mockRepo.EXPECT().
		CreateOne(t.Context(), repoParam).
		Return(nil, repoErr)

	service := svc.NewService(mockRepo)

	result, err := service.Create(t.Context(), cmd)

	assert.Nil(t, result)
	assert.ErrorIs(t, err, repoErr)
}
