package productos

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("no se encontro el producto")
)

// Producto es una estructura que define un producto de un supermercado.
type Producto struct {
	Nombre      string
	Descripcion string
	Cantidad    string
	Precio      string
}

// Storage ....
type Storage struct {
	Productos []Producto
}

func (s *Storage) PrintInfo() {
	fmt.Printf("%v+", s.Productos)
}

// BuscarProductoPorNombre es una func...
func (s *Storage) BuscarProductoPorNombre(nombre string) (Producto, error) {

	for _, valor := range s.Productos {
		if valor.Nombre == nombre {
			return valor, nil
		}
	}

	return Producto{}, ErrNotFound

}
