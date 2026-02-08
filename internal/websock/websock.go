package websock

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/heshanthenura/sentrigov2/internal/config"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex

const (
	pongWait   = 10 * time.Second
	pingPeriod = 5 * time.Second
)

func WsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	log.Println("new user connected")
	log.Println("total clients:", len(clients))
	// _ = conn.WriteMessage(websocket.TextMessage, []byte("welcome"))
	_ = conn.WriteJSON(config.GetConfig())

	var writeMu sync.Mutex

	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
		log.Println("user disconnected")
	}()

	conn.SetReadLimit(1024 * 1024)

	_ = conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		_ = conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	go func() {
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}()

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for range ticker.C {
		writeMu.Lock()
		_ = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		err := conn.WriteControl(
			websocket.PingMessage,
			nil,
			time.Now().Add(10*time.Second),
		)
		writeMu.Unlock()

		if err != nil {
			return
		}
	}
}
