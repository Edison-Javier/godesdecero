package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ObtenerTablaDeMultiplicar() string {
	var text string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Por favor ingrese un numero del 1 al 10")
		if scanner.Scan() {
			number, err := strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}

			if number > 0 && number <= 10 {
				for i := 1; i <= 10; i++ {
					text += fmt.Sprintln(number, "x", i, "=", number*i)
				}
				break
			}
		}
	}

	return text

}
