package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New();

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GET: Hello Fiber.");
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("POST: Hello Fiber.");
	})

	// Parameter and Parameter Optional. 
	app.Get("/hello/:name/:surname?", func(c *fiber.Ctx) error {
		name := c.Params("name");
		surname := c.Params("surname");
		return c.SendString("Name: " + name + " " + surname);
	})

	// Query 
	app.Get("/hi", func(c *fiber.Ctx) error {
		name := c.Query("name");
		surname := c.Query("surname");
		return c.SendString("Hi " + name + " " + surname); 
	})

	app.Listen(":8009"); 
}