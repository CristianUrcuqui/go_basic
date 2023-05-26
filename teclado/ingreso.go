package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numerouno int
var numerodos int
var leyenda string
var err error

func IngresoNumero() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese Numero 1: ")
	if scanner.Scan() {
		numerouno, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Dato ingresado incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese Numero 2: ")
	if scanner.Scan() {
		numerodos, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Dato ingresado incorrecto: " + err.Error())
		}
	}

	fmt.Println("Ingrese Leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numerouno*numerodos)

}
