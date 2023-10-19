package metodos

import "fmt"

// Persona es una estructura que describe un objeto Persona.
type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

// Empleado es una estructura que describe un objeto Empleado.
type Empleado struct {
	Persona Persona
	Cargo   string
	Legajo  string
}

// Empleador es una estructura que describe un objeto Empleador.
type Empleador struct {
	Persona     Persona
	RazonSocial string
	Cuit        string
}

// Podemos tener mas de un metodo con la misma firma en el mismo package, pero que hagan referencia a distintas estructuras.
func (e Empleado) Saludar() {
	fmt.Println("Hola soy un empleado")
}

func (e Empleador) Saludar() {
	fmt.Println("Hola soy un empleador")
}

func main() {
	// Creamos un objeto de tipo persona empleado.
	personaEmpleado := Persona{
		Nombre:   "John",
		Apellido: "Doe",
		Edad:     30,
	}

	// Creamos un objeto de tipo empleado.
	empleado := Empleado{
		Persona: personaEmpleado,
		Cargo:   "Programador",
		Legajo:  "123456",
	}

	// Creamos un objeto de tipo persona empleador.
	personaEmpleador := Persona{
		Nombre:   "Jane",
		Apellido: "Doe",
		Edad:     30,
	}

	// Creamos un objeto de tipo empleador.
	empleador := Empleador{
		Persona:     personaEmpleador,
		RazonSocial: "Empresa S.A.",
		Cuit:        "123456789",
	}

	// Llamamos a sus metodos
	empleado.Saludar()
	empleador.Saludar()
}
