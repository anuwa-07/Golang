package main

import (
	"log"
	"net/http"
	"time"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)


// make JSON obj to send data
type Note struct {
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedOn time.Time `json:createdon`
}

// store for the Notes
var noteStore = make(map[string]Note)

// Variable to generate key for the collection
var id int = 0

// HTTP Post - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note

	// decode the incomming json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatal(err)
	}

	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note

	jsn, err := json.Marshal(note)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsn)
}


// HTTP Get - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}


// HTTP Put - /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	k := vars["id"]

	var noteToUpdate Note

	// decode the incomming Note json
	err = json.NewDecoder(r.Body).Decode(&noteToUpdate)
	if err != nil {
		log.Fatal(err)
	}

	if note, ok := noteStore[k]; ok {
		noteToUpdate.CreatedOn = note.CreatedOn

		// delecte existing item and add the updated item
		delete(noteStore, k)
		noteStore[k] = noteToUpdate
	} else {
		log.Printf("Unable to find key of Notes to delete")
	}
	w.WriteHeader(http.StatusNoContent)
}


//HTTP Delete - /api/notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	// Remove from Store
	if _, ok := noteStore[k]; ok {
	//delete existing item
	delete(noteStore, k)
	} else {
	log.Printf("Could not find key of Note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}


//Entry point of the program
func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	
	server := &http.Server{
		Addr:
		":8080",
		Handler: r,
	}
	// 
	log.Println("Listening...")
	server.ListenAndServe()
}





