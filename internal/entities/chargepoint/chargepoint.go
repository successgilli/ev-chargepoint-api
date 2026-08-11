package chargepoint

type ChargePointStatus string

const (
	AvailableStatus ChargePointStatus = "AVAILABLE"
	OccupiedStatus  ChargePointStatus = "OCCUPIED"
	OfflineStatus   ChargePointStatus = "OFFLINE"
)

type NewChargePoint struct {
	Name   string
	Lat    float64
	Long   float64
	Status ChargePointStatus
}

type ChargePoint struct {
	ID     string
	Name   string
	Lat    float64
	Long   float64
	Status ChargePointStatus
}
