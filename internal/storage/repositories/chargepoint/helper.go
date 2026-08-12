package chargepoint

import "math"


type boundingBox struct {
	MinLat float64
	MaxLat float64
	MinLon float64
	MaxLon float64
}

// calculateBoundingBox computes the coordinate limits for a given radius in km
func calculateBoundingBox(lat, lon, radiusKm float64) boundingBox {
	const kmPerDegreeLat = 111.0

	// 1. Latitude delta remains constant everywhere on Earth
	deltaLat := radiusKm / kmPerDegreeLat

	// 2. Longitude delta narrows towards the poles, apply cosine adjustment
	// Convert latitude to radians first
	radLat := lat * math.Pi / 180
	deltaLon := radiusKm / (kmPerDegreeLat * math.Cos(radLat))

	return boundingBox{
		MinLat: lat - deltaLat,
		MaxLat: lat + deltaLat,
		MinLon: lon - deltaLon,
		MaxLon: lon + deltaLon,
	}
}
