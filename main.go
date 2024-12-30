package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Post("/send", func(c *fiber.Ctx) error {
		data := struct {
			Message string `json:"message"`
		}{}
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).SendString("Invalid request")
		}

		err := produceToKafka(data.Message)
		if err != nil {
			return c.Status(500).SendString("Failed to send message")
		}

		return c.SendString("Message sent to Kafka")
	})

	log.Fatal(app.Listen(":3000"))
}
