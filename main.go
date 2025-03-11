package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var (
	choices []int = []int{0, 1, 1, 1}
)

func main() {
	server := fiber.New()
	server.Get("/metrics", monitor.New())
	server.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	server.Get("/json", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world",
		})
	})
	server.Post("/create", func(c *fiber.Ctx) error {
		if c.Get("Content-Type") != "application/json" {
			return c.SendStatus(http.StatusBadRequest)
		}

		success := choices[rand.Intn(len(choices))]
		if success < 1 {
			return c.SendStatus(http.StatusBadRequest)
		}

		payload := map[string]any{}
		err := c.BodyParser(&payload)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}
		return c.SendStatus(http.StatusNoContent)
	})
	server.Listen(":3000")
}
