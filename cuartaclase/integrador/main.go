package integrador

import (
	"fmt"
	"time"
)

// Autor es una estructura que describe un objeto Autor.
type Autor struct {
	Nombre   string
	Apellido string
}

// Libro es una estructura que describe un objeto Libro.
type Libro struct {
	Autor            Autor
	Descripcion      string
	FechaPublicacion time.Time
	Titulo           string
}

// SetNombre es un metodo que permite modificar el nombre del autor.
// El metodo recibe como parametro un puntero a la estructura Autor.
// Esto nos permite modificar el valor de la estructura Autor.
// Si no recibieramos un puntero, el metodo no podria modificar el valor de la estructura Autor.
// Estariamos trabajando con una copia de la estructura Autor.
func (a *Autor) SetNombre(nombre string) {
	a.Nombre = nombre
}

// SetTitle es un metodo que permite modificar el titulo del libro.
func (l *Libro) SetTitle(titulo string) {
	l.Titulo = titulo
}

func main() {

	autor := Autor{
		Nombre:   "Juan",
		Apellido: "Perez",
	}

	libro := Libro{
		Autor:            autor,
		Descripcion:      "Libro de programacion",
		FechaPublicacion: time.Now(),
		Titulo:           "El mejor libro de programacion",
	}

	fmt.Println("El titulo del libro es:", libro.Titulo)
	libro.SetTitle("Nuevo titulo")
	fmt.Println("El titulo del libro es:", libro.Titulo)
	fmt.Println("El nombre del autor es:", libro.Autor.Nombre)
	libro.Autor.SetNombre("Nuevo nombre")
	fmt.Println("El nombre del autor es:", libro.Autor.Nombre)

}
