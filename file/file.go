package file

import (
	"bufio"
	"fmt"
	"goDesdeCero/ejercicios"
	"os"
)

var ruta_archivo string = "./file/txt/ejercicioManejoArchivos.txt"

func CrearArchivoConTablaMultiplicar() {

	var text = ejercicios.TablaMultiplicar()
	archivo, err := os.Create(ruta_archivo)
	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error())
		return
	}
	fmt.Fprintln(archivo, text)
	archivo.Close()
}

func AgregarTablaAlArchivo() {
	var text string = ejercicios.TablaMultiplicar()
	if !Append(ruta_archivo, text) {
		fmt.Println("Error al concatenar contenido ")
	}
}

func Append(archivo string, text string) bool {
	arch, err := os.OpenFile(ruta_archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}
	_, err = arch.WriteString(text)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(ruta_archivo)
	if err != nil {
		fmt.Println("Error al leer el archivo")
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> ", registro)
	}
}
