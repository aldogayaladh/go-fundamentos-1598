package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivision = errors.New("No se puede divir por cero")
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("El programa ha finalizado")
			fmt.Println("Recover", err)
		}
	}()

	resultado, err := division(10, 0)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("El resultado de la divison es:", resultado)
	}

}

func division(a, b int) (int, error) {
	if b < 1 {
		return 0, fmt.Errorf("El error en la division es: %w", ErrDivision)
	}

	return a / b, nil
}
