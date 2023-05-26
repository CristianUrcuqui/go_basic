package ejercicios

import (
	"strconv"
)

func EjercicioUno(numero string) (int, string) {
	text, err := strconv.Atoi(numero)
	if err != nil {
		return 0, "Invalid" + err.Error()
	}
	if text > 100 {
		return text, numero
	} else {
		return text, "Es menor a 100"
	}
}
