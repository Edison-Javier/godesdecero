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

func TipoVariables() {
	Nombre = "Xavier"
	Estado = true
	Sueldo = 1750.80
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConvertToText(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}
