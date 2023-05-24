package main

import (
	"fmt"
	"go_basic/variables"
)

func main() {
	estado, text := variables.ConviertoTexto(123)
	fmt.Println(estado)
	fmt.Println(text)
}
