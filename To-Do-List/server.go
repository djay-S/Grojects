package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strconv"
)

// var toDoList []string `json:"toDoList"`
var (
	toDoList []string
	toDos    []ToDo
)

const (
	CONTENT_TYPE     string = "Content-Type"
	APPLICATION_JSON string = "application/json"
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

	http.HandleFunc("/add/id/{id}", insertToDoItemToIdx)

	http.HandleFunc("/delete/id/{id}", deleteToDoItem)

	http.HandleFunc("/delete", deleteToDoList)

	http.HandleFunc("/update/id/{id}", updateToDoItem)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getToDoList(w http.ResponseWriter, r *http.Request) {
	log.Print("Fetching To Do List")
	//	jsonToDo, _ := json.Marshal(toDoList)
	//	fmt.Fprintf(w, "%s", jsonToDo)
	//	jsonToDos, _ := json.Marshal(toDos)
	//	fmt.Fprintf(w, "%s", jsonToDos)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

	if err := json.NewEncoder(w).Encode(toDos); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
		return
	}
}

func addToDoItem(w http.ResponseWriter, r *http.Request) {
	log.Print("Adding new To Do Item")
	//	jsonDecoder := json.NewDecoder(r.Body)
	//	var toDo ToDo
	//	jsonDecoder.Decode(&toDo)
	//	toDos = append(toDos, toDo)
	//	jsonToDos, _ := json.Marshal(toDos)
	//	fmt.Fprintf(w, "%s", jsonToDos)
	var toDo ToDo
	if err := json.NewDecoder(r.Body).Decode(&toDo); err != nil {
		http.Error(w, "Invalid RequestPayload", http.StatusBadRequest)
		log.Printf("Invalid request payload: %v", err)
		return
	}

	toDos = append(toDos, toDo)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	if err := json.NewEncoder(w).Encode(toDos); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
		return
	}
}

func insertToDoItemToIdx(w http.ResponseWriter, r *http.Request) {
	//	id, _ := strconv.Atoi(r.PathValue("id"))
	log.Print("Insert a new To Do Item")
	//	jsonDecoder := json.NewDecoder(r.Body)
	//	var toDo ToDo
	//	jsonDecoder.Decode(&toDo)
	//	toDos = slices.Insert(toDos, id, toDo)
	//	jsonToDos, _ := json.Marshal(toDos)
	//	fmt.Fprintf(w, "%s", jsonToDos)
	idStr := r.URL.Path[len("/add/id/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Invalid ID: %v", err)
		return
	}
	log.Printf("Inserting new to do item at index: %d", id)
	var toDo ToDo
	if err := json.NewDecoder(r.Body).Decode(&toDo); err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		log.Printf("Invalid request payload: %v", err)
		return
	}

	if id < 0 || id > len(toDos) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		log.Printf("Invalid index: %v", err)
		return
	}

	toDos = slices.Insert(toDos, id, toDo)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

	if err := json.NewEncoder(w).Encode(toDos); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Failed to encode JSON: %v", err)
		return
	}
}

func deleteToDoItem(w http.ResponseWriter, r *http.Request) {
	//	id, _ := strconv.Atoi(r.PathValue("id"))
	log.Print("Deleting To Do item of id")
	//	toDos = append(toDos[:id], toDos[id+1:]...)
	//	fmt.Fprintf(w, "%s", toDos)
	idStr := r.URL.Path[len("/delete/id/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Invalid Id: %v", err)
		return
	}
	log.Printf("Deleting To Do item at index :%d", id)

	if id < 0 || id > len(toDos) {
		http.Error(w, "Invalid Index", http.StatusBadRequest)
		log.Printf("Invalid Index: %v", err)
		return
	}

	toDos = append(toDos[:id], toDos[id+1:]...)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

	if err := json.NewEncoder(w).Encode(toDos); err != nil {
		http.Error(w, "Failed to encode to JSON", http.StatusInternalServerError)
		log.Printf("Failed to encode to JSON: %v", err)
		return
	}
}

func deleteToDoList(w http.ResponseWriter, r *http.Request) {
	log.Print("Deleting To Do List")
	toDos = toDos[:0]
	w.WriteHeader(http.StatusNoContent)
}

func updateToDoItem(w http.ResponseWriter, r *http.Request) {
	log.Print("Updating To Do Item")
	//	id, _ := strconv.Atoi(r.PathValue("id"))
	//	jsonDecoder := json.NewDecoder(r.Body)
	//	var toDo ToDo
	//	jsonDecoder.Decode(&toDo)
	//	toDos[id] = toDo
	//	jsonToDos, _ := json.Marshal(toDos)
	//	fmt.Fprintf(w, "%s", jsonToDos)
	idStr := r.URL.Path[len("/update/id/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Invalid ID: %v", err)
		return
	}

	if id < 0 || id > len(toDos) {
		http.Error(w, "Invalid Index", http.StatusBadRequest)
		log.Printf("Invalid Index: %v", err)
		return
	}

	var toDo ToDo
	if err := json.NewDecoder(r.Body).Decode(&toDo); err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		log.Printf("Invalid Request Payload: %v", err)
		return
	}

	toDos[id] = toDo
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)

	if err := json.NewEncoder(w).Encode(toDos); err != nil {
		http.Error(w, "Failed to Encode to JSON", http.StatusInternalServerError)
		log.Printf("Failed to Encode to JSON: %v", err)
		return
	}
}

/*
1) json.NewEncoder().Encode() vs json.Marshall:
	json.NewEncoder(w).Encode(toDos) this is more efficient than json.Marshall(),
	since json.Marshal save the byte[] into memory whereas json.NewEncoder stores the json in to the HTTP responseWriter in this case
	Use json.Marshal when you need the entire JSON string in memory.
	Use json.NewEncoder when you want to stream JSON data to an io.Writer, especially in web server contexts.
2) r.PathValue() vs r.URL.Path()
	r.PathValue is designed to extract values from named path parameters defined in your routing patterns, and requires the new go 1.22 servemux. It handles the parsing for you.
	r.URL.Path provides the raw path, and you're responsible for parsing and extracting the desired information.
	r.PathValue is safer and more convenient when working with named parameters, but requires the go 1.22 servemux.
	r.URL.Path is more flexible but requires more manual work.
*/
