package main

import (
	//"BackendGo/handlers"
	Funciones "MIA_P2_201313965/Backend/Funciones"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	Id      int    `json:"id"`
	Command string `json:"comand"`
	Params  string `json:"params"`
}

type Tasks []Task

var TasksData = Tasks{
	{
		Id:      1,
		Command: "Comand",
		Params:  "params",
	},
}

type Disk []Funciones.MBR

type Mbr struct {
	Mbr_tamano int32 `json:"tamano"`

	Mbr_fecha_creacion [16]byte `json:"fecha_creacion"` // de tipo time

	Mbr_dsk_signature int32 `json:"signature"`

	Dsk_fit [1]byte `json:"fit"` // B (mejor ajuste)  F(primer ajuste) W(peor ajuste)

	Mbr_partitions Partition `json:"particiones"` // este arreglo simulara las 4 particiones

}

type Partition struct {
	Part_status bool `json:"status"` // es de tipo bool(indica si la particion esta montada o no)

	Part_type [1]byte `json:"type"` //(indica el tipo de particion: primaria(P) o extendida(E))

	Part_fit [1]byte `json:"fit"` // indica el tipo de ajuste(B mejor ajuste  F primer ajuste W peor ajuste)

	Part_start int32 `json:"start"` // indica en que byte del disco inicia la particion

	Part_size int32 `json:"size"` //(part_s) contiene el tamano total de la particion en bytes (por defecto es cero)

	Part_name [16]byte `json:"name"` // contiene el nombre de la particion

	Part_correlative int32 `json:"correlative"` // contiene el correlativo de la particion

	Part_id [4]byte `json:"id"`
}

var slice_fecha [16]byte
var slice_fit [1]byte
var slice_type [1]byte
var slice_name [16]byte
var slice_id [4]byte

type Datos []Mbr

var diskData = Datos{

	{
		Mbr_tamano:         0,
		Mbr_fecha_creacion: slice_fecha,
		Mbr_dsk_signature:  0,
		Dsk_fit:            slice_fit,
		Mbr_partitions: Partition{
			Part_status:      true,
			Part_type:        slice_type,
			Part_fit:         slice_fit,
			Part_start:       0,
			Part_size:        0,
			Part_name:        slice_name,
			Part_correlative: 0,
			Part_id:          slice_id,
		},
	},
}

/*
name := Name{FirstName: "Dikxya", Surname: "Lhyaho"}
employee := Employee{Position: "Senior Developer", Name: name}
*/
/*
var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("json map: %v\n", data)
*/

func insertComand(w http.ResponseWriter, r *http.Request) {

	var newTask Task

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}

	json.Unmarshal(body, &newTask)

	input := newTask.Command + " " + newTask.Params

	Funciones.Analyze(input) // enviando el comando

	fmt.Fprintf(w, "Comando enviado con exito")

	//newTask.ID = len(TasksData) + 1
	//TasksData = append(TasksData, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func getDisk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	letra := vars["name"]

	/*
		for _, task := range TasksData {
			if task.ID == disk {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(task)
			}
		}
	*/

	file, err := Funciones.AbrirArchivo("./archivos/" + letra + ".dsk")
	if err != nil {
		return
	}

	var tempMBR Funciones.MBR

	if err := Funciones.LeerObjeto(file, &tempMBR, 0); err != nil {
		return
	}
	// Print object
	fmt.Println("\nImprimiendo informacion del disco ", letra)
	fmt.Println()
	Funciones.PrintMBR(tempMBR)

	for i := 0; i < len(tempMBR.Mbr_partitions); i++ {
		/*fmt.Fprintf(w, "Particion %v Nombre: %s Tipo: %s fit: %s tamano: %v start: %v id: %s correlativo: %v status: %t", i,
		string(tempMBR.Mbr_partitions[i].Part_name[:]), string(tempMBR.Mbr_partitions[i].Part_type[:]), string(tempMBR.Mbr_partitions[i].Part_fit[:]), tempMBR.Mbr_partitions[i].Part_size,
		tempMBR.Mbr_partitions[i].Part_start, string(tempMBR.Mbr_partitions[i].Part_id[:]), tempMBR.Mbr_partitions[i].Part_correlative, tempMBR.Mbr_partitions[i].Part_status)*/
		fmt.Fprintf(w, "Nombre: %v ", string(tempMBR.Mbr_partitions[i].Part_name[:]))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(diskData)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	// Index Routes
	router.HandleFunc("/", welcome)
	router.HandleFunc("/insert", insertComand).Methods("POST")
	router.HandleFunc("/disk/{name}", getDisk).Methods("GET")

	// Tasks Routes
	/*
		router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
		router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
		router.HandleFunc("/tasks/{id}", handlers.GetOneTask).Methods("GET")
		router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
		router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")*/

	fmt.Println("Server started on port ", 3000)
	log.Fatal(http.ListenAndServe(":3000", router)) // para crear un servidor

	///go get -u github.com/gorilla/mux       obtiene paquete para poder crear rutas para API rest
	//go get github.com/githubnemo/CompileDaemon   para actualizar cambios automaticamente en el servidor sin necesidad de cerrarlo

}
