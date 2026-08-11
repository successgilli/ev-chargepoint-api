package chargepoint_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/google/uuid"

	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
	svc "github.com/successgilli/api-app/internal/service/chargepoint"
	repo "github.com/successgilli/api-app/internal/storage/repositories/chargepoint"
)

func TestService_Create_Success(t *testing.T) {
	mockRepo := repo.NewMockRepository(t)

	uid, err := uuid.NewV7()
	require.NoError(t, err)

	input := entity.NewChargePoint{
		Name:   "Station A",
		Status: entity.AvailableStatus,
	}
	expected := &entity.ChargePoint{
		ID:     uid.String(),
		Name:   "Station A",
		Status: entity.AvailableStatus,
	}

	mockRepo.EXPECT().
		CreateOne(t.Context(), input).
		Return(expected, nil)

	service := svc.NewService(mockRepo)

	result, err := service.Create(t.Context(), input)

	require.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestService_Create_RepositoryError(t *testing.T) {
	mockRepo := repo.NewMockRepository(t)

	input := entity.NewChargePoint{Name: "Station A"}
	repoErr := errors.New("db connection failed")

	mockRepo.EXPECT().
		CreateOne(t.Context(), input).
		Return(nil, repoErr)

	service := svc.NewService(mockRepo)

	result, err := service.Create(t.Context(), input)

	assert.Nil(t, result)
	assert.ErrorIs(t, err, repoErr)
}
