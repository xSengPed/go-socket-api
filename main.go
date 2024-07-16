package main

import (
	"context"
	socketmessage "go-socket-api/pkg/socket_message"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// NewHttpHandler is a constructor for the http handler
func NewHttpHandler(lc fx.Lifecycle, messageHandler *socketmessage.Handler) *fiber.App {

	app := fiber.New()
	api := app.Group("/api")
	v := api.Group("/v1")

	/* Mount for service path*/
	v.Mount("/socket-message", messageHandler.App)

	/* lc (LifeCycle)  create lc hook*/
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			/* use go routine for listening on port 3000 */
			go app.Listen(":3000")
			return nil
		},
		OnStop: func(context.Context) error {
			app.Shutdown()
			return nil
		},
	})

	return app
}
func main() {
	fx.New(
		/* Provide handler for using for dependency injection (DI)*/
		fx.Provide(socketmessage.NewHandler),
		/* Invoke NewHttpHandler for create fiber app instance then run the api */
		fx.Invoke(NewHttpHandler),
	).Run()
}
