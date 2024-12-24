package utils

import (
	"math"
	"sort"
)

// Estructura encargada de representar un punto en el plano cartesiano.
type Point struct {
	X float64
	Y float64
}

// Función encargada de determinar si 3 puntos se encuentran sobre la misma recta sin el sol.
// Parámetros: Tres puntos p1, p2 y p3.
func CheckLine(p1, p2, p3 Point) bool {

	if CompareEqual(p1.X, p2.X) && CompareEqual(p1.X, p3.X) && !CompareEqual(p1.X, 0.0) { //El divisor de la pendiente es 0, división por 0
		return true
	}
	if CompareEqual(p1.X, p2.X) {
		p1, p3 = p3, p1
	}

	m := (p1.Y - p2.Y) / (p1.X - p2.X) //Pendiente de la recta formada por p1 y p2

	if CompareEqual(p1.Y, m*p1.X) { //Se intersecta con el sol
		return false
	}
	if CompareEqual(p3.Y-p1.Y, m*(p3.X-p1.X)) { //El tercer punto está sobre la misma recta
		return true
	}
	return false
}

// Función encargada de determinar si el sol (0,0) está contenido en un triángulo dado.
// El algoritmo usado para esta implementación es el algoritmo general para determinar si un
// punto está contenido dentro de un polígono. Consiste en trazar un segmento desde el punto deseado hasta un
// punto muy lejano cualquiera. Si la cantidad de intersecciones con segmentos es impar, el punto está contenido en el polígono.
// Para mas información del algoritmo, visitar Competitive Programmer’s Handbook, Antti Laaksonen, Chapter 29.
// Parámetros: Tres puntos: p1,p2 y p3 que determinan el triángulo.
func SunContained(p1, p2, p3 Point) bool {
	sun := Point{X: 0, Y: 0}
	farAway := Point{X: 100000, Y: 1}
	cont := 0
	if Intersection(sun, farAway, p1, p2) {
		cont++
	}
	if Intersection(sun, farAway, p1, p3) {
		cont++
	}
	if Intersection(sun, farAway, p2, p3) {
		cont++
	}
	if cont%2 == 1 {
		return true
	}
	return false
}

// Función encargada de determinar si dos segmentos se tocan o intersectan.
// Hay tres casos que determinan la intersección de los segmentos:
// 1. Estos son paralelos y están solapados.
// 2. Comparten algún punto en común.
// 3. Se atraviesan.
// Parámetros: Cuatro puntos p1,p2,p3 y p4. En caso de que el segmento formado por (p1,p2)
// intersecte al segmento formado por (p3 y p4) se retorna true.
func Intersection(p1, p2, p3, p4 Point) bool {
	cp1 := CrossProduct(p1, p2, p3)
	cp2 := CrossProduct(p1, p2, p4)
	cp3 := CrossProduct(p3, p4, p1)
	cp4 := CrossProduct(p3, p4, p2)
	//Los siguientes 4 ifs corroboran si un punto está contenido en el segmento opuesto o es igual a algún punto de otro segmento
	if CompareEqual(cp1, 0) && midCheck(p1, p2, p3) {
		return true
	}
	if CompareEqual(cp2, 0) && midCheck(p1, p2, p4) {
		return true
	}
	if CompareEqual(cp3, 0) && midCheck(p3, p4, p1) {
		return true
	}
	if CompareEqual(cp4, 0) && midCheck(p3, p4, p2) {
		return true
	}
	//El siguiente if corrobora si los segmentos se atraviesan
	if sign(cp1) != sign(cp2) && sign(cp3) != sign(cp4) {
		return true
	}
	return false
}

// Función encargada de determinar si un punto se encuentra contenido entre otros dos puntos.
// Parámetros: Tres puntos p1,p2,p3. Si p3 está entre p1 y p2, se retorna true.
func midCheck(p1, p2, p3 Point) bool {
	// Creamos un slice con los tres puntos
	v := []Point{p1, p2, p3}

	// Ordenamos el slice usando sort.Slice y la función `comp`
	sort.Slice(v, func(i, j int) bool {
		return comparePoint(v[i], v[j])
	})

	// Verificamos si el punto intermedio es igual a `p3`
	return EqualPoint(v[1], p3)
}

// Función encargada de determinar dos puntos son iguales.
// Parámetros: Dos puntos p1,p2. Si p1 es igual a p2, se retorna true.
func EqualPoint(p1, p2 Point) bool {
	return CompareEqual(p1.X, p2.X) && CompareEqual(p1.Y, p2.Y)
}

// Función encargada de determinar si un punto es mayor a otro.
// Parámetros: Dos puntos p1,p2. Si p1 es menor a p2, se retorna true.
func comparePoint(p1, p2 Point) bool {
	if p1.X == p2.X {
		return p1.Y < p2.Y
	}
	return p1.X < p2.X
}

// Función encargada de calcular el producto cruz entre dos vectores formados por tres puntos.
// Los vectores formados son los siguientes (p2-p1) y (p3-p1).
// Parámetros: Tres puntos p1,p2,p3 con los cuales se crean dos vectores y se calcula el producto cruz entre ellos.
func CrossProduct(p1, p2, p3 Point) float64 {
	u := Point{X: p2.X - p1.X, Y: p2.Y - p1.Y}
	v := Point{X: p3.X - p1.X, Y: p3.Y - p1.Y}
	return u.X*v.Y - v.X*u.Y
}

// Función encargada de calcular el perímetro de un triángulo dados 3 puntos.
// Parámetros: Tres puntos sobre los cuales se calcula el perímetro del triángulo formado por ellos.
func TrianglePerimeter(p1, p2, p3 Point) float64 {
	return euclidDistance(p1, p2) + euclidDistance(p2, p3) + euclidDistance(p1, p3)
}

// Función encargada de calcular la distancia euclidiana entre dos puntos.
// Parámetros: Dos puntos sobre los cuales se calcula la distancia euclidiana.
func euclidDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}

// Función encargada de convertir coordenadas radiales a cartesianas.
// Parámetros: Dos decimales r y theta.
func Rad2Cart(r, angle float64) Point {
	theta := angle * (math.Pi / 180)

	x := r * math.Cos(theta)
	y := r * math.Sin(theta)
	return Point{X: x, Y: y}
}

// Función para calcular el resultado de una congruencia de la forma (ax)mod180 = (bx)mod180.
// Parámetros: Dos enteros a y b sobre los cuales se calcula el resultado de la ecuación.
func Congruence(a int, b int) int {
	return 180 / gcd(180, abs(a-b))
}

// Función para calcular el mínimo comun múltiplo(MCM).
// Parámetros: Dos enteros a y b sobre los cuales se calcula el MCM.
func Mcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

// Función para calcular el máximo común divisor (MCD).
// Parámetros: Dos enteros a y b sobre los cuales se calcula el MCD.
func gcd(a int, b int) int {
	m := max(a, b)
	n := min(a, b)
	if n != 0 {
		return gcd(n, m%n)
	}
	return m
}

// Función para calcular el máximo de dos números.
// Parámetros: Dos enteros a y b sobre los cuales se calcula el máximo.
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Función para calcular el mínimo de dos números.
// Parámetros: Dos enteros a y b sobre los cuales se calcula el mínimo.
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Función para calcular el valor absoluto de un número.
// Parámetros: Un enteros a sobre el cual se calcula el valor absoluto.
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

// Función para calcular el techo de una división dados dos números.
// Parámetros: Dos enteros a y b sobre los cuales se calcula el techo de la división.
func CeilDiv(a int, b int) int {
	var res = a / b
	if a%b > 1 {
		res++
	}
	return res
}

// Función encargada de corregir el problema de la precisión de los floats.
// Se comparan dos números y si la diferencia entre ellos es menor a un epsilon, se consideran iguales.
// Parámetros: Dos floats a y b.
func CompareEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.1
}

// Función para calcular el signo de un número.
// Parámetros: Un entero a, se retorna -1 si es negativo, 0 si es cero y 1 si es positivo.
func sign(a float64) int {
	if CompareEqual(a, 0) {
		return 0
	} else if a > 0.0 {
		return 1
	} else {
		return -1
	}
}
