package main

import (
	"errors"
	"fmt"
)

// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de
// calificaciones de los estudiantes de un curso. Requieren calcular los valores mínimo,
// máximo y promedio de sus calificaciones.
// Para esto, se solicita generar una función que indique qué tipo de cálculo se quiere realizar
// (mínimo, máximo o promedio) y que devuelva otra función y un mensaje
// (en caso de que el cálculo no esté definido), que se le pueda pasar una cantidad N de enteros y
// devuelva el cálculo que se indicó en la función anterior. Veamos un ejemplo:

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {

	minFunc, err := operation(minimo)
	fmt.Println("Error es:", err)
	promedioFunc, err := operation(promedio)
	fmt.Println("Error es:", err)
	maxFunc, err := operation(maximo)

	fmt.Println("Error es:", err)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	promedioValue := promedioFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Minimo", minValue)
	fmt.Println("Maximo", maxValue)
	fmt.Println("Promedio", promedioValue)

}

func calcularMinimo(valores ...int) int {
	var min int = 10000
	for _, valor := range valores {
		if valor < min {
			min = valor
		}
	}
	return min
}

func calcularMaximo(valores ...int) int {
	var max int
	for _, valor := range valores {
		if valor > max {
			max = valor
		}
	}
	return max
}

// func(...int) int
func calcularPromedio(valores ...int) int {
	var suma int

	for _, valor := range valores {
		suma += valor
	}

	promedio := suma / len(valores)
	return promedio
}

func operation(operacion string) (func(...int) int, error) {
	switch operacion {
	case minimo:
		funcMin := calcularMinimo
		return funcMin, nil
	case maximo:
		funcMax := calcularMaximo
		return funcMax, nil
	case promedio:
		funcPromedio := calcularPromedio
		return funcPromedio, nil
	default:
		return nil, errors.New("funcion no existe")
	}
}
