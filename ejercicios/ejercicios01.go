package ejercicios

import (
	"strconv"
)

func ConvertToInt(value string) (int, string) {
	number, error := strconv.Atoi(value)

	if error != nil {
		return 0, "Hubo un error. Motivo: " + error.Error()
	}

	if number > 100 {
		return number, "Es mayor a 100"
	} else {
		return number, "Es menor a 100"
	}
}
