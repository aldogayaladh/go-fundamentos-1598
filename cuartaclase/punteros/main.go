package punteros

import "fmt"

func main() {
	var contador int
	contador = 1
	fmt.Println("Contador", contador)
	// Pasamos la direccion de memoria de la variable contador. &contador
	incrementar(&contador)
	// Imprimimos el valor de la variable contador
	fmt.Println("Contador despues de incrementar", contador)
}

// incrementar incrementa en 1 el valor de un contador
// Recibe un puntero a un entero
func incrementar(contador *int) {
	// Desreferenciamos el puntero para obtener el valor de la variable
	// Incrementamos el valor de la variable
	*contador++
}
