package main

import (
	"fmt"
	"goDesdeCero/variables"
)

func main() {
	estado, texto := variables.CovnertirEnteroToString(6700000)
	fmt.Println(estado)
	fmt.Println(texto)
}
