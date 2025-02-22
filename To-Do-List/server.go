package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// var toDoList []string `json:"toDoList"`
var (
	toDoList []string
	toDos    []ToDo
)

type ToDo struct {
	Priority    int    `json:"priority"`
	Description string `json:"description"`
}

func main() {
	log.Print("Server started...")
	toDoList = append(toDoList, "Start", "Postman")
	http.HandleFunc("/get", getToDoList)

	http.HandleFunc("/add", addToDoItem)

	http.HandleFunc("/delete/id/{id}", deleteToDoItem)

	http.HandleFunc("/delete", deleteToDoList)

	http.HandleFunc("/update/id/{id}", updateToDoItem)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getToDoList(w http.ResponseWriter, r *http.Request) {
	log.Print("Fetching To Do List")
	//	jsonToDo, _ := json.Marshal(toDoList)
	//	fmt.Fprintf(w, "%s", jsonToDo)
	jsonToDos, _ := json.Marshal(toDos)
	fmt.Fprintf(w, "%s", jsonToDos)
}

func addToDoItem(w http.ResponseWriter, r *http.Request) {
	log.Print("Adding new To Do Item")
	jsonDecoder := json.NewDecoder(r.Body)
	var toDo ToDo
	jsonDecoder.Decode(&toDo)
	toDos = append(toDos, toDo)
	jsonToDos, _ := json.Marshal(toDos)
	fmt.Fprintf(w, "%s", jsonToDos)
}

func deleteToDoItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	log.Print("Deleting To Do item of id")
	toDos = append(toDos[:id], toDos[id+1:]...)
	fmt.Fprintf(w, "%s", toDos)
}

func deleteToDoList(w http.ResponseWriter, r *http.Request) {
	log.Print("Deleting To Do List")
	toDos = toDos[:0]
}

func updateToDoItem(w http.ResponseWriter, r *http.Request) {
	log.Print("Updating To Do Item")
	id, _ := strconv.Atoi(r.PathValue("id"))
	jsonDecoder := json.NewDecoder(r.Body)
	var toDo ToDo
	jsonDecoder.Decode(&toDo)
	toDos[id] = toDo
	//	slices.Insert(toDos, id, toDo)
	jsonToDos, _ := json.Marshal(toDos)
	fmt.Fprintf(w, "%s", jsonToDos)
}
