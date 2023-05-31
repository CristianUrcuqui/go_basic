package ejercicio2

import (
	"bufio" // captura de datos por el usuario
	"fmt"
	"os"      // sistema operativo
	"strconv" // convertidor de tipo de datos
)

var err error
var numero int
var texto string

func Calculadora() string {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d\n", numero, i, i*numero)
	}

	return texto

}
