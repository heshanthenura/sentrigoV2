package api

import (
	"net/http"
	"sentrigoV2/engine/internal/api/ui"

	"github.com/gin-gonic/gin"
)

func StartAPI() error {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	files, err := ui.FS()
	if err != nil {
		return err
	}

	router.NoRoute(gin.WrapH(
		http.FileServer(http.FS(files)),
	))

	return router.Run(":8080")
}
