package commands

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"strconv"
)

// Estructura de un objeto mount que representara la particon montada Primaria
type mountedParts struct {
	partition interface{}
	number    int
	id        string
}

// Estructura de un objeto mount que representara todas las particiones de un disco que sean montadas
type mounted struct {
	path   string
	letter byte
	parts  []mountedParts
}

// Lista de los discos que han sido montados con sus respectivos
var mountedDisks []mounted

// Funcion Mount, añade particiones que fueron montadas por el comando
func Mount(path string, name string, consola *string) {
	//Obtenemos el mbr del disco
	file, mbr, err := readFile(path)
	if err != nil {
		return
	}
	defer file.Close()
	//Indice de la particion encontrada
	index := -1
	//Letra que se asigna por disco
	var diskLetter byte = 'a'
	//Disco montando anteriormente para usarlo temporalmente
	disk := mounted{}
	//Buscamos en la lista de discos que fueron montados
	for i, disks := range mountedDisks {
		if path == disks.path {
			//Se encontro el disco que alguna vez fue montado
			diskLetter, disk, index = disks.letter, disks, i
			break
		}
		diskLetter = byte(int(diskLetter) + 1)
	}
	//	Obtenemos el nombre a asignar
	var realName [16]byte
	copy(realName[:], name)
	//Verificamos que no haya sido montada esta particion
	if existPart(disk, realName, consola) {
		fmt.Println("Alert: Esta particion ya ha sido montada")
		*consola = *consola + "Alert: Esta particion ya ha sido montada\n"
		return
	}
	//Variable que almacena el estado si se encontro o no la particion
	findPart, findEBR := fmt.Errorf("NOT FOUND"), fmt.Errorf("NOT FOUND")
	//Variable que almacena temporalmente la particion encontrada
	partition := partition{}
	ebr := extendedBootRecord{}
	//Recorremos la lista de particiones del MBR
	for _, part := range mbr.Partitions {
		if part.Name == realName {
			//Se encontro la particion
			findPart, partition = nil, part
			break
		}
		//Buscamos en las particiones logicas
		if part.Type == 'E' {
			indexEBR := part.Start
			for i := 1; true; i++ {
				file.Seek(indexEBR, 0)
				//Se obtiene la data del archivo binario
				data := readNextBytes(file, int64(binary.Size(ebr)))
				buffer := bytes.NewBuffer(data)
				err := binary.Read(buffer, binary.BigEndian, &ebr)
				if err != nil {
					log.Fatal("binary.Read failed", err)
				}
				//Verificamos si existe
				if ebr.Name == realName {
					findPart = nil
					findEBR = nil
					break
				}
				//Si ya no hay siguientes
				if ebr.Next == -1 {
					break
				}
				indexEBR = ebr.Next
			}
		}
	}
	//En dado caso no se encuentre la particion
	if findPart != nil {
		fmt.Println("[ERROR] Esta particion no fue encontrada en el disco.")
		*consola = *consola + "[ERROR] Esta particion no fue encontrada en el disco.\n"
		return
	}
	//Las particiones extendidas no pueden ser montadas
	if partition.Type == 'E' {
		fmt.Println("[ERROR] No es permitido montar una particion extendida.")
		*consola = *consola + "[ERROR] No es permitido montar una particion extendida.\n"
		return
	}
	//Si el disco no ha sido montado lo montamos
	if disk.path == "" {
		disk.letter, disk.path = diskLetter, path
	}
	//Creamos la nueva particion a montar
	part := mountedParts{}
	//En dado caso vamos a montar una particion logica
	if findEBR == nil {
		part.partition = ebr
	} else {
		part.partition = partition
	}
	if len(disk.parts) == 0 {
		part.number = 1
	} else {
		part.number = disk.parts[len(disk.parts)-1].number + 1
	}
	part.id = "74" + strconv.Itoa(part.number) + string(disk.letter)
	fmt.Println(part.id)
	//La montamos al disco
	disk.parts = append(disk.parts, part)
	//Montamos el disco a la lista de discos montados
	if index == -1 {
		mountedDisks = append(mountedDisks, disk)
	} else {
		mountedDisks[index] = disk
	}
	fmt.Println("[-] La particion ha sido montada con exito.")
	*consola = *consola + "[-] La particion ha sido montada con exito.\n"
}

func existPart(disk mounted, name [16]byte, consola *string) bool {
	//Revisamos si no existe en las primarias o extendidas
	for _, part := range disk.parts {
		typePart := typeOf(part.partition, consola)
		switch typePart {
		case 0:
			aux := part.partition.(partition)
			if aux.Name == name {
				return true
			}
		case 1:
			aux := part.partition.(extendedBootRecord)
			if aux.Name == name {
				return true
			}
		}
	}
	return false
}

func typeOf(x interface{}, consola *string) int {
	// type switch
	switch x.(type) {
	case partition:
		return 0
	case extendedBootRecord:
		return 1
	default:
		fmt.Println("Error: No se encontro el tipo de particion")
		*consola = *consola + "Error: No se encontro el tipo de particion\n"
		return -1
	}
}

// Funcion para mostrar todos las particiones montadas en el sistema
func ShowMountedDisks(consola *string) {
	fmt.Println("[-] Particiones montadas:")
	//Recorremos la lista de discos montados
	for _, disk := range mountedDisks {
		//Path del disco temporal
		path := disk.path
		//Recorremos las particiones de ese disco que han sido montadas
		for i, part := range disk.parts {
			partitionType := typeOf(part.partition, consola)
			switch partitionType {
			case 0:
				aux := part.partition.(partition)
				fmt.Printf("	[%d] -id = %s -path = \"%s\" -name = \"%s\"\n", i+1, part.id, path, aux.Name)
				*consola = *consola + "	[%d] -id = %s -path = \"%s\" -name = \"%s\"\n" + string(i+1) + part.id + path
			case 1:
				aux := part.partition.(extendedBootRecord)
				fmt.Printf("	[%d] -id = %s -path = \"%s\" -name = \"%s\"\n", i+1, part.id, path, aux.Name)
				*consola = *consola + "	[%d] -id = %s -path = \"%s\" -name = \"%s\"\n" + string(i+1) + part.id + path
			}
		}
	}
}
