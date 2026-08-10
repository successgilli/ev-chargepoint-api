package chargepoint

import (
	"context"
	"database/sql"
	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
)

type ChargePointSQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *ChargePointSQLiteRepository {
	return &ChargePointSQLiteRepository{
		db: db,
	}
}

// CreateOne creates one charge point
func (c *ChargePointSQLiteRepository) CreateOne(ctx context.Context, cp entity.NewChargePoint) (*entity.ChargePoint, error) {
	result, err := c.db.ExecContext(ctx,
		`INSERT INTO charge_points (name, latitude, longitude, status) VALUES (?, ?, ?, ?)`,
		cp.Name, cp.Lat, cp.Long, cp.Status,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &entity.ChargePoint{
		ID:   id,
		Name: cp.Name,
		Lat:  cp.Lat,
		Long: cp.Long,
	}, nil
}
