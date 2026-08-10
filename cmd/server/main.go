package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/successgilli/api-app/internal/controller"
	"github.com/successgilli/api-app/internal/storage/sqlite"

	apiDocs "github.com/successgilli/api-app/api"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// baseAPI is the base API
// TODO: move to config
const baseAPI = "/api/v1"
const dbPath = "ev.db"

func main() {
	r := gin.Default()

	db, err := sqlite.Init(dbPath)
	if err != nil {
		panic(fmt.Errorf("error initializing DB %w", err))
	}

	registerRoutes(r, db)

	r.Run(":8080")
}

// registerRoutes registers routes
func registerRoutes(r *gin.Engine, db *sql.DB) {
	api := r.Group(baseAPI)

	controller.RegisterRoutes(api, db)
	registerSwagger(api)
}

// registerSwagger registers swagger
// TODO: map to only dev
func registerSwagger(r *gin.RouterGroup) {
	apiDocs.SwaggerInfo.BasePath = baseAPI
	apiDocs.SwaggerInfo.Version = "v0.1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
