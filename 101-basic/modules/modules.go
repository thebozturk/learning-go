package main

import (
	"github.com/gofiber/fiber/v2"
)
// go mod init ile go modülü oluşturabiliriz ve go modülü içerisindeki paketleri go get ile indirebiliriz

// fiber framework'ü ile web uygulamaları geliştirebiliriz
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}