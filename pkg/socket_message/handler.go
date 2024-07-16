package socketmessage

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

type Handler struct {
	App *fiber.App
}

func initWebsocket(c *websocket.Conn) {
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			break
		}

		slog.Info("Message received: %s", msg)

		err = c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Hello %s", msg)))
		if err != nil {
			break
		}

	}
}

func NewHandler() *Handler {
	app := fiber.New()
	app.Get("/", websocket.New(initWebsocket))
	return &Handler{
		App: app,
	}
}
