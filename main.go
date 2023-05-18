package main

import (
	"fmt"

	"github.com/godesdecero/variables"
)

func main() {
	//variables.TipoVariables()
	status, text := variables.ConvertToText(1700)

	fmt.Println(status)
	fmt.Println(text)
}
