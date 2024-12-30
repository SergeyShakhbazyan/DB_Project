package app

import (
	"ProjectDB/config"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	Router   *gin.Engine
	Database *config.Connection
}

func (a *App) Initialize() {
	a.Router = gin.Default()
	a.Database = config.NewConnection()
}

func (a *App) Run(port string) {
	if err := a.Router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
