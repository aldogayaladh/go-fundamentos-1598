package typeassertion

import "fmt"

func main() {
	// Creamos una variable de tipo interface
	// Esta variable puede almacenar cualquier tipo de dato
	var variable interface{}

	// Le asignamos un valor de tipo entero
	variable = 1

	// Realizamos una asercion de tipo
	// Comprobamos si la variable es de tipo string
	// Si es de tipo string, la variable cadena tendra el valor de la variable
	// Si no es de tipo string, la variable cadena tendra el valor zero de un string
	// En este caso, cadena tendra el valor ""
	// Y la variable ok tendra el valor false
	cadena, ok := variable.(string)
	if ok {
		fmt.Println("Variable", cadena)
	} else {
		// Se cumple esta condicion
		// Imprimimos el valor zero de un string
		fmt.Println("No es un string", cadena)
	}
}
