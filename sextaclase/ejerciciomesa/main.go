package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	categoriacsv    = "./categorias.csv"
	estimacionescsv = "./estimaciones.csv"
)

// Producto describe un producto
type Producto struct {
	Codigo         string
	Nombre         string
	PrecioActual   float64
	CantidadActual int
}

// Categoria describe una categoria
type Categoria struct {
	Codigo         string
	Nombre         string
	ListaProductos []Producto
}

// Tienda describe una tienda
type Tienda struct {
	Categorias []Categoria
}

// Estimacion describe una estimacion de ventas por categoria.
type Estimacion struct {
	Categoria              string
	EstimativoPorCategoria float64
}

func main() {
	// escribimos el archivo csv
	writeCategoriaCsv(categoriacsv, loadTienda())
	// leemos el archivo csv y recuperamos la tienda
	tienda := readCategoriaCsv(categoriacsv)
	// calculamos el estimativo por categoria
	estimaciones := calcularEstimativos(tienda)
	// escribimos el archivo csv
	writeEstimativosCsv(estimacionescsv, estimaciones)
	// Imprimos por consola
	printConsole(estimaciones)

}

// loadTienda carga la tienda
func loadTienda() Tienda {
	// lista de productos de electrodomesticos
	listaProductosElectrodomesticos := []Producto{
		{
			Codigo:         "001",
			Nombre:         "Heladera",
			PrecioActual:   1000,
			CantidadActual: 10,
		},
		{
			Codigo:         "002",
			Nombre:         "Lavarropas",
			PrecioActual:   2000,
			CantidadActual: 20,
		},
		{
			Codigo:         "003",
			Nombre:         "Televisor",
			PrecioActual:   3000,
			CantidadActual: 30,
		},
	}
	// categoria de electrodomesticos
	categoriaElectrodomesticos := Categoria{
		Codigo:         "001",
		Nombre:         "Electrodomesticos",
		ListaProductos: listaProductosElectrodomesticos}

	// lista de productos de muebles
	listaProductosMuebles := []Producto{
		{
			Codigo:         "001",
			Nombre:         "Mesa",
			PrecioActual:   1000,
			CantidadActual: 10,
		},
		{
			Codigo:         "002",
			Nombre:         "Silla",
			PrecioActual:   2000,
			CantidadActual: 20,
		},
		{
			Codigo:         "003",
			Nombre:         "Rack",
			PrecioActual:   3000,
			CantidadActual: 30,
		},
	}

	// categoria de muebles
	categoriaMuebles := Categoria{
		Codigo:         "002",
		Nombre:         "Muebles",
		ListaProductos: listaProductosMuebles,
	}

	// lista de productos de herramientas
	listaProductosHerramientas := []Producto{
		{
			Codigo:         "001",
			Nombre:         "Martillo",
			PrecioActual:   1000,
			CantidadActual: 10,
		},
		{
			Codigo:         "002",
			Nombre:         "Destornillador",
			PrecioActual:   2000,
			CantidadActual: 20,
		},
		{
			Codigo:         "003",
			Nombre:         "Taladro",
			PrecioActual:   3000,
			CantidadActual: 30,
		},
	}

	// categoria de herramientas
	categoriaHerramientas := Categoria{
		Codigo:         "003",
		Nombre:         "Herramientas",
		ListaProductos: listaProductosHerramientas,
	}

	// carga la tienda
	tienda := Tienda{}
	tienda.Categorias = append(tienda.Categorias, categoriaElectrodomesticos)
	tienda.Categorias = append(tienda.Categorias, categoriaMuebles)
	tienda.Categorias = append(tienda.Categorias, categoriaHerramientas)
	return tienda
}

// writeCategoriaCsv escribe el archivo csv
func writeCategoriaCsv(filename string, tienda Tienda) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, categoria := range tienda.Categorias {
		f.WriteString(fmt.Sprintf("%s,%s,%v\n", categoria.Codigo, categoria.Nombre, categoria.ListaProductos))
	}
}

// readCategoriaCsv lee el archivo csv
func readCategoriaCsv(filename string) Tienda {
	// Creo una tienda
	tienda := Tienda{}
	// Lee el contenido del archivo CSV
	contenido, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Convierte el contenido en una cadena
	contenidoStr := string(contenido)

	// Divide el contenido en líneas
	lineas := strings.Split(contenidoStr, "\n")

	// Crea un slice para almacenar las categorías
	categorias := make([]Categoria, 0)

	var categoria Categoria

	// Procesa las líneas
	for i := 0; i < len(lineas)-1; i++ {
		campos := strings.Split(lineas[i], ",")
		if len(campos) == 3 {
			if categoria.Codigo != "" {
				categorias = append(categorias, categoria)
			}
			categoria = Categoria{
				Codigo: campos[0],
				Nombre: campos[1],
			}

			productoStr := campos[2]
			productoStr = strings.TrimPrefix(productoStr, "[")
			productoStr = strings.TrimSuffix(productoStr, "]")
			productoCampos := strings.Split(productoStr, " ")
			if len(productoCampos) == 12 {
				precioActual, err := strconv.ParseFloat(productoCampos[2], 64)
				if err != nil {
					log.Fatal("Error al convertir el precio actual a float", err)
				}
				cantidadActual, err := strconv.ParseInt(productoCampos[10], 10, 64)
				if err != nil {
					log.Fatal("Error al convertir la cantidad actual a int", err)
				}
				producto := Producto{
					Codigo:         productoCampos[0],
					Nombre:         productoCampos[1],
					PrecioActual:   precioActual,
					CantidadActual: int(cantidadActual),
				}
				categoria.ListaProductos = append(categoria.ListaProductos, producto)
			}
		}
	}

	// Agregar la última categoría
	if categoria.Codigo != "" {
		categorias = append(categorias, categoria)
	}

	tienda.Categorias = append(tienda.Categorias, categorias...)

	return tienda

}

// calcularEstimativos calcula los estimativos de ventas por categoría.
func calcularEstimativos(tienda Tienda) []Estimacion {
	estimaciones := make([]Estimacion, 0)
	for _, categoria := range tienda.Categorias {
		estimacion := Estimacion{
			Categoria:              categoria.Nombre,
			EstimativoPorCategoria: 0,
		}
		for _, producto := range categoria.ListaProductos {
			estimacion.EstimativoPorCategoria += producto.PrecioActual * float64(producto.CantidadActual)
		}
		estimaciones = append(estimaciones, estimacion)
	}
	return estimaciones
}

// writeEstimativosCsv escribe el archivo csv
func writeEstimativosCsv(filename string, estimaciones []Estimacion) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, estimacion := range estimaciones {
		f.WriteString(fmt.Sprintf("%s,%.2f,\n", estimacion.Categoria, estimacion.EstimativoPorCategoria))
	}
}

// printConsole imprime en consola las estimaciones
func printConsole(estimaciones []Estimacion) {
	fmt.Println("Categoría\t\t\tEstimativo")
	var total float64
	for _, estimacion := range estimaciones {
		if estimacion.Categoria == "Muebles" {
			fmt.Printf("%s\t\t\t\t%.2f\n", estimacion.Categoria, estimacion.EstimativoPorCategoria)
		} else if estimacion.Categoria == "Herramientas" {
			fmt.Printf("%s\t\t\t%.2f\n", estimacion.Categoria, estimacion.EstimativoPorCategoria)
		} else {
			fmt.Printf("%s\t\t%.2f\n", estimacion.Categoria, estimacion.EstimativoPorCategoria)
		}

		total += estimacion.EstimativoPorCategoria
	}

	fmt.Printf("\t\tTotal\t\t%.2f\n", total)

}
