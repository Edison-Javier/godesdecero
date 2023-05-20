package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["MX"] = "Mexico"
	paises["CO"] = "Colombia"
	paises["AR"] = "Argentina"

	fmt.Println(paises)
	fmt.Println(paises["CO"])

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Boca Junior": 100,
		"Junior":      45,
	}

	fmt.Println(campeonato)

	/*for equipo, puntaje := range campeonato {
		fmt.Println("Equipo:", equipo, "tiene un puntaje de:", puntaje)
	}*/

	delete(campeonato, "Real Madrid")

	fmt.Println(campeonato)

	puntaje, existe := campeonato["Junior"]
	fmt.Printf("El puntaje capturado es: %d. El equipo existe %t\n", puntaje, existe)
}
