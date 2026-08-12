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
func (c *ChargePointSQLiteRepository) CreateOne(ctx context.Context, param CreateChargePointParam) (*entity.ChargePoint, error) {
	uid, err := uuid.NewV7()
	if err != nil {
		log.Println("error generating uuid: ", err)
		return nil, apperrors.ErrInternal
	}

	chargePointID := uid.String()

	_, dbErr := c.db.ExecContext(ctx, createChargePointQuery, chargePointID, param.Name, param.Lat, param.Long, param.Status)
	if dbErr != nil {
		log.Println("charge point insert error: ", dbErr)
		return nil, apperrors.ErrInternal
	}

	return &entity.ChargePoint{
		ID:     chargePointID,
		Name:   param.Name,
		Lat:    param.Lat,
		Long:   param.Long,
		Status: param.Status,
	}, nil
}

// GetByID Gets a charge point by ID
func (c *ChargePointSQLiteRepository) GetByID(ctx context.Context, id string) (*entity.ChargePoint, error) {
	row := c.db.QueryRowContext(ctx, getChargePointByIDQuery, id)

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

// GetAllNearBy Gets all charge points near by
func (c *ChargePointSQLiteRepository) GetAllNearBy(ctx context.Context, param FindNearByParam) ([]entity.ChargePoint, error) {
	boundingBox := calculateBoundingBox(param.Latitude, param.Longitude, param.RadiusKM)

	results, err := c.db.QueryContext(
		ctx,
		findNearbyChargePointsQuery,
		param.Latitude,
		param.Longitude,
		param.Latitude,
		boundingBox.MinLat,
		boundingBox.MaxLat,
		boundingBox.MinLon,
		boundingBox.MaxLon,
		param.RadiusKM,
	)
	if err != nil {
		log.Println("error fetching nearby query: ", err)
		return nil, apperrors.ErrInternal
	}
	defer func() {
		if closeErr := results.Close(); closeErr != nil {
			log.Println("error closing query: ", closeErr)
		}
	}()

	chargePoints := []entity.ChargePoint{}

	for results.Next() {
		var cp entity.ChargePoint

		if err := results.Scan(&cp.ID, &cp.Name, &cp.Lat, &cp.Long, &cp.Status); err != nil {
			log.Println("error fetching nearby query: ", err)
			break
		}

		chargePoints = append(chargePoints, cp)
	}

	if results.Err() != nil {
		log.Println("row error fetching nearby query: ", err)
	}

	return chargePoints, nil
}
