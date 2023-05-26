package ejercicio2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var err error
var numero int

func Calculadora() {

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
		fmt.Printf("%d x %d = %d\n", numero, i, i*numero)
	}

}