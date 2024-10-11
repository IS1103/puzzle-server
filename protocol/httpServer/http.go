package httpServer

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/antage/eventsource.v1"
)

func StartServer(port string) {
	es := eventsource.New(&eventsource.Settings{
		Timeout:        5 * time.Second,
		CloseOnTimeout: false,
		IdleTimeout:    30 * time.Minute,
	}, nil)
	defer es.Close()
	r := gin.Default()

	{
		r.GET("/events", esSSE)
		r.Run(":" + port)
	}
}

func esSSE(c *gin.Context) {
	w := c.Writer

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)

	if !ok {
		log.Panic("server not support") //浏览器不兼容
	}

	// 使用循環定期推送數據
	for i := 0; i < 10; i++ { // 可以改為無限循環 `for { ... }`
		message := fmt.Sprintf("Message %d", i+1)
		_, err := fmt.Fprintf(w, "data: %s\n\n", message)
		if err != nil {
			return
		}

		flusher.Flush()             // 立即刷新緩衝區，將數據發送到客戶端
		time.Sleep(1 * time.Second) // 模擬推送間隔
	}

	// _, err := fmt.Fprintf(w, "data: %s\n\n", "dsdf")
	// if err != nil {
	// 	return
	// }
}
