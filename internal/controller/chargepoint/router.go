package chargepoint

import (
	"database/sql"
	"github.com/gin-gonic/gin"

	service "github.com/successgilli/api-app/internal/service/chargepoint"
	repository "github.com/successgilli/api-app/internal/storage/repositories/chargepoint"
)

func RegisterRoutes(r *gin.RouterGroup, sqlDB *sql.DB) {
	repo := repository.NewSQLiteRepository(sqlDB)
	svc := service.NewService(repo)
	ctr := NewController(svc)

	api := r.Group("/chargepoints")

	api.POST("/", ctr.Create)
	api.GET("/:id", ctr.GetByID)
}
