package chargepoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
)

func TestChargePointStatus_IsValid(t *testing.T) {
	tests := []struct {
		name   string
		status entity.ChargePointStatus
		want   bool
	}{
		{
			name:   "available is valid",
			status: entity.AvailableStatus,
			want:   true,
		},
		{
			name:   "occupied is valid",
			status: entity.OccupiedStatus,
			want:   true,
		},
		{
			name:   "offline is valid",
			status: entity.OfflineStatus,
			want:   true,
		},
		{
			name:   "empty string is invalid",
			status: entity.ChargePointStatus(""),
			want:   false,
		},
		{
			name:   "unknown value is invalid",
			status: entity.ChargePointStatus("BROKEN"),
			want:   false,
		},
		{
			name:   "lowercase valid value is invalid",
			status: entity.ChargePointStatus("available"),
			want:   false,
		},
		{
			name:   "value with extra whitespace is invalid",
			status: entity.ChargePointStatus("AVAILABLE "),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.status.IsValid()
			assert.Equal(t, tt.want, got)
		})
	}
}
