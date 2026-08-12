package chargepoint

// createChargePointQuery inserts a new charge point.
// Params: name, latitude, longitude, status
const createChargePointQuery = `
INSERT INTO charge_points (id, name, latitude, longitude, status) VALUES (?, ?, ?, ?, ?)
`

// getChargePointByIDQuery retrieves a single charge point by its ID.
// Params: id
const getChargePointByIDQuery = `
SELECT id, name, latitude, longitude, status FROM charge_points WHERE id = ?
`

// findNearbyQuery finds all charge points within radiusKm of (lat, lon),
// using an indexed lat/lon bounding box as a pre-filter, then an exact
// haversine (spherical law of cosines) distance calculation.
// Params (in order): lat, lon, lat, minLat, maxLat, minLon, maxLon, radiusKm
const findNearbyChargePointsQuery = `
	SELECT id, name, latitude, longitude, status FROM (
		SELECT
			id,
			name,
			latitude,
			longitude,
			status,
			(
				6371 * acos(
					max(-1.0, min(1.0,
						cos(radians(?)) * cos(radians(latitude)) *
						cos(radians(longitude) - radians(?)) +
						sin(radians(?)) * sin(radians(latitude))
					))
				)
			) AS distance
		FROM charge_points
		WHERE latitude BETWEEN ? AND ?
		AND longitude BETWEEN ? AND ?
	) AS nearby
	WHERE distance <= ?
	ORDER BY distance;
	`
