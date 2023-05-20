package funciones

import "fmt"

func tabla(valor int) func(int) int {
	numero := valor

	return func(secuencia int) int {
		return numero * secuencia
	}
}

func LlamarClosure() {
	tabladel := 2
	multiplicar := tabla(tabladel)

	for i := 1; i <= 10; i++ {
		fmt.Println(tabladel, "x", i, "=", multiplicar(i))
	}
}
