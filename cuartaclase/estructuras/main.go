package estructuras

import (
	"fmt"
	"time"
)

// RequestPersona es una estructura que describe un objeto de solicitud de una Persona.
type RequestPersona struct {
	Nombre          string    `json:"nombre"`
	Edad            int       `json:"edad"`
	Direccion       string    `json:"direccion"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"`
}

// ResponsePersona es una estructura que describe un objeto de respuesta de una Persona.
type ResponsePersona struct {
	Nombre          string    `json:"nombre"`
	Edad            int       `json:"edad"`
	Direccion       string    `json:"direccion"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"`
	Usuario         Usuario   `json:"usuario"`
}

// Persona es una estructura que describe un objeto Persona.
type Persona struct {
	Nombre          string    `json:"nombre"`
	Edad            int       `json:"edad"`
	Direccion       string    `json:"direccion"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"`
	Usuario         Usuario   `json:"usuario"`
}

type Usuario struct {
	Username string `json:"username"`
	Password string `json:"-"` // No se muestra en el JSON de respuesta
}

func main() {
	// Creamos un objeto de tipo Usuario
	usuario := Usuario{
		Username: "johndoe",
		Password: "123456",
	}

	// Composicion de objetos
	// Creamos un objeto de tipo Persona
	// y le asignamos el objeto Usuario
	persona := Persona{
		Nombre:          "John Doe",
		Edad:            30,
		Direccion:       "Calle falsa 123",
		FechaNacimiento: time.Now(),
		Usuario:         usuario,
	}

	// Imprimimos el nombre de la persona
	fmt.Println("El nombre de la persona es:", persona.Nombre)
	// Imprimimos el nombre de usuario
	fmt.Println("El nombre de usuario es:", persona.Usuario.Username)
}
