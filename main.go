package main

import (
	"weather-predictor/config/db"
	"weather-predictor/config/envs"
	"weather-predictor/day"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db.Initdb()
	day.Route(app)

	app.Listen(":" + envs.EnvVariable("PORT"))
}
