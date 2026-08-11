package chargepoint

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"

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
	uid, err := uuid.NewV7()
	if err != nil {
		log.Println("error generating uuid: ", err)
		return nil, apperrors.ErrInternal
	}

	chargePointID := uid.String()

	_, dbErr := c.db.ExecContext(ctx,
		`INSERT INTO charge_points (id, name, latitude, longitude, status) VALUES (?, ?, ?, ?, ?)`,
		chargePointID, cp.Name, cp.Lat, cp.Long, cp.Status,
	)
	if dbErr != nil {
		log.Println("charge point insert error: ", dbErr)
		return nil, apperrors.ErrInternal
	}

	return &entity.ChargePoint{
		ID:   chargePointID,
		Name: cp.Name,
		Lat:  cp.Lat,
		Long: cp.Long,
	}, nil
}

// GetByID Gets a charge point by ID
func (c *ChargePointSQLiteRepository) GetByID(ctx context.Context, id string) (*entity.ChargePoint, error) {
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

		log.Println("charge point db scan error: ", err)
		return nil, apperrors.ErrInternal
	}

	return &cp, nil
}
