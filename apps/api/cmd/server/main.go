package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	log.Println("CafeCounter Api running in port:8080")

	log.Fatal(app.Listen(":8080"))
}