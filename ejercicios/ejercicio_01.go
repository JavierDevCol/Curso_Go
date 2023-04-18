package ejercicios

import (
	"strconv"
)

func StringToInt(numeroText string) (int, string) {

	numeroInt, error := strconv.Atoi(numeroText)
	if error != nil {
		return numeroInt, "error"
	}
	if numeroInt > 100 {
		return numeroInt, "numero mayor a 100"
	} else {
		return numeroInt, "numero menor a 100"
	}

}
