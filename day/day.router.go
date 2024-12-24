package day

import (
	"context"
	"fmt"
	"strconv"
	"weather-predictor/config/db"
	"weather-predictor/config/envs"
	"weather-predictor/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Route(app *fiber.App) {

	daysCollection := db.Client.Database(envs.EnvVariable("CUR_DB")).Collection("days")
	day := app.Group("/day")

	//Handler que realiza un hello world.
	day.Get("/hello-world", func(c *fiber.Ctx) error {
		fmt.Println("Get, Hello World")

		return c.Status(fiber.StatusOK).JSON("Esto es un hola mundo del predictor de climas.")
	})

	//Handler encargado de retornar el número de sequías calculado de forma iterativa.
	//Parámetros: Velocidades angulares de los planetas enviados como query params.
	day.Get("/drought-iterative", func(c *fiber.Ctx) error {

		ferengi, vulcano, betazoide, err := ParseAngularParams(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": *err})
		}
		fmt.Println("Get drought days iterative. Parameters: ", *ferengi, *vulcano, *betazoide)
		var drought_days = utils.DroughtDaysIterative(*ferengi, *vulcano, *betazoide)
		response := map[string]interface{}{
			"message":      "Total de días de sequía calculado de forma iterativa dados los parámetros iniciales del problema.",
			"drought_days": drought_days,
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	//Handler encargado de retornar el número de sequías calculado de forma matemática.
	//Parámetros: Velocidades angulares de los planetas enviados como query params.
	day.Get("/drought-congruence", func(c *fiber.Ctx) error {

		ferengi, vulcano, betazoide, err := ParseAngularParams(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": *err})
		}

		fmt.Println("Get drought days congruence. Parameters: ", *ferengi, *vulcano, *betazoide)

		var drought_days = utils.DroughtDays(*ferengi, *vulcano, *betazoide)
		response := map[string]interface{}{
			"message":      "Total de días de sequía calculado de forma matemática y general dados los parámetros iniciales del problema.",
			"drought_days": drought_days,
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	//Handler encargado de retornar el número de días lluviosos y el día mas lluvioso.
	//Parámetros: Velocidades angulares y radios de los planetas enviados como query params.
	day.Get("/rain", func(c *fiber.Ctx) error {
		fmt.Println("Get rainy days")
		ferengi_angular, ferengi_radius, vulcano_angular, vulcano_radius, betazoide_angular, betazoide_radius, err := ParseAngularRadiusParams(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": *err})
		}
		var rainy_days, rainiest_day = utils.RainyDays(*ferengi_angular, *ferengi_radius, *vulcano_angular, *vulcano_radius, *betazoide_angular, *betazoide_radius)
		response := map[string]interface{}{
			"message":      "Total de días de lluvia. Y dia mas lluvioso, el dia mas lluvioso se repite cada 360 dias.",
			"rainy_days":   rainy_days,
			"rainiest_day": rainiest_day,
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	//Handler encargado de retornar el número de días óptimos.
	//Parámetros: Velocidades angulares y radios de los planetas enviados como query params.
	day.Get("/optimal", func(c *fiber.Ctx) error {
		fmt.Println("Get optimal days")
		ferengi_angular, ferengi_radius, vulcano_angular, vulcano_radius, betazoide_angular, betazoide_radius, err := ParseAngularRadiusParams(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": *err})
		}
		var optimal_days = utils.OptimalDays(*ferengi_angular, *ferengi_radius, *vulcano_angular, *vulcano_radius, *betazoide_angular, *betazoide_radius)
		response := map[string]interface{}{
			"message":      "Total de días optimos.",
			"optimal_days": optimal_days,
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	//Handler encargado de popular la base de datos de acuerdo a unas velocidades angulares y radios dados.
	////Parámetros: Velocidades angulares y radios de los planetas enviados como query params.
	day.Post("/populate", func(c *fiber.Ctx) error {
		fmt.Println("Database Population")
		ferengi_angular, ferengi_radius, vulcano_angular, vulcano_radius, betazoide_angular, betazoide_radius, err := ParseAngularRadiusParams(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": *err})
		}

		err = PopulateDB(daysCollection, *ferengi_angular, *ferengi_radius, *vulcano_angular, *vulcano_radius, *betazoide_angular, *betazoide_radius)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": *err})
		}
		response := map[string]interface{}{
			"message": "Base de datos populada con éxito.",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	//Función encargada de eliminar la información de la base de datos para poder popularla posteriormente con distintas entradas.
	day.Delete("/empty", func(c *fiber.Ctx) error {
		fmt.Println("Empty Database")

		_, err := daysCollection.DeleteMany(context.TODO(), bson.D{})
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error borrando la base de datos."})
		}

		response := map[string]interface{}{
			"message": "Base de datos borrada con éxito.",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	//Handler encargado de retornar la información de un dia específico dado un día y un año.
	//Parámetros: Día y año enviados como query params.
	day.Get("/info", func(c *fiber.Ctx) error {
		fmt.Println("Retrieve day with year and day")

		year := c.Query("year", "1")
		search_day := c.Query("day", "1")

		year_value, err := strconv.Atoi(year)
		if err != nil {
			fmt.Println("Error:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "El valor del año tiene un parámetro inválido."})
		}

		day_value, err := strconv.Atoi(search_day)
		if err != nil {
			fmt.Println("Error:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "El valor del día tiene un parámetro inválido."})
		}

		response, err := daysCollection.Find(context.TODO(), bson.M{"year": year_value, "day": day_value})

		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al recuperar la información de la base de datos."})
		}

		var result []bson.M

		if err = response.All(context.TODO(), &result); err != nil {
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(result)
	})

	//Función encargada de retornar todos los días cuyo estado coincida con el parámetro dado.
	//Párametros: Posible estado del día [Rain,Normal,Drought,Optimal].
	day.Get("/info/status", func(c *fiber.Ctx) error {
		fmt.Println("Retrieve day with status")

		status := c.Query("status", "Rain")

		response, err := daysCollection.Find(context.TODO(), bson.D{{"status", status}})
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al recuperar la información de la base de datos."})
		}

		var result []bson.M

		if err = response.All(context.TODO(), &result); err != nil {
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(result)
	})
}
