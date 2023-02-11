package main

import (
	"github.com/gofiber/fiber/v2"
)

type Food struct {
	Name string `json:"name"`
	Price int `json:"price"`
}

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

	// Query Parser
	app.Get("/food", func(c *fiber.Ctx) error {
		food := Food{};
		c.QueryParser(&food);
		return c.JSON(food); 
	})
	
	// WildCard
	app.Get("/wild/*", func(c *fiber.Ctx) error {
		wildcard := c.Params("*");
		return c.SendString(wildcard); 
	})

	// Static File
	app.Static("/web", "./ui", fiber.Static{
		Index: "index.html",
	})

	app.Listen(":8009"); 

}