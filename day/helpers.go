package day

import (
	"context"
	"fmt"
	"strconv"
	"weather-predictor/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// Función encargada de disminuir la duplicidad de código procesando los query params que consisten de velocidades angulares y manejando errores.
// Parámetros: El contexto.
func ParseAngularParams(c *fiber.Ctx) (*int, *int, *int, *string) {
	ferengi := c.Query("ferengi", "1")
	ferengi_angular, err := strconv.Atoi(ferengi)
	if err != nil {
		fmt.Println("Error:", err)
		error_description := "La velocidad angular de Ferengi tiene un parámetro inválido."
		return nil, nil, nil, &error_description
	}

	vulcano := c.Query("vulcano", "-5")
	vulcano_angular, err := strconv.Atoi(vulcano)
	if err != nil {
		fmt.Println("Error:", err)
		error_description := "La velocidad angular de vulcano tiene un parámetro inválido."
		return nil, nil, nil, &error_description
	}

	betazoide := c.Query("betazoide", "3")
	betazoide_angular, err := strconv.Atoi(betazoide)
	if err != nil {
		fmt.Println("Error:", err)
		error_description := "La velocidad angular de betazoide tiene un parámetro inválido."
		return nil, nil, nil, &error_description
	}
	return &ferengi_angular, &vulcano_angular, &betazoide_angular, nil
}

// Función encargada de disminuir la duplicidad de código procesando los query params que consisten de velocidades angulares y radios y manejando errores.
// Parámetros: El contexto.
func ParseAngularRadiusParams(c *fiber.Ctx) (*int, *int, *int, *int, *int, *int, *string) {
	ferengi := c.Query("ferengi_a", "1")
	ferengi_angular, err := strconv.Atoi(ferengi)
	if err != nil {
		fmt.Println("Error:", err)
		error_description := "La velocidad angular de Ferengi tiene un parámetro inválido."
		return nil, nil, nil, nil, nil, nil, &error_description
	}

	ferengi_rad := c.Query("ferengi_r", "1")
	ferengi_radius, err := strconv.Atoi(ferengi_rad)
	if err != nil {
		fmt.Println("Error:", err)
		error_description := "El radio de Ferengi tiene un parámetro inválido."
		return nil, nil, nil, nil, nil, nil, &error_description
	}

	vulcano := c.Query("vulcano_a", "-5")
	vulcano_angular, err := strconv.Atoi(vulcano)
	if err != nil {
		fmt.Println("Error:", err)
		error_description := "La velocidad angular de vulcano tiene un parámetro inválido."
		return nil, nil, nil, nil, nil, nil, &error_description
	}

	vulcano_rad := c.Query("vulcano_r", "-5")
	vulcano_radius, err := strconv.Atoi(vulcano_rad)
	if err != nil {
		fmt.Println("Error:", err)
		error_description := "El radio de vulcano tiene un parámetro inválido."
		return nil, nil, nil, nil, nil, nil, &error_description
	}

	betazoide := c.Query("betazoide_a", "3")
	betazoide_angular, err := strconv.Atoi(betazoide)
	if err != nil {
		fmt.Println("Error:", err)
		error_description := "La velocidad angular de betazoide tiene un parámetro inválido."
		return nil, nil, nil, nil, nil, nil, &error_description
	}

	betazoide_rad := c.Query("betazoide_r", "3")
	betazoide_radius, err := strconv.Atoi(betazoide_rad)
	if err != nil {
		fmt.Println("Error:", err)
		error_description := "El radio de betazoide tiene un parámetro inválido."
		return nil, nil, nil, nil, nil, nil, &error_description
	}

	return &ferengi_angular, &ferengi_radius, &vulcano_angular, &vulcano_radius, &betazoide_angular, &betazoide_radius, nil
}

// Función encargada de popular la base de datos de forma iterativa.
// Parámetros: Referencia a la colección de la base de datos, velocidades angulares y radios.
func PopulateDB(daysCollection *mongo.Collection, a_ang, a_dis, b_ang, b_dis, c_ang, c_dis int) *string {
	a_position, b_position, c_position := 0, 0, 0
	for i := 0; i < utils.DAYS; i++ {
		day_status := "Normal"
		rain_amount := 0.0
		p1 := utils.Rad2Cart(float64(a_dis), float64(a_position))
		p2 := utils.Rad2Cart(float64(b_dis), float64(b_position))
		p3 := utils.Rad2Cart(float64(c_dis), float64(c_position))
		if utils.SunContained(p1, p2, p3) {
			day_status = "Rain"
			rain_amount = utils.TrianglePerimeter(p1, p2, p3)
		} else if (a_position%180) == (b_position%180) && (a_position%180) == (c_position%180) {
			day_status = "Drought"
		} else if utils.CheckLine(p1, p2, p3) {
			day_status = "Optimal"
		}

		day := Day{
			Year:           (i / 365) + 1,
			Day:            (i % 365) + 1,
			Status:         day_status,
			RainAmount:     rain_amount,
			FerengiAngle:   a_position,
			VulcanoAngle:   b_position,
			BetazoideAngle: c_position,
		}
		results, err := daysCollection.InsertOne(context.TODO(), day)
		_ = results
		if err != nil {
			fmt.Println(err)
			error_description := "Error al guardar los días en base de datos."
			return &error_description
		}
		a_position = (a_position + a_ang + 360) % 360
		b_position = (b_position + b_ang + 360) % 360
		c_position = (c_position + c_ang + 360) % 360
	}
	return nil
}
