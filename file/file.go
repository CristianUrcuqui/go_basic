package file

import (
	"bufio"
	"fmt"
	"go_basic/ejercicio2"
	"os"
)

var fileName string = "./file/archivos/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicio2.Calculadora()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating : " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicio2.Calculadora()
	if !AppendTabla(fileName, texto) {
		fmt.Printf("Error creating")
	}

}

func AppendTabla(filen, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error " + err.Error())
		return false
	}

	arch.Close()
	return true
}

// Leer con ioutil
/*
func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error" + err.Error())
		return
	}

	fmt.Println(string(archivo))
}
*/
// Leer con for
func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
