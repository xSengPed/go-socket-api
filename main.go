package main

import (
	"context"
	"go-socket-api/configs"
	"go-socket-api/pkg/members"
	socketmessage "go-socket-api/pkg/socket_message"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// NewHttpHandler is a constructor for the http handler
func NewHttpHandler(lc fx.Lifecycle, config *configs.Config, db *gorm.DB, membersHandler *members.MemberHandler, messageHandler *socketmessage.Handler) *fiber.App {

	db.AutoMigrate(&members.Member{})
	app := fiber.New()
	api := app.Group("/api")
	v := api.Group("/v1")
	ws := v.Group("/ws")

	v.Mount("/member", membersHandler.App)
	/* Mount for service path*/
	ws.Mount("/socket-message", messageHandler.App)

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
		fx.Provide(configs.New, configs.NewDatabaseConnect, members.NewService, members.NewMemberHandler, socketmessage.NewHandler),
		/* Invoke NewHttpHandler for create fiber app instance then run the api */
		fx.Invoke(NewHttpHandler),
	).Run()
}
