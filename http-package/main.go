package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// GET http://localhost:3000/user/asfasfasl/fasfasklşfka
	// reverse proxy: http://localhost:3000/user/asfasfasl/fasfasklşfka -> http://localhost:3000/user/*
	app.Get("/:key/*", ProxyHandler)

	app.Listen(":3000")
}
