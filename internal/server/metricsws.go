package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/heshanthenura/sentrigov2/internal/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func MetricsWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(500, "Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer conn.Close()
	for {
		metrics, err := utils.GetMetrics()
		if err != nil {
			log.Println("Failed to get metrics:", err)
			break
		}

		if err := conn.WriteJSON(metrics); err != nil {
			log.Println("WebSocket write error:", err)
			break
		}

		time.Sleep(3 * time.Second)
	}
}
