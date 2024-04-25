package Funciones

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Mkfs(id string, type_ string, fs_ string) string {

	datos := ""

	id = strings.ToUpper(id)

	//fmt.Println("\n\n=========================Iniciando MKFS===========================")
	//fmt.Println()

	datos += "\n\n=========================Iniciando MKFS==========================="

	driveletter := string(id[0])

	// Open bin file
	filepath := "./archivos/" + driveletter + ".dsk"
	file, err := AbrirArchivo(filepath)
	if err != nil {
		return ""
	}

	var TempMBR MBR
	// Read object from bin file
	if err := LeerObjeto(file, &TempMBR, 0); err != nil {
		return ""
	}

	// Print object
	PrintMBR(TempMBR)

	var index int = -1
	// Iterate over the partitions
	for i := 0; i < 4; i++ {
		if TempMBR.Mbr_partitions[i].Part_size != 0 {
			if strings.Contains(string(TempMBR.Mbr_partitions[i].Part_id[:]), id) {
				//fmt.Println("\n*********************Particion Encontrada****************")
				datos += "\n*********************Particion Encontrada****************"

				if TempMBR.Mbr_partitions[i].Part_status { // si la particion es true es porque esta montada
					//fmt.Println("\n ********************La Particion esta montada**********************")
					datos += "\n ********************La Particion esta montada**********************"
					index = i
				} else {
					//fmt.Println("\n ********************La Particion no esta montada**********************")
					datos += "\n ********************La Particion no esta montada**********************"
					return datos
				}
				break
			}
		}
	}

	if index != -1 {
		ImprimirParticion(TempMBR.Mbr_partitions[index])
	} else {
		//fmt.Println("\n*********************Particion NO Encontrada****************")
		datos += "\n*********************Particion NO Encontrada****************"
		return datos
	}

	numerador := int32(TempMBR.Mbr_partitions[index].Part_size - int32(binary.Size(Superblock{})))
	denominador_base := int32(4 + int32(binary.Size(Inode{})) + 3*int32(binary.Size(Fileblock{})))

	var temp int32 = 0
	if fs_ == "2fs" {
		temp = 0
	} else if fs_ == "3fs" {

		temp = int32(binary.Size(Journaling{}))

		//fmt.Println("\n *******************El tamano del journaling (en bytes) es: ", temp)
		datos += "\n *******************El tamano del journaling (en bytes) es: " + strconv.Itoa(int(temp))

	} else {
		//fmt.Println("\nIngrese un sistema de archivo correcto")
		datos += "\nIngrese un sistema de archivo correcto"

	}
	denominador := denominador_base + temp

	//fmt.Println("\nEl valor del numerador es: ", numerador)
	//fmt.Println("\nEl valor del denominador es: ", denominador)
	n := int32(numerador / denominador)

	//fmt.Println("\n*************************El numero de estructuras N es: ", n)
	datos += "\n*************************El numero de estructuras N es: " + strconv.Itoa(int(n))

	// var newMRB Structs.MRB
	var newSuperblock Superblock
	newSuperblock.S_inodes_count = n
	newSuperblock.S_blocks_count = 3 * n

	newSuperblock.S_free_blocks_count = 3 * n
	newSuperblock.S_free_inodes_count = n

	//copy(newSuperblock.S_mtime[:], "06/03/2024")      Esto no se evaluara
	//copy(newSuperblock.S_umtime[:], "06/03/2024")		Esto no se evaluara
	//newSuperblock.S_mnt_count = 0                    (No se evaluara cuantas veces fue montado el sistema)

	if fs_ == "2fs" {
		datos += ext2(n, TempMBR.Mbr_partitions[index], newSuperblock, file)
	} else {
		ext3(n, TempMBR.Mbr_partitions[index], newSuperblock, file)
	}

	// Close bin file
	defer file.Close()

	//fmt.Println("\n\n=========================Finalizando MKFS===========================")

	datos += "\n\n=========================Finalizando MKFS==========================="

	return datos
}

func ext2(n int32, partition Partition, newSuperblock Superblock, file *os.File) string {
	datos := ""
	//fmt.Println("\n\n=========================Creando ext2===========================")
	datos += "\n\n=========================Creando ext2==========================="

	newSuperblock.S_filesystem_type = 2
	newSuperblock.S_bm_inode_start = partition.Part_start + int32(binary.Size(Superblock{}))
	newSuperblock.S_bm_block_start = newSuperblock.S_bm_inode_start + n
	newSuperblock.S_inode_start = newSuperblock.S_bm_block_start + 3*n
	newSuperblock.S_block_start = newSuperblock.S_inode_start + n*int32(binary.Size(Inode{})) // n = numero de inodos
	newSuperblock.S_inode_size = int32(binary.Size(Inode{}))
	newSuperblock.S_block_size = int32(binary.Size(Folderblock{}))

	// se resta dos veces -1 porque hay que crear dos inodos y dos bloques
	newSuperblock.S_free_inodes_count -= 1
	newSuperblock.S_free_blocks_count -= 1
	newSuperblock.S_free_inodes_count -= 1
	newSuperblock.S_free_blocks_count -= 1

	// escribiendo ceros en bitmap de inodos
	for i := int32(0); i < n; i++ {
		err := EscribirObjeto(file, byte(0), int64(newSuperblock.S_bm_inode_start+i))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	// escribiendo ceros en bitmap de bloques
	for i := int32(0); i < 3*n; i++ {
		err := EscribirObjeto(file, byte(0), int64(newSuperblock.S_bm_block_start+i))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	// llenando con -1 los primeros 15 bloques
	var newInode Inode
	for i := int32(0); i < 15; i++ {
		newInode.I_block[i] = -1 //-1 no han sido utilizados
	}

	// escribiendo los inodos vacios en el archivo

	for i := int32(0); i < n; i++ {
		err := EscribirObjeto(file, newInode, int64(newSuperblock.S_inode_start+i*int32(binary.Size(Inode{}))))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	// escribiendo los bloques vacios en el archivo (los fileblock y folderblock tiene el mismo tamano de 64bytes)

	var newFileblock Fileblock
	for i := int32(0); i < 3*n; i++ {
		err := EscribirObjeto(file, newFileblock, int64(newSuperblock.S_block_start+i*int32(binary.Size(Fileblock{}))))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	//creando el primer Inode en posicion 0
	var Inode0 Inode //Inode 0
	Inode0.I_uid = 1
	Inode0.I_gid = 1
	Inode0.I_size = int32(binary.Size(Folderblock{}))

	// carpeta -> 0   archivo ->1
	copy(Inode0.I_type[:], "0")

	copy(Inode0.I_perm[:], "664")

	for i := int32(0); i < 15; i++ {
		Inode0.I_block[i] = -1
	}

	Inode0.I_block[0] = 0 // el inode 0 apunta al bloque 0

	// . | 0
	// .. | 0
	// user.txt | 1
	//

	var Folderblock0 Folderblock //Bloque 0 -> carpetas

	for i := int32(0); i < 4; i++ {
		Folderblock0.B_content[i].B_inodo = -1
	}

	Folderblock0.B_content[0].B_inodo = 0
	copy(Folderblock0.B_content[0].B_name[:], ".")
	Folderblock0.B_content[1].B_inodo = 0
	copy(Folderblock0.B_content[1].B_name[:], "..")
	Folderblock0.B_content[2].B_inodo = 1
	copy(Folderblock0.B_content[2].B_name[:], "user.txt")

	//creando el Inode 1

	var Inode1 Inode //Inode 1
	Inode1.I_uid = 1
	Inode1.I_gid = 1
	Inode1.I_size = int32(binary.Size(Fileblock{}))
	copy(Inode1.I_type[:], "1") // es de tipo archivo
	copy(Inode1.I_perm[:], "664")

	for i := int32(0); i < 15; i++ {
		Inode1.I_block[i] = -1
	}

	Inode1.I_block[0] = 1 // el Inode 1 apunta al bloque 1

	/*var name_bytes [16]byte
	copy(name_bytes[:], []byte(name))
	*/
	data := "1,G,root\n1,U,root,root,123\n"

	var Fileblock1 Fileblock //Bloque 1 -> archivo

	copy(Fileblock1.B_content[:], []byte(data))

	// Inodo 0 -> Bloque 0 -> Inodo 1 -> Bloque 1
	// Crear la carpeta raiz /
	// Crear el archivo user.txt "1,G,root\n1,U,root,root,123\n"

	// escribiendo el superblock
	err := EscribirObjeto(file, newSuperblock, int64(partition.Part_start))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// escribiendo bitmap inodes con unos

	err = EscribirObjeto(file, byte(1), int64(newSuperblock.S_bm_inode_start))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = EscribirObjeto(file, byte(1), int64(newSuperblock.S_bm_inode_start+1))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// escribiendo bitmap blocks con unos
	err = EscribirObjeto(file, byte(1), int64(newSuperblock.S_bm_block_start))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = EscribirObjeto(file, byte(1), int64(newSuperblock.S_bm_block_start+1))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	//fmt.Println("Inode 0:", int64(newSuperblock.S_inode_start))
	datos += "\nInode 0:" + strconv.Itoa(int(int64(newSuperblock.S_inode_start)))
	//fmt.Println("Inode 1:", int64(newSuperblock.S_inode_start+int32(binary.Size(Inode{}))))
	datos += "\nInode 1:" + strconv.Itoa(int(int64(newSuperblock.S_inode_start+int32(binary.Size(Inode{})))))

	// escribiendo inodes
	err = EscribirObjeto(file, Inode0, int64(newSuperblock.S_inode_start)) //Inode 0
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = EscribirObjeto(file, Inode1, int64(newSuperblock.S_inode_start+int32(binary.Size(Inode{})))) //Inode 1
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// escribiendo blocks
	err = EscribirObjeto(file, Folderblock0, int64(newSuperblock.S_block_start)) //Bloque 0
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = EscribirObjeto(file, Fileblock1, int64(newSuperblock.S_block_start+int32(binary.Size(Fileblock{})))) //Bloque 1

	if err != nil {
		fmt.Println("Error: ", err)
	}

	//fmt.Println("\n\n=========================Finalizando ext2===========================")
	datos += "\n\n=========================Finalizando ext2==========================="

	return datos
}

func ext3(n int32, partition Partition, newSuperblock Superblock, file *os.File) {

	fmt.Println("\n\n=========================Creando ext3===========================")

	newSuperblock.S_filesystem_type = 3
	newSuperblock.S_bm_inode_start = partition.Part_start + int32(binary.Size(Superblock{})) + int32(binary.Size(Journaling{}))
	newSuperblock.S_bm_block_start = newSuperblock.S_bm_inode_start + n
	newSuperblock.S_inode_start = newSuperblock.S_bm_block_start + 3*n
	newSuperblock.S_block_start = newSuperblock.S_inode_start + n*int32(binary.Size(Inode{})) // n = numero de inodos

	// se resta dos veces -1 porque hay que crear dos inodos y dos bloques al inicio
	fmt.Println("El numero de free_inodes_count es: ", newSuperblock.S_free_inodes_count)
	fmt.Println("El numero de free_blocks_count es: ", newSuperblock.S_free_blocks_count)
	newSuperblock.S_free_inodes_count -= 1
	newSuperblock.S_free_blocks_count -= 1
	newSuperblock.S_free_inodes_count -= 1
	newSuperblock.S_free_blocks_count -= 1

	//escribiendo el journaling
	var journaling Journaling

	/*
		var operacion [10]byte
		var path [100]byte
		var content [100]byte


		copy(operacion[:], "mkdir")
		copy(path[:], "/")
		copy(content[:], "Mi nombre es Luis Gonzalez")
		journaling.Contenido[0].Operation = operacion
		journaling.Contenido[0].Path = path
		journaling.Contenido[0].Content = content*/

	error_ := EscribirObjeto(file, journaling, int64(partition.Part_start+int32(binary.Size(Superblock{}))))

	if error_ != nil {
		fmt.Println("Error: ", error_)
	}

	// escribiendo ceros en bitmap de inodos
	for i := int32(0); i < n; i++ {
		err := EscribirObjeto(file, byte(0), int64(newSuperblock.S_bm_inode_start+i))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	// escribiendo ceros en bitmap de bloques
	for i := int32(0); i < 3*n; i++ {
		err := EscribirObjeto(file, byte(0), int64(newSuperblock.S_bm_block_start+i))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	// llenando con -1 los primeros 15 bloques
	var newInode Inode
	for i := int32(0); i < 15; i++ {
		newInode.I_block[i] = -1 //-1 no han sido utilizados
	}

	// escribiendo los inodos en el archivo

	for i := int32(0); i < n; i++ {
		err := EscribirObjeto(file, newInode, int64(newSuperblock.S_inode_start+i*int32(binary.Size(Inode{}))))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	// escribiendo los bloques en el archivo (los fileblock y folderblock tiene el mismo tamano de 64bytes)

	var newFileblock Fileblock
	for i := int32(0); i < 3*n; i++ {
		err := EscribirObjeto(file, newFileblock, int64(newSuperblock.S_block_start+i*int32(binary.Size(Fileblock{}))))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	//creando el primer Inode en posicion 0
	var Inode0 Inode //Inode 0
	Inode0.I_uid = 1
	Inode0.I_gid = 1
	Inode0.I_size = int32(binary.Size(Folderblock{}))

	copy(Inode0.I_type[:], "0") // es de tipo carpeta
	copy(Inode0.I_perm[:], "664")

	for i := int32(0); i < 15; i++ {
		Inode0.I_block[i] = -1
	}

	Inode0.I_block[0] = 0 // el inode 0 apunta al bloque 0

	// . | 0
	// .. | 0
	// user.txt | 1
	//

	var Folderblock0 Folderblock //Bloque 0 -> carpetas

	for i := int32(0); i < 4; i++ {
		Folderblock0.B_content[i].B_inodo = -1
	}

	Folderblock0.B_content[0].B_inodo = 0
	copy(Folderblock0.B_content[0].B_name[:], ".")
	Folderblock0.B_content[1].B_inodo = 0
	copy(Folderblock0.B_content[1].B_name[:], "..")
	Folderblock0.B_content[2].B_inodo = 1
	copy(Folderblock0.B_content[2].B_name[:], "user.txt")

	//creando el Inode 1

	var Inode1 Inode //Inode 1
	Inode1.I_uid = 1
	Inode1.I_gid = 1
	Inode1.I_size = int32(binary.Size(Fileblock{}))

	copy(Inode1.I_type[:], "1") // es de tipo archivo
	copy(Inode1.I_perm[:], "664")

	for i := int32(0); i < 15; i++ {
		Inode1.I_block[i] = -1
	}

	Inode1.I_block[0] = 1 // el Inode 1 apunta al bloque 1

	data := "1,G,root\n1,U,root,root,123\n"
	var Fileblock1 Fileblock //Bloque 1 -> archivo
	copy(Fileblock1.B_content[:], data)

	// Inodo 0 -> Bloque 0 -> Inodo 1 -> Bloque 1
	// Crear la carpeta raiz /
	// Crear el archivo user.txt "1,G,root\n1,U,root,root,123\n"

	// escribiendo el superblock
	err := EscribirObjeto(file, newSuperblock, int64(partition.Part_start))

	if err != nil {
		fmt.Println("Error: ", err)
	}
	// escribiendo bitmap inodes con unos

	err = EscribirObjeto(file, byte(1), int64(newSuperblock.S_bm_inode_start))
	if err != nil {
		fmt.Println("Error: ", err)
	}
	err = EscribirObjeto(file, byte(1), int64(newSuperblock.S_bm_inode_start+1))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// escribiendo bitmap blocks con unos
	err = EscribirObjeto(file, byte(1), int64(newSuperblock.S_bm_block_start))
	if err != nil {
		fmt.Println("Error: ", err)
	}
	err = EscribirObjeto(file, byte(1), int64(newSuperblock.S_bm_block_start+1))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("\n**********Inode 0 inicia en: ", int64(newSuperblock.S_inode_start))
	fmt.Println("\n**********Inode 1 inicia en: ", int64(newSuperblock.S_inode_start+int32(binary.Size(Inode{}))))

	// escribiendo inodes
	err = EscribirObjeto(file, Inode0, int64(newSuperblock.S_inode_start)) //Inode 0
	if err != nil {
		fmt.Println("Error: ", err)
	}
	err = EscribirObjeto(file, Inode1, int64(newSuperblock.S_inode_start+int32(binary.Size(Inode{})))) //Inode 1
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// escribiendo blocks
	err = EscribirObjeto(file, Folderblock0, int64(newSuperblock.S_block_start)) //Bloque 0
	if err != nil {
		fmt.Println("Error: ", err)
	}
	err = EscribirObjeto(file, Fileblock1, int64(newSuperblock.S_block_start+int32(binary.Size(Fileblock{})))) //Bloque 1

	if err != nil {
		fmt.Println("Error: ", err)
	}

	var tempjournaling Journaling

	error2_ := LeerObjeto(file, &tempjournaling, int64(partition.Part_start+int32(binary.Size(Superblock{}))))

	if error2_ != nil {
		fmt.Println("Error: ", err)
	}

	PrintJournaling(tempjournaling)

	var tempSuperblock Superblock

	error3_ := LeerObjeto(file, &tempSuperblock, int64(partition.Part_start))

	if error3_ != nil {
		fmt.Println("Error: ", err)
	}

	PrintSuperblock(tempSuperblock)

	fmt.Println("\n\n=========================Finalizando ext3===========================")

}
