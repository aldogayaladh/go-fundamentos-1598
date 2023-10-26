package main

import (
	"errors"
	"fmt"
)

const (
	salaryConst = 10000
)

var (
	// ErrSalary es un error que utiliza el paquete errors
	ErrSalary = errors.New("Salario menor o igual a 10000")
)

func main() {

	// Validamos el salario
	result, err := validateSalary(5000)
	// Verificamos si el error es del tipo ErrSalary
	if errors.Is(err, ErrSalary) {
		fmt.Println(err.Error())
	} else {
		fmt.Println("El salario es:", result)
	}

}

// validateSalary valida el salario
func validateSalary(salary int) (int, error) {

	if salary <= salaryConst {
		return 0, ErrSalary
	}

	return salary, nil
}
