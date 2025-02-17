package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
}

func deleteToDoItem(w http.ResponseWriter, r *http.Request) {
	log.Print("Deleting To Do item")
	jsonDecoder := *json.NewDecoder(r.Body)
	var deleteToDo ToDo
	jsonDecoder.Decode(deleteToDo)
}

func deleteToDoList(w http.ResponseWriter, r *http.Request) {
	log.Print("Deleting To Do List")
}

func updateToDoItem(w http.ResponseWriter, r *http.Request) {
	log.Print("Updating To Do Item")
}
