package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese numero: ")
		if scanner.Scan() {
			numero, err := strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}
			for i := 1; i < 11; i++ {
				fmt.Println(numero, " X ", i, " = ", numero*i)
			}
			fmt.Println("desea ingresar otro numero? (Y/N)")
			if scanner.Scan() {
				if scanner.Text() == "N" || scanner.Text() == "n" {
					break
				}
			}
		}
	}
}
