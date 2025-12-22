package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heshanthenura/sentrigov2/internal/capture"
)

func StartServer() {
	r := gin.Default()

	r.GET("/api/interfaces", func(c *gin.Context) {
		ifaces, err := capture.ListInterfaces()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, ifaces)
	})

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/api/ws/metrics", MetricsWS)

	r.Static("/static", "./frontend/dist")
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	log.Println("Starting HTTP server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
