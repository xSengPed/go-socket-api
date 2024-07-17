package members

import "github.com/gofiber/fiber/v2"

type MemberHandler struct {
	App *fiber.App
}

func NewMemberHandler(s *Service) *MemberHandler {

	app := fiber.New()
	app.Get("/create", func(c *fiber.Ctx) error {
		err := s.Create()
		if err != nil {
			return c.SendString("Create member Failed")
		}
		return c.SendString("Create member")
	})
	return &MemberHandler{
		App: app,
	}
}
