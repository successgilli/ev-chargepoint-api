package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/successgilli/api-app/internal/controller/chargepoint"
)

func RegisterRoutes(r *gin.RouterGroup, sqlDB *sql.DB) {
	chargepoint.RegisterRoutes(r, sqlDB)
}

func RegisterCustomValidators() error {
	return chargepoint.RegisterCustomValidators()
}
