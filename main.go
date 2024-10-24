package main

import (
	"log"
	"telegram-chatbot/clients"
	"telegram-chatbot/config"
	"telegram-chatbot/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	bot := clients.Init()
	handlers.Init(bot)
	log.Println("Server starting on port:", config.Config("PORT"))
	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
