package ejercicios

import (
	"fmt"
	"strconv"
)

func ConvertToInt(value string) (int, string) {
	intValue, error := strconv.Atoi(value)

	var resp string
	if intValue > 100 {
		resp = "Es mayor a 100"
	} else {
		resp = "Es menor a 100"
	}

	fmt.Println(error)
	return intValue, resp
}
