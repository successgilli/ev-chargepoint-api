package chargepoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateBoundingBox_Properties(t *testing.T) {
	tests := []struct {
		name     string
		lat, lon float64
		radiusKm float64
	}{
		{"amsterdam, 10km", 52.3676, 4.9041, 10},
		{"equator, 100km", 0, 0, 100},
		{"southern hemisphere, 50km", -33.8688, 151.2093, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := calculateBoundingBox(tt.lat, tt.lon, tt.radiusKm)

			// the box must actually contain the center point
			assert.LessOrEqual(t, box.MinLat, tt.lat)
			assert.GreaterOrEqual(t, box.MaxLat, tt.lat)
			assert.LessOrEqual(t, box.MinLon, tt.lon)
			assert.GreaterOrEqual(t, box.MaxLon, tt.lon)

			// larger radius must always produce a larger (or equal) box
			biggerBox := calculateBoundingBox(tt.lat, tt.lon, tt.radiusKm*2)
			assert.Less(t, biggerBox.MinLat, box.MinLat)
			assert.Greater(t, biggerBox.MaxLat, box.MaxLat)
		})
	}
}
