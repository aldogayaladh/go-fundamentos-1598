package main

import (
	"errors"
	"fmt"
)

const (
	salaryConst = 10000
)

// ErrorPersonalizado es un error personalizado
type ErrorPersonalizado struct {
	Codigo  int
	Mensaje string
}

// Error es un metodo que implementa la interfaz error
func (e *ErrorPersonalizado) Error() string {
	return fmt.Sprintf("Codigo: %d - Mensaje: %s", e.Codigo, e.Mensaje)
}

func main() {

	// Validamos el salario
	result, err := validateSalary(5000)
	// Buscamos el error personalizado
	var errorABuscar *ErrorPersonalizado
	// Si el error es del tipo ErrorPersonalizado
	if errors.As(err, &errorABuscar) {
		fmt.Println(err.Error())
	} else {
		fmt.Println("El salario es:", result)
	}

}

// validateSalary valida el salario
func validateSalary(salary int) (int, error) {
	if salary <= salaryConst {
		// Retornamos un error personalizado
		return 0, &ErrorPersonalizado{Codigo: 1, Mensaje: "Error personalizado"}
	}

	return salary, nil
}
