package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Javier"
	Estado = true
	Sueldo = 6700000.55
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func CovnertirEnteroToString(number int) (bool, string) {
	texto := strconv.Itoa(number)
	return true, texto
}
