package main

import (
	"fmt"
	"go-fundamentos-1598/internal/productos"
	"log"
	"os"
	"strings"
)

const (
	filename = "./productos.csv"
)

func main() {
	// Recupermos el panic y lo handleamos
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	// Creamos un storage en memoria
	storage := productos.Storage{
		Productos: readfile(filename),
	}

	canalProducto := make(chan productos.Producto)
	defer close(canalProducto)
	canalError := make(chan error)
	defer close(canalError)

	var entrada string

	fmt.Print("Ingrese nombre de producto a buscar: ")
	_, err := fmt.Scanln(&entrada)
	if err != nil {
		log.Fatal(err)
	}

	// Generamos una go routine
	go func(chan productos.Producto, chan error) {
		producto, err := storage.BuscarProductoPorNombre(entrada)
		if err != nil {
			canalError <- err
			return
		}

		canalProducto <- producto

	}(canalProducto, canalError)

	// Evaluamos los canales
	select {
	case pr := <-canalProducto:
		fmt.Println(pr)
	case err := <-canalError:
		//log.Fatal(err)
		log.Println(err)
		os.Exit(1)
	}

}

// radfile..
func readfile(filname string) []productos.Producto {

	listaProductos := make([]productos.Producto, 0)

	file, err := os.ReadFile(filname)
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n")

	for i := 0; i < len(data); i++ {

		if len(data[i]) > 0 {
			line := strings.Split(string(data[i]), ",")
			producto := productos.Producto{
				Nombre:      line[0],
				Descripcion: line[1],
				Cantidad:    line[2],
				Precio:      line[3],
			}

			listaProductos = append(listaProductos, producto)

		}

	}

	return listaProductos

}
