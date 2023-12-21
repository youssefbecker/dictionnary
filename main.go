package main

import (
	"dico_go/dictionnary"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)


func add(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	filePath := "dictionary.json"
	dict := dictionnary.NewDictionnary(filePath)
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var entry dictionnary.Entry
	err = json.Unmarshal(body, &entry)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	_, err = dict.Add(entry.Nom, entry.Prenom)
	if err != nil {
		http.Error(w, "Error adding entry", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func get(w http.ResponseWriter, req *http.Request) {
	filePath := "dictionary.json"
	dict := dictionnary.NewDictionnary(filePath)
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing 'nom' parameter", http.StatusBadRequest)
		return
	}
	nom := parts[2]
	entry, err := dict.Get(nom)
	if err != nil {
		http.Error(w, "Entry not found", http.StatusNotFound)
		return
	}
	jsonData, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func remove(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	filePath := "dictionary.json"
	dict := dictionnary.NewDictionnary(filePath)
	parts := strings.Split(req.URL.Path, "/")

	nom := parts[2]
	_, err := dict.Remove(nom)
	if err != nil {
		http.Error(w, "Entry not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("removed"))
}

func list(w http.ResponseWriter, req *http.Request) {
	filePath := "dictionary.json"
	dict := dictionnary.NewDictionnary(filePath)
	entries, err := dict.List()
	if err != nil {
		http.Error(w, "Error listing entries", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.MarshalIndent(entries, "", "  ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

type ChannelStore struct {
	action chan string
}



var cs = ChannelStore{
	action: make(chan string, 1),
}

func main() {
	

	http.HandleFunc("/add", add)
	http.HandleFunc("/get/", get)
	http.HandleFunc("/remove/", remove)
	http.HandleFunc("/list", list)

	http.ListenAndServe(":8090", nil)
}
