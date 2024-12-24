package utils

// Variable que representa la cantidad de dias totales en 10 año.
var DAYS = 10 * 365

// Función que permite calcular la cantidad de dias óptimos.
// Parámetros: Tres pares de enteros a,b y c que son la velocidad angular y el radio de cada punto.
func OptimalDays(a_ang, a_dis, b_ang, b_dis, c_ang, c_dis int) int {
	var optimal_days = 0
	var a_position, b_position, c_position int = 0, 0, 0
	for i := 0; i < DAYS; i++ {
		p1 := Rad2Cart(float64(a_dis), float64(a_position))
		p2 := Rad2Cart(float64(b_dis), float64(b_position))
		p3 := Rad2Cart(float64(c_dis), float64(c_position))
		if CheckLine(p1, p2, p3) {
			optimal_days++
		}

		a_position = (a_position + a_ang + 360) % 360
		b_position = (b_position + b_ang + 360) % 360
		c_position = (c_position + c_ang + 360) % 360
	}
	return optimal_days
}

// Función que permite calcular la cantidad de dias lluviosos y el dia mas lluvioso contando la cantidad
// de veces en las que el sol se encuentra en el triángulo formado por los plantas.
// Parámetros: Tres pares de enteros a,b y c que son la velocidad angular y el radio de cada punto.
func RainyDays(a_ang, a_dis, b_ang, b_dis, c_ang, c_dis int) (int, int) {
	var rainy_days, rainiest_day = 0, 0
	var max_perimeter float64 = 0
	var a_position, b_position, c_position int = 0, 0, 0
	for i := 0; i < DAYS; i++ {
		p1 := Rad2Cart(float64(a_dis), float64(a_position))
		p2 := Rad2Cart(float64(b_dis), float64(b_position))
		p3 := Rad2Cart(float64(c_dis), float64(c_position))
		if SunContained(p1, p2, p3) {
			rainy_days++
			perimeter := TrianglePerimeter(p1, p2, p3)
			if perimeter > max_perimeter {
				max_perimeter = perimeter
				rainiest_day = i
			}
		}

		a_position = (a_position + a_ang + 360) % 360
		b_position = (b_position + b_ang + 360) % 360
		c_position = (c_position + c_ang + 360) % 360
	}
	return rainy_days, rainiest_day
}

// Función generalizada para calcular los dias de sequía de forma matemática dados las velocidades angulares de los planetas.
// Parámetros: Tres enteros a,b y c que representan las velocidades angulares de los planetas.
func DroughtDays(a_ang int, b_ang int, c_ang int) int {
	return CeilDiv(DAYS, Mcm(Congruence(a_ang, b_ang), Congruence(a_ang, c_ang)))
}

// Función generalizada para calcular los dias de sequía de forma iterativa dados las velocidades angulares de los planetas.
// Parámetros: Tres enteros a,b y c que representan las velocidades angulares de los planetas.
func DroughtDaysIterative(a_ang int, b_ang int, c_ang int) int {
	var a_position, b_position, c_position int = 0, 0, 0
	var drought_days = 0
	for i := 0; i < DAYS; i++ {
		if a_position == b_position && a_position == c_position {
			drought_days++
		}
		a_position = (a_position + a_ang + 180) % 180
		b_position = (b_position + b_ang + 180) % 180
		c_position = (c_position + c_ang + 180) % 180
	}
	return drought_days
}

// Función encargada de chequear la correctitud de los algoritmos implementados para calcular los dias de sequía.
// Parámetros: Tres enteros a,b y c que representan las velocidades angulares de los planetas.
func CheckDroughtDaysCorrectnes() bool {
	var correct = true
	for i := 1; i < 50; i++ {
		for j := 1; j < 50; j++ {
			for k := 1; k < 50; k++ {
				if DroughtDays(i, j, k) != DroughtDaysIterative(i, j, k) {
					correct = false
					break
				}
			}
			if correct == false {
				break
			}
		}
		if correct == false {
			break
		}
	}
	return correct
}
