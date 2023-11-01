package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	var input string
	channelDevolucion := make(chan string)
	channelPrestamo := make(chan string)

	// Escribimos en el channel devolucion
	go Devolucion(channelDevolucion)

	// Imprimimos el contenido del canal
	go PrintDevolucion(channelDevolucion)

	// Escribimos en el channel prestamo
	go Prestamo(channelPrestamo)

	// Imprimimos el contenido del canal
	go PrintPrestamo(channelPrestamo)

	fmt.Scan(&input)
	if input != "" {
		fmt.Println("Saliendo del programa...")
		os.Exit(0)
	}

}

// Devolucion ....
func Devolucion(channelDevolucion chan string) {
	defer close(channelDevolucion)
	for {
		time.Sleep(time.Second * 1)
		channelDevolucion <- "Devolucion  procesada"
	}
}

// Prestamo ....
func Prestamo(channelPrestamo chan string) {
	defer close(channelPrestamo)
	for {
		time.Sleep(time.Second / 2)
		channelPrestamo <- "Prestamo procesado"

	}
}

// PrintDevolucion ....
func PrintDevolucion(channelDevolucion chan string) {
	for devolucion := range channelDevolucion {
		fmt.Println(devolucion)
	}
}

// PrintPrestamo ....
func PrintPrestamo(channelPrestamo chan string) {
	for prestamo := range channelPrestamo {
		fmt.Println(prestamo)
	}
}

// output
// Devolucion procesada - 1 segundo
// Prestamo procesado
// Prestamo procesado
