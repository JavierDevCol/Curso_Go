package main

import (
	"fmt"
	"goDesdeCero/condicionales"
	"goDesdeCero/ejercicios"
	"goDesdeCero/variables"
	"runtime"
)

func main() {
	estado, texto := variables.CovnertirEnteroToString(6700000)
	fmt.Println(estado)
	fmt.Println(texto)
	fmt.Println(condicionales.EsWindows(runtime.GOOS))
	fmt.Println(ejercicios.StringToInt("19u"))
}
