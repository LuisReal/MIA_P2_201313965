package main

import (
	//"BackendGo/handlers"
	Funciones "MIA_P2_201313965/Backend/Funciones"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Task struct {
	Command string `json:"comand"`
}

type dataConsola struct {
	Data string `json:"data"`
}

type ArrayConsola []dataConsola

var TasksConsola = ArrayConsola{
	{
		Data: "",
	},
}

type Login struct {
	IdPartition string `json:"id"`
	Usuario     string `json:"user"`
}

type Sesion []Login

type Tasks []Task

var TasksData = Tasks{
	{
		Command: "Comand",
	},
}

type Disk struct {
	Nombre string `json:"nombre"`
}

type Discos []Disk

/*
var DisksData = Discos{
	{
		Nombre: "",
	},
}*/

type Mbr struct {
	Mbr_tamano int32 `json:"tamano"`

	Mbr_fecha_creacion string `json:"fecha_creacion"` // de tipo time

	Mbr_dsk_signature int32 `json:"signature"`

	Dsk_fit string `json:"fit"` // B (mejor ajuste)  F(primer ajuste) W(peor ajuste)

	Mbr_partitions [4]Partition `json:"particiones"` // este arreglo simulara las 4 particiones

}

type Partition struct {
	Part_status bool `json:"status"` // es de tipo bool(indica si la particion esta montada o no)

	Part_type string `json:"type"` //(indica el tipo de particion: primaria(P) o extendida(E))

	Part_fit string `json:"fit"` // indica el tipo de ajuste(B mejor ajuste  F primer ajuste W peor ajuste)

	Part_start int32 `json:"start"` // indica en que byte del disco inicia la particion

	Part_size int32 `json:"size"` //(part_s) contiene el tamano total de la particion en bytes (por defecto es cero)

	Part_name string `json:"name"` // contiene el nombre de la particion

	Part_correlative int32 `json:"correlative"` // contiene el correlativo de la particion

	//Part_identificador string `json:"identificador"`

	Part_id string `json:"id"`
}

func insertComand(w http.ResponseWriter, r *http.Request) {

	var newTask Task
	//var login Login
	//var sesionArray Sesion
	var consola dataConsola

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}

	json.Unmarshal(body, &newTask)

	input := newTask.Command

	//fmt.Fprintf(w, "imprimiendo newTask\n%v", input)

	data := Funciones.Analyze(input) // enviando el comando

	consola.Data = data

	//fmt.Fprintf(w, "\nimprimiendo data consola\n%v", consola)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	/*
		if Funciones.User_.Status {
			//fmt.Fprintf(w, "Devolviendo la informacion del login\n")
			login.Usuario = Funciones.User_.Nombre
			login.IdPartition = Funciones.User_.Id

			sesionArray = append(sesionArray, login)

			json.NewEncoder(w).Encode(sesionArray)
		} else {
			//json.NewEncoder(w).Encode(newTask)
			json.NewEncoder(w).Encode(consola)
		}*/

	json.NewEncoder(w).Encode(consola)

}

func getDisk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	letra := vars["name"]

	file, err := Funciones.AbrirArchivo("./archivos/" + letra + ".dsk")
	if err != nil {
		return
	}

	defer file.Close()

	var tempMBR Funciones.MBR

	if err := Funciones.LeerObjeto(file, &tempMBR, 0); err != nil {
		return
	}

	var indice_fecha int

	var newMbr Mbr

	for k := 0; k < len(tempMBR.Mbr_fecha_creacion[:]); k++ {
		if tempMBR.Mbr_fecha_creacion[k] == 0 { //quitando espacios(los ceros restantes) al slice de B_name
			indice_fecha = k
			break
		}

	}
	//fmt.Println("\nEl indice es: ", indice)
	fecha := string(tempMBR.Mbr_fecha_creacion[:indice_fecha])

	var indice_fit int

	for k := 0; k < len(tempMBR.Dsk_fit[:]); k++ {
		if tempMBR.Dsk_fit[k] == 0 { //quitando espacios(los ceros restantes) al slice de B_name
			indice_fit = k
			break
		}

	}
	//fmt.Println("\nEl indice es: ", indice)
	fit := string(tempMBR.Dsk_fit[:indice_fit])

	newMbr.Dsk_fit = fit
	newMbr.Mbr_dsk_signature = tempMBR.Mbr_dsk_signature
	newMbr.Mbr_fecha_creacion = fecha

	for j := 0; j < 4; j++ {

		var indice_fit int

		for k := 0; k < len(tempMBR.Mbr_partitions[j].Part_fit[:]); k++ {
			if tempMBR.Mbr_partitions[j].Part_fit[k] == 0 { //quitando espacios(los ceros restantes) al slice de B_name
				indice_fit = k
				break
			}

		}

		var fit string

		if indice_fit != 0 {
			fit = string(tempMBR.Mbr_partitions[j].Part_fit[:indice_fit])
		}

		var indice_tipo int

		for k := 0; k < len(tempMBR.Mbr_partitions[j].Part_type[:]); k++ {
			if tempMBR.Mbr_partitions[j].Part_type[k] == 0 { //quitando espacios(los ceros restantes) al slice de B_name
				indice_tipo = k
				break
			}

		}

		var tipo string

		if indice_tipo != 0 {
			tipo = string(tempMBR.Mbr_partitions[j].Part_type[:indice_tipo])
		}

		var indice_name int

		for k := 0; k < len(tempMBR.Mbr_partitions[j].Part_name[:]); k++ {
			if tempMBR.Mbr_partitions[j].Part_name[k] == 0 { //quitando espacios(los ceros restantes) al slice de B_name
				indice_name = k
				break
			}

		}

		var name string

		if indice_name != 0 {
			name = string(tempMBR.Mbr_partitions[j].Part_name[:indice_name])
		}

		var indice_id int

		for k := 0; k < len(tempMBR.Mbr_partitions[j].Part_id[:]); k++ {
			if tempMBR.Mbr_partitions[j].Part_id[k] == 0 { //quitando espacios(los ceros restantes) al slice de B_name

				indice_id = k

				break
			}

		}

		var id string

		var bytes_id [4]byte

		if indice_id != 0 {

			if tempMBR.Mbr_partitions[j].Part_id != bytes_id { // si el slice Part_id no esta vacio

				id = string(tempMBR.Mbr_partitions[j].Part_id[:indice_id])
				newMbr.Mbr_partitions[j].Part_id = id
			}

		} else {

			if tempMBR.Mbr_partitions[j].Part_id != bytes_id { // si el slice Part_id no esta vacio

				id = string(tempMBR.Mbr_partitions[j].Part_id[:])
				newMbr.Mbr_partitions[j].Part_id = id
			}

		}

		newMbr.Mbr_partitions[j].Part_correlative = tempMBR.Mbr_partitions[j].Part_correlative
		newMbr.Mbr_partitions[j].Part_fit = fit

		newMbr.Mbr_partitions[j].Part_name = name
		newMbr.Mbr_partitions[j].Part_size = tempMBR.Mbr_partitions[j].Part_size
		newMbr.Mbr_partitions[j].Part_start = tempMBR.Mbr_partitions[j].Part_start
		newMbr.Mbr_partitions[j].Part_status = tempMBR.Mbr_partitions[j].Part_status
		newMbr.Mbr_partitions[j].Part_type = tipo
	}

	mbr_json, err := json.MarshalIndent(newMbr, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(mbr_json))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMbr)

}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido")
}

func getAllDisks(w http.ResponseWriter, r *http.Request) {

	var newDisk Disk

	var discos Discos

	lista_discos, err := os.ReadDir("./archivos")

	if err != nil {
		fmt.Println("Hubo un error al leer los discos")
	}

	for _, f := range lista_discos {

		newDisk.Nombre = f.Name()

		discos = append(discos, newDisk)
	}

	//json.Unmarshal(reqBody, &Disco)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(discos)
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	// Index Routes
	router.HandleFunc("/", welcome)
	router.HandleFunc("/insert", insertComand).Methods("POST")
	router.HandleFunc("/disk/{name}", getDisk).Methods("GET")
	router.HandleFunc("/discos", getAllDisks).Methods("GET")

	// Tasks Routes
	/*
		router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
		router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
		router.HandleFunc("/tasks/{id}", handlers.GetOneTask).Methods("GET")
		router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
		router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")*/

	fmt.Println("Server started on port ", 3000)
	http.ListenAndServe(":3000",

		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		)(router)) // para crear un servidor

	///go get -u github.com/gorilla/mux       obtiene paquete para poder crear rutas para API rest
	//go get github.com/githubnemo/CompileDaemon   para actualizar cambios automaticamente en el servidor sin necesidad de cerrarlo

}

/*
for i := 0; i < len(tempMBR.Mbr_partitions); i++ {
		/*fmt.Fprintf(w, "Particion %v Nombre: %s Tipo: %s fit: %s tamano: %v start: %v id: %s correlativo: %v status: %t", i,
		string(tempMBR.Mbr_partitions[i].Part_name[:]), string(tempMBR.Mbr_partitions[i].Part_type[:]), string(tempMBR.Mbr_partitions[i].Part_fit[:]), tempMBR.Mbr_partitions[i].Part_size,
		tempMBR.Mbr_partitions[i].Part_start, string(tempMBR.Mbr_partitions[i].Part_id[:]), tempMBR.Mbr_partitions[i].Part_correlative, tempMBR.Mbr_partitions[i].Part_status)
		fmt.Fprintf(w, "Nombre: %v ", string(tempMBR.Mbr_partitions[i].Part_name[:]))
	}
*/
