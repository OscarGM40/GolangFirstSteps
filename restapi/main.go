package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

/* FAKE DATA   */
type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []task // slice o array dinámico de tareas

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Tarea 1",
		Content: "Ejercicio 1",
	},
}


/*----------- CONTROLADORES  -------------  */
/* necesito del writter para escribir en un header o un body*/
func getTasks(w http.ResponseWriter, r *http.Request) {
	/* escribo una cabecera con este header */
	w.Header().Set("Content-Type", "application/json")
	/* mando un 200 */
	w.WriteHeader(http.StatusOK)
	/* y ahora escribo el json en el body de la response */
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	// defino una variable de tipo task
	var newTask task;

	// en vez de newDecoder para leer usaré ioutil.ReadAll
	// _ = json.NewDecoder(r.Body).Decode(&newTask)
	reqBody,err := ioutil.ReadAll(r.Body);
	
	if (err != nil) {
		fmt.Fprintf(w, "Error al leer el body %v\n", err);
	}

	/* unmarshal parsea el json de la request y lo guarda en un puntero */
	json.Unmarshal(reqBody, &newTask); //puntero a newTask

	// le sumo uno al ID de la tarea
	newTask.ID = len(tasks) + 1;
	// y concateno la tarea
	tasks = append(tasks, newTask)

	/* fijate que NewEncoder es para escribir a la response y NewDecoder para leer de la response.Además si escribo tengo que codificar con Encode y si leo decodificar con Decode*/
	w.Header().Set("Content-Type", "application/json")
	/* parece que hay que escribir cabeceras de una en una */
	/* fijate que uno es Header y el otro WriteHeader,asinto */
	w.WriteHeader(http.StatusCreated);
	json.NewEncoder(w).Encode(newTask)
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la API en Go!\n")
}

func getTaskById(w http.ResponseWriter, r *http.Request) {
	/* mux.Vars(r) extrae las variables de la request */
	vars := mux.Vars(r)
	/* pero ojo,porque los devuelve como strings siempre */
	/* fijate que Atoi puede dar un error */
	taskID,err := strconv.Atoi(vars["id"]);
	
	if (err != nil) {
		fmt.Fprintf(w, "Error al convertir el id %v\n", err);
		return;
	}
	// fmt.Fprintf(w, "Task ID: %v\n", taskID)
	/* busco la tarea en la lista */
	for _,task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}else{
			fmt.Fprintf(w, "No se encontro la tarea con el id %v\n", taskID)
			return
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID,err := strconv.Atoi(vars["id"]);
	
	if (err != nil) {
		fmt.Fprintf(w, "Error al convertir el id %v\n", err);
		return;
	}

	/* si el id es válido busco la tarea y la borro */
	for i,t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i + 1:]...);
			fmt.Fprintf(w, "Se elimino la tarea con el id %v\n", taskID);
			return;
		}else{
			fmt.Fprintf(w, "No se encontro la tarea con el id %v\n", taskID);
			return;
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) //extrae las variables de la request
	taskID,err := strconv.Atoi(vars["id"]);
	
	if (err != nil) {
		fmt.Fprintf(w, "Error al convertir el id %v\n", err);
		return;
	}
	// leo el body de la request controlando el error
	reqBody,err := ioutil.ReadAll(r.Body);
	
	if (err != nil) {
		fmt.Fprintf(w, "Error al leer el body %v\n", err);
	}
	
	var updatedTask task;
	json.Unmarshal(reqBody, &updatedTask); //puntero a updatedTask

	for i,t := range tasks {
		if t.ID == taskID {
			/* primero la elimino y despues la guardo */
			tasks = append(tasks[:i], tasks[i + 1:]...);
			/* al eliminarla necesito asignarla el mismo ID */
			updatedTask.ID = taskID;
			tasks = append(tasks, updatedTask);
			fmt.Fprintf(w, "Se actualizo la tarea con el id %v\n", taskID);
			return;
		}else{
			fmt.Fprintf(w, "No se encontro la tarea con el id %v\n", taskID);
			return;
		}
	}
}

func main() {
	/* strictSlash a true no permite usar una trailing '/' */
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeRoute).Methods("GET")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}",getTaskById).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	
	
	log.Fatal(http.ListenAndServe(":3000", router))
}
