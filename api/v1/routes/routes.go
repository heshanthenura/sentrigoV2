package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/heshanthenura/sentrigov2/api/v1/handlers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	api.GET("/interfaces", handlers.GetAllInterfaces)

	api.POST("/capture/start", handlers.StartCapture)
	api.POST("/capture/stop", handlers.StopCapture)

}
