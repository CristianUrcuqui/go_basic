package variables

import (
	"fmt"
	"strconv"
	"time"
)

// Variables universales publicas se deben declarar con la primera letra en Mayuscula ejemplo:
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Cristian"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now().UTC()
	fmt.Println("Nombre:", Nombre)
	fmt.Println("Estado:", Estado)
	fmt.Println("Sueldo:", Sueldo)
	fmt.Println("Fecha:", Fecha)
}

func ConviertoTexto(numero int) (bool, string) {
	text := strconv.Itoa(numero)
	return true, text
}
