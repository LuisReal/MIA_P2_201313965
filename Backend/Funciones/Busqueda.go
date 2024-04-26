package Funciones

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InitSearch(path string, file *os.File, tempSuperblock Superblock) int32 {

	fmt.Println("======Start INITSEARCH======")
	//fmt.Println("path:", path)
	// path = "/ruta/nueva"

	// split the path by /
	TempStepsPath := strings.Split(path, "/")

	//fmt.Println("\nEl arreglo TempStepsPath es\n", TempStepsPath)

	StepsPath := TempStepsPath[1:]

	//fmt.Println("StepsPath:", StepsPath, "len(StepsPath):", len(StepsPath))
	/*
		for _, step := range StepsPath {
			fmt.Println("step:", step)
		}*/

	var Inode0 Inode
	// Read object from bin file
	if err := LeerObjeto(file, &Inode0, int64(tempSuperblock.S_inode_start)); err != nil {
		return -1
	}

	fmt.Println("======End INITSEARCH======")

	return SarchInodeByPath(StepsPath, Inode0, file, tempSuperblock)
}

func pop(s *[]string) string {
	lastIndex := len(*s) - 1
	last := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return last
}

// login -user=root -pass=123 -id=A119
func SarchInodeByPath(StepsPath []string, Inode_ Inode, file *os.File, tempSuperblock Superblock) int32 {
	fmt.Println("======Start SARCHINODEBYPATH======")
	index := int32(0)
	SearchedName := strings.Replace(pop(&StepsPath), " ", "", -1)

	//fmt.Println("========== SearchedName:", SearchedName)

	// Iterate over i_blocks from Inode
	for _, block := range Inode_.I_block {
		if block != -1 {
			if index < 13 {
				//CASO DIRECTO

				var crrFolderBlock Folderblock
				// Read object from bin file                                       // un Inodo y un bloque las estructuras miden lo mismo
				if err := LeerObjeto(file, &crrFolderBlock, int64(tempSuperblock.S_block_start+block*int32(binary.Size(Inode{})))); err != nil {
					return -1
				}

				for _, folder := range crrFolderBlock.B_content {
					// fmt.Println("Folder found======")
					//fmt.Println("Folder === Name:", string(folder.B_name[:]), "B_inodo", folder.B_inodo)

					if strings.Contains(string(folder.B_name[:]), SearchedName) {

						//fmt.Println("len(StepsPath)", len(StepsPath), "StepsPath", StepsPath)
						if len(StepsPath) == 0 {
							//fmt.Println("Folder found======")
							return folder.B_inodo
						} else {
							fmt.Println("NextInode======")
							var NextInode Inode
							// Read object from bin file
							if err := LeerObjeto(file, &NextInode, int64(tempSuperblock.S_inode_start+folder.B_inodo*int32(binary.Size(Inode{})))); err != nil {
								return -1
							}
							return SarchInodeByPath(StepsPath, NextInode, file, tempSuperblock)
						}
					}
				}

			} else {
				//CASO INDIRECTO
			}
		}
		index++
	}

	fmt.Println("======End SARCHINODEBYPATH======")

	return 0
}

//busca el id del grupo al que pertenece el usuario

func SearchByUser(grupo string, Inodo Inode, file *os.File, tempSuperblock Superblock) (string, error) {

	fmt.Println("=================== Inicio SearchByUser ===================")

	var bloque int
	var indice int
	var fileblock Fileblock
	var cadena string
	var fileblock_start int32

	for i := 0; i < len(Inodo.I_block); i++ { //iterando bloques de inodo1

		if Inodo.I_block[i] != -1 {

			bloque = int(Inodo.I_block[i]) //obtiene el numero del ultimo bloque de archivos creado
			indice = i

			fileblock_start = tempSuperblock.S_block_start + int32(bloque)*int32(binary.Size(Fileblock{}))

			if err := LeerObjeto(file, &fileblock, int64(fileblock_start)); err != nil { //bloque1
				return "", err
			}

			cadena += string(fileblock.B_content[:])

		}
	}

	fmt.Printf("\nel ultimo bloque creado es: %d, index: %d", bloque, indice)

	fmt.Println("\n Imprimiendo cadena\n", cadena)

	lines := strings.Split(cadena, "\n")

	if len(lines) > 0 {
		lines = lines[:len(lines)-1]
	}

	for i := 0; i < len(lines); i++ {

		datos := strings.Split(lines[i], ",")

		t := strings.TrimSpace(datos[0]) // elimina espacios para poder ser leido correctamente

		sv, _ := strconv.Atoi(string(t)) // contiene el id del grupo

		//fmt.Println("\nsv : ", sv)
		id := strconv.Itoa(sv)

		if len(datos) > 2 {

			if string(datos[2]) == grupo {
				fmt.Println("=================== Fin SearchByUser ===================")
				return id, nil

			}
		}

	}

	fmt.Println("=================== Fin SearchByUser ===================")

	return "", nil
}

func SearchFolder(folder_name string) int32 {

	fmt.Println("=================== SearchFolder ===================")

	var folder_slice [12]byte
	copy(folder_slice[:], []byte(folder_name))

	id := User_.Id

	driveletter := string(id[0])

	file, err := AbrirArchivo("./archivos/" + strings.ToUpper(driveletter) + ".dsk")
	if err != nil {
		return 0
	}

	defer file.Close()

	var tempMBR MBR

	if err := LeerObjeto(file, &tempMBR, 0); err != nil {
		return 0
	}

	var index int = -1

	for i := 0; i < 4; i++ {
		if tempMBR.Mbr_partitions[i].Part_size != 0 {
			if strings.Contains(string(tempMBR.Mbr_partitions[i].Part_id[:]), strings.ToUpper(id)) {
				fmt.Println("\n****Particion Encontrada en SearchFolder*****")
				index = i
			}
		}
	}

	var tempSuperblock Superblock

	if err := LeerObjeto(file, &tempSuperblock, int64(tempMBR.Mbr_partitions[index].Part_start)); err != nil {
		fmt.Println(err)
	}

	Inodo_start := tempSuperblock.S_inode_start

	var Inodo Inode

	var inodo_vacio Inode

	inodos_ocupados := tempSuperblock.S_inodes_count - tempSuperblock.S_free_inodes_count

	for i := 0; i < int(inodos_ocupados); i++ {

		if err := LeerObjeto(file, &Inodo, int64(Inodo_start+int32(i)*int32(binary.Size(Inode{})))); err != nil {
			fmt.Println(err)
		}

		if Inodo == inodo_vacio {
			continue
		}

		for _, block := range Inodo.I_block {
			if block != -1 {

				// carpeta -> 0   archivo ->1

				if string(Inodo.I_type[:]) == "0" { // si es carpeta

					var crrFolderBlock Folderblock

					if err := LeerObjeto(file, &crrFolderBlock, int64(tempSuperblock.S_block_start+block*int32(binary.Size(Folderblock{})))); err != nil {
						fmt.Println(err)
					}

					for _, folder := range crrFolderBlock.B_content {

						if folder.B_name == folder_slice {
							return folder.B_inodo // retorna el numero de inodo
						}

					}

				}

			}

		}

	}

	return 0
}
