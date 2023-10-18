package main

import "fmt"

func main() {

	// Realizar una aplicación que contenga una variable con el número del mes.
	// Según el número, imprimir el mes que corresponda en texto.
	// ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?

	// Entero con el mes a buscar.
	mesBusqueda := 3

	// Solucion 1 : Uso de if - else.
	if mesBusqueda == 1 {
		fmt.Println("Enero")
	} else if mesBusqueda == 2 {
		fmt.Println("Febrero")
	} else if mesBusqueda == 3 {
		fmt.Println("Marzo")
	} else if mesBusqueda == 4 {
		fmt.Println("Abril")
	} else if mesBusqueda == 5 {
		fmt.Println("Mayo")
	} else if mesBusqueda == 6 {
		fmt.Println("Junio")
	} else if mesBusqueda == 7 {
		fmt.Println("Julio")
	} else if mesBusqueda == 8 {
		fmt.Println("Agosto")
	} else if mesBusqueda == 9 {
		fmt.Println("Septiembre")
	} else if mesBusqueda == 10 {
		fmt.Println("Octubre")
	} else if mesBusqueda == 11 {
		fmt.Println("Noviembre")
	} else if mesBusqueda == 12 {
		fmt.Println("Diciembre")
	} else {
		fmt.Println("Mes no existe")
	}

	// Solucion 2. Uso de switch - case.

	switch mesBusqueda {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("Mes no existe")

	}

	// Solucion 4. Uso de un array + for.
	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	var mesBuscado string
	for i := 0; i < len(meses); i++ {
		if i == mesBusqueda {
			mesBuscado = string(meses[i-1])
			break
		}
	}

	if mesBuscado != "" {
		fmt.Println(mesBuscado)
	} else {
		fmt.Println("Mes no existe")
	}

	// Solucion 4. Uso de un mapa.
	mapaMeses := map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	mes, ok := mapaMeses[mesBusqueda]
	if !ok {
		fmt.Println("Mes no existe")
	} else {
		fmt.Println(mes)
	}

}
