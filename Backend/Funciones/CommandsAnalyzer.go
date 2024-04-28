package Funciones

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var User_ User

var re = regexp.MustCompile(`-(\w+)=("[^"]+"|\S+)`)

var contador int = 0
var abecedario = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func getCommandAndParams(input string) (string, string) {

	if input != " " {

		parts := strings.Fields(input)

		//fmt.Println("\nImprimiendo parts: ", parts)
		if len(parts) > 0 {

			command := strings.ToLower(parts[0])
			params := strings.Join(parts[1:], " ")

			return command, params
		}
	}

	return "", input
}

var dot []string
var nombre_dot []string

func Analyze(input string) (string, []string, []string, User) {

	datos := ""
	var usuario User

	scanner := bufio.NewScanner(strings.NewReader(input))

	var newLine string

	for scanner.Scan() {
		//fmt.Println("Leyendo linea del input")
		//fmt.Println(scanner.Text())

		linea := scanner.Text()

		if linea != "" {

			for i := 0; i < len(linea); i++ {

				if string(linea[i]) == "#" {

					break
				}

				newLine += string(linea[i])
			}

			if newLine != "" {

				command, params := getCommandAndParams(newLine)

				newLine = ""

				fmt.Println("\n\nCommand: ", command, "Params: ", params)

				datos_, dot_, nombre_archivo_, usuario_ := AnalyzeComand(command, params)
				usuario = usuario_
				datos += datos_
				dot = append(dot, dot_)
				nombre_dot = append(nombre_dot, nombre_archivo_)

			}

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	//se valida ejecucion de comando execute

	return datos, dot, nombre_dot, usuario
}

func AnalyzeComand(command string, params string) (string, string, string, User) {
	datos := ""
	dot := ""
	nombre_archivo := ""
	var usuario User

	if command == "mkdisk" {
		datos += bn_mkdisk(params)
	} else if command == "rmdisk" {
		bn_rmdisk(params)
	} else if command == "fdisk" {
		datos += bn_fdisk(params)
	} else if command == "mount" {
		datos += bn_mount(params)
	} else if command == "unmount" {
		datos += bn_unmount(params)
	} else if command == "mkfs" {
		datos += bn_mkfs(params)
	} else if command == "login" {
		datos += bn_login(params)
	} else if command == "mkgrp" {
		bn_mkgrp(params)
	} else if command == "rmgrp" {
		bn_rmgrp(params)
	} else if command == "mkusr" {
		bn_mkusr(params)
	} else if command == "rmusr" {
		bn_rmusr(params)
	} else if command == "logout" {
		datos_, usuario_ := bn_logout()
		datos += datos_
		usuario = usuario_

	} else if command == "pause" {
		bn_pause()
	} else if command == "mkdir" {
		datos += bn_mkdir(params)
	} else if command == "mkfile" {
		datos += bn_mkfile(params)
	} else if command == "cat" {
		bn_cat(params)
	} else if command == "remove" {
		bn_remove(params)
	} else if command == "move" {
		bn_move(params)
	} else if command == "rep" {
		datos_, dot_, nombre_archivo_ := bn_reportes(params)
		nombre_archivo = nombre_archivo_
		datos += datos_
		dot += dot_
	} else {
		fmt.Println("Error: Command not found")
		datos += "Error: Command not found"
	}

	return datos, dot, nombre_archivo, usuario
}

func bn_move(params string) {

	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	fs := flag.NewFlagSet("move", flag.ExitOnError)
	path := fs.String("path", "", "path")
	dest := fs.String("dest", "", "destino")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := strings.ToLower(match[1])
		//flagValue := strings.ToLower(match[2])
		flagValue := match[2]
		flagValue = strings.Trim(flagValue, "\"") // elimina comillas si la ruta trae comillas

		switch flagName {
		case "path", "dest":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")

			return
		}
	}

	// Call the function
	Move(*path, *dest)
}

func bn_remove(params string) {

	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	fs := flag.NewFlagSet("remove", flag.ExitOnError)
	path := fs.String("path", "", "path")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := strings.ToLower(match[1])
		//flagValue := strings.ToLower(match[2])
		flagValue := match[2]
		flagValue = strings.Trim(flagValue, "\"") // elimina comillas si la ruta trae comillas

		switch flagName {
		case "path":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")

			return
		}
	}

	// Call the function
	Remove(*path)
}

func bn_cat(params string) {
	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	fs := flag.NewFlagSet("cat", flag.ExitOnError)
	file := fs.String("file", "", "File")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := strings.ToLower(match[1])
		//flagValue := strings.ToLower(match[2])
		flagValue := match[2]
		flagValue = strings.Trim(flagValue, "\"") // elimina comillas si la ruta trae comillas

		switch flagName {
		case "file":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")

			return
		}
	}

	// Call the function
	Cat(*file)
}

func bn_mkdir(params string) string { //mkdir -path=/bin
	datos := ""
	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	fs := flag.NewFlagSet("mkdir", flag.ExitOnError)
	path := fs.String("path", "", "Path")
	r := fs.String("r", "0", "r") // si viene un parametro r el valor seria una cadena vacia ("") y sino viene por defecto seria 0

	//execute -path=/home/darkun/Escritorio/prueba.mia

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		//flagValue := strings.ToLower(match[2])
		flagValue := match[2]
		fmt.Println("Flagvalue es: ", flagValue)
		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "r", "path":
			fs.Set(flagName, flagValue)
		default:
			//fmt.Println("Error: Flag not found")
			datos += "Error: Flag not found"

			return datos
		}
	}

	//execute -path=/home/darkun/Escritorio/prueba.mia
	slice_params := strings.Fields(params)
	//fmt.Println("EL slice_params es: ", slice_params)

	slice_path := strings.Fields(*path)
	//fmt.Println("EL slice_parmas es: ", slice_params)

	newOutput := strings.Join(slice_path, " ")
	//fmt.Println("EL newOuput es: ", newOutput)

	newInput := strings.Replace(newOutput, " ", "\"", -1) //reemplazando con comillas el espacio entre "archivos 19" por "archivos"19"
	//fmt.Println("El newInput es: ", newInput)
	fmt.Println("El newInput es: ", newInput)
	datos += "El newInput es: " + newInput

	newSlice := []string{slice_params[0], newInput}
	fmt.Println("newSlice es: ", newSlice)

	for i := 0; i < len(newSlice); i++ {

		if newSlice[i] == "-r" {
			datos += "Existe el parametro -r"
			*r = "1"

		}
	}

	datos += "\nEl valor de -r es: " + *r
	// Call the function
	data, _ := Mkdir(newInput, *r)

	datos += data

	return datos
}

func bn_mkfile(params string) string { //mkdir -path=/bin
	datos := ""
	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	fs := flag.NewFlagSet("mkfile", flag.ExitOnError)
	path := fs.String("path", "", "Path")
	r := fs.String("r", "0", "r") // si viene un parametro r el valor seria una cadena vacia ("") y sino viene por defecto seria 0
	size := fs.Int("size", 0, "Size")
	cont := fs.String("cont", "", "Cont")

	//execute -path=/home/darkun/Escritorio/prueba.mia

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		//flagValue := strings.ToLower(match[2])
		flagValue := match[2]
		//fmt.Println("Flagvalue es: ", flagValue)
		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "r", "path", "size", "cont":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
			datos += "Error: Flag not found"

			return datos
		}
	}

	//execute -path=/home/darkun/Escritorio/prueba.mia
	slice_params := strings.Fields(params)
	//fmt.Println("EL slice_params es: ", slice_params)

	slice_path := strings.Fields(*path)
	//fmt.Println("EL slice_parmas es: ", slice_params)

	newOutput := strings.Join(slice_path, " ")
	//fmt.Println("EL newOuput es: ", newOutput)

	newInput := strings.Replace(newOutput, " ", "\"", -1) //reemplazando con comillas el espacio entre "archivos 19" por "archivos"19"
	fmt.Println("El newInput es: ", newInput)
	datos += "El newInput es: " + newInput

	newSlice := []string{slice_params[0], newInput}
	fmt.Println("newSlice es: ", newSlice)

	for i := 0; i < len(newSlice); i++ {

		if newSlice[i] == "-r" {
			fmt.Println("Existe el parametro -r")
			*r = "1"

		}
	}

	fmt.Println("\nEl valor de -r es: ", *r)
	// Call the function
	data, _ := Mkfile(newInput, *r, *size, *cont)

	datos += data

	return datos
}

func bn_pause() {

	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	//read -n 1 -s -r -p "Press any key to continue"

	fmt.Println("\nPresione cualquier tecla para continuar")
	cmd := exec.Command("bash", "-c", "read -n 1 ")
	cmd.Stdin = os.Stdin
	out, err := cmd.CombinedOutput()

	fmt.Println("error:", err)
	fmt.Printf("output: %q\n", out)

}

func bn_reportes(params string) (string, string, string) {
	datos := ""
	dot := ""
	nombre_archivo := ""
	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	fs := flag.NewFlagSet("rep", flag.ExitOnError)
	name := fs.String("name", "", "Name")
	path := fs.String("path", "", "Path")
	id := fs.String("id", "m", "")
	ruta := fs.String("ruta", "", "ruta")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := strings.ToLower(match[1])
		//flagValue := strings.ToLower(match[2])
		flagValue := match[2]
		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "name", "path", "id", "ruta":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
			datos += "Error: Flag not found"
			return datos, "", ""
		}
	}

	// Call the function
	data, _, grafo, nombre_archivo_ := Reportes(*name, *path, *id, *ruta)
	nombre_archivo = nombre_archivo_
	datos += data
	dot += grafo

	return datos, dot, nombre_archivo
}

func bn_mkdisk(params string) string {

	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	// Define flags
	fs := flag.NewFlagSet("mkdisk", flag.ExitOnError)
	size := fs.Int("size", 0, "Tamaño")
	fit := fs.String("fit", "", "Ajuste")
	unit := fs.String("unit", "", "Unidad")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "size", "fit", "unit":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")

			return ""
		}
	}

	fmt.Println("\n       El valor del contador es: ", contador)

	// Call the function
	datos := Mkdisk(*size, *fit, *unit, abecedario[contador])
	contador++

	return datos
}

func bn_rmdisk(params string) {

	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	// Define flags
	fs := flag.NewFlagSet("rmdisk", flag.ExitOnError)
	letra := fs.String("driveletter", "", "Disco")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "driveletter":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")

			return
		}
	}

	// Call the function
	Rmdisk(*letra)

}

func bn_fdisk(input string) string {
	datos := ""
	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	// Define flags
	fs := flag.NewFlagSet("fdisk", flag.ExitOnError)
	size := fs.Int("size", -1, "Tamaño")
	driveletter := fs.String("driveletter", "", "Letra")
	name := fs.String("name", "", "Nombre")
	unit := fs.String("unit", "", "Unidad")
	type_ := fs.String("type", "", "Tipo")
	fit := fs.String("fit", "", "Ajuste")
	delete := fs.String("delete", "", "Elimina particion")
	add := fs.Int("add", 0, "Agrega espacio")

	input_ := strings.Split(input, " ")
	//fmt.Println("\nImprimiendo SLICE input: ", input_)
	var formateo string

	//fmt.Println("\nImprimendo input_[1]: ", input_[1])
	if input_[0] == "-delete=full" {
		formateo = "rapido"
	} else {

		for i := 1; i < len(input_); i++ {

			if input_[i] == "-delete=full" {
				formateo = "completo"
				break
			}
		}

	}

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)
	//fmt.Println("\nImprimiendo matches: ", matches)
	// Process the input

	for _, match := range matches {
		flagName := match[1]

		//fmt.Println("\nmatch[1]: ", match[1])
		//fmt.Println("\nmatch[2]: ", match[2])
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "size", "fit", "unit", "driveletter", "name", "type", "delete", "add":
			fs.Set(flagName, flagValue)
		default:
			//fmt.Println("Error: Flag not found")
			datos += "Error: Flag not found"

			return datos
		}
	}

	//Funciones.Fdisk(10, "A", "Particion1", "b", " ", "bf", "", 0)
	// Call the function

	//fmt.Println("\nImprimiendo valor de formateo: ", formateo)
	datos += Fdisk(*size, *driveletter, *name, *unit, *type_, *fit, *delete, *add, formateo)

	return datos
}

func bn_mount(input string) string {
	datos := ""
	//execute -path=/home/darkun/Escritorio/avanzado.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	// Define flags
	fs := flag.NewFlagSet("mount", flag.ExitOnError)
	driveletter := fs.String("driveletter", "", "Letra")
	name := fs.String("name", "", "Nombre")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "driveletter", "name":
			fs.Set(flagName, flagValue)
		default:
			//fmt.Println("Error: Flag not found")
			datos += "Error: Flag not found"
			return datos
		}
	}

	// Call the function
	datos += Mount(*driveletter, *name)

	return datos
}

func bn_unmount(input string) string {
	datos := ""

	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	// Define flags
	fs := flag.NewFlagSet("unmount", flag.ExitOnError)
	id := fs.String("id", "", "ID")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input

	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]

		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {

		case "id":
			fs.Set(flagName, flagValue)
		default:
			//fmt.Println("Error: Flag not found")
			datos += "Error: Flag not found"

			return datos
		}
	}

	// Call the function
	datos += UnMount(*id)

	return datos
}

func bn_mkfs(input string) string {
	datos := ""
	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	// Define flags
	fs := flag.NewFlagSet("mkfs", flag.ExitOnError)
	id := fs.String("id", "", "Id")
	type_ := fs.String("type", "", "Tipo")
	fs_ := fs.String("fs", "", "Fs")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "id", "type", "fs":
			fs.Set(flagName, flagValue)
		default:
			//fmt.Println("Error: Flag not found")
			datos += "Error: Flag not found"

			return datos
		}
	}

	// Call the function
	datos += Mkfs(*id, *type_, *fs_)

	return datos
}

func bn_login(input string) string {

	datos := ""

	//execute -path=/home/darkun/Escritorio/prueba.mia

	//execute -path=/home/darkun/Escritorio/basico.mia

	// Define flags
	fs := flag.NewFlagSet("login", flag.ExitOnError)
	user := fs.String("user", "", "Usuario")
	pass := fs.String("pass", "", "Contraseña")
	id := fs.String("id", "", "Id")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "user", "pass", "id":
			fs.Set(flagName, flagValue)
		default:
			//fmt.Println("Error: Flag not found")
			datos += "Error: Flag not found"
			return datos
		}
	}

	// Call the function
	dat, _ := Login(*user, *pass, *id)

	//fmt.Println("Devolviendo data: \n", dat)

	datos += dat

	/*EL usuario root puede ejecutar los siguientes comandos:
	MKGRP
	RMGRP
	MKUSR
	RMUSR
	*/

	return datos
}

func bn_logout() (string, User) {
	datos := ""
	var usuario User

	fmt.Println("\n\n========================= Iniciando Logout =========================")
	datos += "\n\n========================= Iniciando Logout ========================="

	if User_.Status {
		fmt.Println("\n Cerrando sesion de usuario: ", User_.Nombre)
		datos += "\n Cerrando sesion de usuario: " + User_.Nombre

		usuario.Nombre = User_.Nombre
		usuario.Status = User_.Status

		User_.Nombre = ""
		User_.Status = false
		User_.Id = ""
	} else {
		fmt.Println("\nERROR: No hay sesion actual")
		datos += "\nERROR: No hay sesion actual"
	}

	fmt.Println("\n\n========================= Finalizando Logout =========================")
	datos += "\n\n========================= Finalizando Logout ========================="

	return datos, usuario
}

func bn_mkgrp(input string) {

	//execute -path=/home/darkun/Escritorio/scripts.sdaa

	if User_.Nombre == "root" && User_.Status { //si el usuario es root y esta logueado(true)
		// Define flags
		fs := flag.NewFlagSet("mkgrp", flag.ExitOnError)
		name := fs.String("name", "", "nombre de grupo")

		// Parse the flags
		fs.Parse(os.Args[1:])

		// find the flags in the input
		matches := re.FindAllStringSubmatch(input, -1)

		// Process the input
		for _, match := range matches {
			flagName := match[1]
			flagValue := strings.ToLower(match[2])

			flagValue = strings.Trim(flagValue, "\"")

			switch flagName {
			case "name":
				fs.Set(flagName, flagValue)
			default:
				fmt.Println("Error: Flag not found")

				return
			}
		}

		Mkgrp(*name, User_.Id)

	} else {
		fmt.Println("\n\n******************Necesita iniciar sesion como ususario ROOT***********************")
		return
	}
}

func bn_rmgrp(input string) {

	//execute -path=/home/darkun/Escritorio/scripts.sdaa

	if User_.Nombre == "root" && User_.Status { //si el usuario es root y esta logueado(true)
		// Define flags
		fs := flag.NewFlagSet("rmgrp", flag.ExitOnError)
		name := fs.String("name", "", "nombre de grupo")

		// Parse the flags
		fs.Parse(os.Args[1:])

		// find the flags in the input
		matches := re.FindAllStringSubmatch(input, -1)

		// Process the input
		for _, match := range matches {
			flagName := match[1]
			flagValue := strings.ToLower(match[2])

			flagValue = strings.Trim(flagValue, "\"")

			switch flagName {
			case "name":
				fs.Set(flagName, flagValue)
			default:
				fmt.Println("Error: Flag not found")

				return
			}
		}

		Rmgrp(*name, User_.Id)

	} else {
		fmt.Println("\n\n******************Necesita iniciar sesion como ususario ROOT para poder REMOVER un grupo***********************")
		return
	}
}

func bn_mkusr(input string) {

	//execute -path=/home/darkun/Escritorio/scripts.sdaa

	if User_.Nombre == "root" && User_.Status { //si el usuario es root y esta logueado(true)
		// Define flags
		fs := flag.NewFlagSet("mkusr", flag.ExitOnError)
		user := fs.String("user", "", "nombre de usuario")
		pass := fs.String("pass", "", "contrasena de usuario")
		group := fs.String("grp", "", "grupo al que pertenecera el usuario")

		// Parse the flags
		fs.Parse(os.Args[1:])

		// find the flags in the input
		matches := re.FindAllStringSubmatch(input, -1)

		// Process the input
		for _, match := range matches {
			flagName := match[1]
			flagValue := strings.ToLower(match[2])

			flagValue = strings.Trim(flagValue, "\"")

			switch flagName {
			case "user", "pass", "grp":
				fs.Set(flagName, flagValue)
			default:
				fmt.Println("Error: Flag not found")

				return
			}
		}

		Mkusr(*user, *pass, *group, User_.Id)

	} else {
		fmt.Println("\n\n******************Necesita iniciar sesion como ususario ROOT para poder crear un usuario ***********************")
		return
	}
}

func bn_rmusr(input string) {

	//execute -path=/home/darkun/Escritorio/scripts.sdaa

	if User_.Nombre == "root" && User_.Status { //si el usuario es root y esta logueado(true)
		// Define flags
		fs := flag.NewFlagSet("rmusr", flag.ExitOnError)
		user := fs.String("user", "", "nombre de usuario")

		// Parse the flags
		fs.Parse(os.Args[1:])

		// find the flags in the input
		matches := re.FindAllStringSubmatch(input, -1)

		// Process the input
		for _, match := range matches {
			flagName := match[1]
			flagValue := strings.ToLower(match[2])

			flagValue = strings.Trim(flagValue, "\"")

			switch flagName {
			case "user":
				fs.Set(flagName, flagValue)
			default:
				fmt.Println("Error: Flag not found")

				return
			}
		}

		Rmusr(*user, User_.Id)

	} else {
		fmt.Println("\n\n******************Necesita iniciar sesion como ususario ROOT para poder REMOVER un grupo***********************")
		return
	}
}
