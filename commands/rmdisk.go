package commands

import (
	"fmt"
	"os"
)

// Funcion para eliminar un archivo que represente un disco duro
func RMDisk(path string, consola *string) {
	fmt.Println("Desea eliminar este disco duro?")
	fmt.Println("0 - Cancelar\n1 - Confirmar")
	var input int
	fmt.Scanln(&input)
	if input == 1 {
		err := os.Remove(path)
		if err != nil {
			fmt.Println("Error: El sistema no puede encontrar el archivo especificado.")
			*consola = *consola + "Error: El sistema no puede encontrar el archivo especificado.\n"
			return
		}
		fmt.Println("[-] Disco Eliminado con exito.")
		*consola = *consola + "[-] Disco Eliminado con exito.\n"
	} else {
		fmt.Println("[-] Cancelado")
		*consola = *consola + "[-] Cancelado\n"
	}

}
