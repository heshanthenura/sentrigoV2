package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heshanthenura/sentrigov2/api/v1/routes"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
