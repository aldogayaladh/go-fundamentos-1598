package interfaces

import "fmt"

// Animal es una interfaz que define un metodo saludar. (contrato)
// Puede tener uno o mas metodos.
type Animal interface {
	Saludar()
}

// Gato es una estructura que implementa la interfaz Animal.
type Gato struct {
	saludo string
}

// Saludar es un metodo que implementa la interfaz Animal.
func (g Gato) Saludar() {
	fmt.Println(g.saludo)
}

// Caminar es un metodo que implementa el animal Gato.
func (g Gato) Caminar() {
	fmt.Println("El gato camina")
}

// Perro es una estructura que implementa la interfaz Animal.
type Perro struct {
	saludo string
}

// Saludar es un metodo que implementa la interfaz Animal.
func (p Perro) Saludar() {
	fmt.Println(p.saludo)
}

// Rana es una estructura que implementa la interfaz Animal.
type Rana struct {
	saludo string
}

// Saludar es un metodo que implementa la interfaz Animal.
func (r Rana) Saludar() {
	fmt.Println(r.saludo)
}

// NewAnimal es una funcion que devuelve un objeto de tipo Animal.
func NewAnimal(nombre string) Animal {
	switch nombre {
	case "gato":
		return Gato{saludo: "Miau"}
	case "perro":
		return Perro{saludo: "Guau"}
	case "rana":
		return Rana{saludo: "croc"}
	default:
		return nil
	}
}

func main() {
	// Creamos un objeto de tipo Gato.
	animalGato := NewAnimal("gato")
	animalGato.Saludar()
	// Creamos un objeto de tipo Perro.
	animalPerro := NewAnimal("perro")
	animalPerro.Saludar()
	// Creamos un objeto de tipo Rana.
	animalRana := NewAnimal("rana")
	animalRana.Saludar()
}
