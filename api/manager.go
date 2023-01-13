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
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	router.Get("/", func(c *fiber.Ctx) error {
		data := map[string]interface{}{
			"message": "Hello World 🖖",
		}
		return c.JSON(data)
	})

	return &Manager{
		router: router,
	}
}

func (m *Manager) Serve(port string) {
	log.Fatal(m.router.Listen(":" + port))
}

func (m *Manager) Close() error {
	return m.router.Shutdown()
}
