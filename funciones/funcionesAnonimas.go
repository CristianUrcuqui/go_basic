package funciones

import "fmt"

func Calculos() {
	var numero3 int = 32
	var numero4 int = 322

	suma := func(numero1, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4

	}

	fmt.Println(suma(10, 25))
}
