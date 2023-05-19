package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/godesdecero/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GuardarTabla() {
	tabla := ejercicios.ObtenerTablaDeMultiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo. Motivo: " + err.Error())
		return
	}

	fmt.Fprint(archivo, tabla)
	archivo.Close()
}

func AppendTextFile() {
	tabla := ejercicios.ObtenerTablaDeMultiplicar()
	if !Append(tabla) {
		fmt.Println("Fallo el Append Text File.")
	}
}

func Append(texto string) bool {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error al momento de guardar el archivo. Motivo:" + err.Error())
		return false
	}

	_, err = file.WriteString(texto)
	if err != nil {
		fmt.Println("Fallo el WriteString. Motivo:" + err.Error())
		return false
	}

	file.Close()
	return true
}

func LeerArchivo() {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Fallo lectura del archivo. Motivo: " + err.Error())
	}

	fmt.Println(string(file))
}

func LeerArchivo2() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Fallo lectura del archivo. Motivo: " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		println(">" + line)
	}

	file.Close()
}
