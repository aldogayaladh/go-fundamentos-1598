package main

import (
	"encoding/json"
	"fmt"
)

// // // Employee
// // // Una empresa necesita realizar una buena gestión de sus empleados,
// //  para esto realizaremos un pequeño programa nos ayudará a gestionar correctamente
//     dichos empleados. Los objetivos son:
// // // Crear una estructura Person con los campos ID, Name, DateOfBirth.
// // // Crear una estructura Employee con los campos: ID, Position y una
//  composición con la estructura Person.
// // // Realizar un método a la estructura Employee que se llame PrintEmployee(),
// lo que hará es realizar la impresión de los campos de un empleado.
// // // Instanciar en la función main() tanto una Person como un Employee cargando
// sus respectivos campos y, por último, ejecutar el método PrintEmployee().
// // // Si logramos realizar este pequeño programa, pudimos ayudar a la empresa
// a solucionar la gestión de los empleados.

// Persona es una estructura que descibre el modelo de una persona.
type Persona struct {
	ID              string `json:"id"`
	Nombre          string `json:"nombre"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Password        string `json:"-"`
}

// Empleado es una estructura que descibre a un empleado.
type Empleado struct {
	ID       string  `json:"id"`
	Posicion string  `json:"posicion"`
	Persona  Persona `json:"persona"`
}

// PrintInfo es un metodo que imprime la informacion de un empleado.
func (e Empleado) PrintInfo() {
	fmt.Println("Id empleado:", e.ID)
	fmt.Println("Posicion empleado:", e.Posicion)
	fmt.Println("Id persona:", e.Persona.ID)
	fmt.Println("Nombre:", e.Persona.Nombre)
	fmt.Println("Fecha Nacimiento:", e.Persona.FechaNacimiento)
}

func main() {

	persona := Persona{
		ID:              "AAABCCC",
		Nombre:          "Juan Perez",
		FechaNacimiento: "1994-03-12",
		Password:        "123456",
	}

	empleado := Empleado{
		ID:       "ZZZTTRR",
		Posicion: "Desarrollador",
		Persona:  persona,
	}

	// Llamado del metodo print info de empleado.
	empleado.PrintInfo()

	// Extra. Conversion a JSON + Tags.
	{
		jsonEmpleado, err := json.Marshal(empleado)

		if err != nil {
			fmt.Println("Error is:", err)
		}

		fmt.Println(string(jsonEmpleado))
	}

}
