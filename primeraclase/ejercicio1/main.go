package main

import "fmt"

func main() {

	// 	La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
	//  A - Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
	//  B - Luego, imprimir cada una de las letras.

	var palabra string = "diccionario"

	// Punto A.

	// len - Cantidad de letras que tiene la palabra.
	cantidad := len(palabra) // retorna un entero.

	// Salida por consola formateada de la palabra + cantidad de letras.
	fmt.Printf("La cantidad de letras que tiene la palabra %s es: %d\n", palabra, cantidad)

	// Punto B.

	// Solucion 1.
	// Recorrer la palabra con un for convencional.
	for i := 0; i < len(palabra); i++ {
		fmt.Println(string(palabra[i]))

	}

	// Solucion 2.
	// Recorrer la palabra por clave - valor usando range.
	for clave, valor := range palabra {
		fmt.Printf("Clave: %d - Valor: %s\n", clave, string(valor))
	}

}
