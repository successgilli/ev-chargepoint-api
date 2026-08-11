package chargepoint

import (
	"context"
	"database/sql"
	"errors"

	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
	"github.com/successgilli/api-app/internal/pkg/apperrors"
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

// GetByID Gets a charge point by ID
func (c *ChargePointSQLiteRepository) GetByID(ctx context.Context, id int) (*entity.ChargePoint, error) {
	row := c.db.QueryRowContext(ctx,
		"SELECT id, name, latitude, longitude, status FROM charge_points WHERE id = ?",
		id,
	)

	var cp entity.ChargePoint

	err := row.Scan(&cp.ID, &cp.Name, &cp.Lat, &cp.Long, &cp.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}

	return &cp, nil
}
