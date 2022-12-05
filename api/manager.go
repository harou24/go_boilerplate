package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Manager struct {
	router *fiber.App
}

func NewManager() *Manager {
	return &Manager{
		router: fiber.New(),
	}
}

func (m *Manager) Serve() {

	m.router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	m.router.Get("/", func(c *fiber.Ctx) error {
		data := map[string]any{
			"message": "Hello World 🖖",
		}
		return c.JSON(data)
	})

	log.Fatal(m.router.Listen(":5000"))
}
