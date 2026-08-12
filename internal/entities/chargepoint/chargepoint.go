package chargepoint

type ChargePointStatus string

const (
	AvailableStatus ChargePointStatus = "AVAILABLE"
	OccupiedStatus  ChargePointStatus = "OCCUPIED"
	OfflineStatus   ChargePointStatus = "OFFLINE"
)

// ChargePoint is a charge point entity
type ChargePoint struct {
	ID     string
	Name   string
	Lat    float64
	Long   float64
	Status ChargePointStatus
}

// IsValid checks if a status is valid
func (s ChargePointStatus) IsValid() bool {
	switch s {
	case AvailableStatus, OccupiedStatus, OfflineStatus:
		return true
	default:
		return false
	}
}
