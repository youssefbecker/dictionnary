package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Dictionary struct {
	Entries map[string]string `json:"entries"`
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		Entries: make(map[string]string),
	}
}

func (d *Dictionary) Add(word, definition string) {
	d.Entries[word] = definition
}

func (d *Dictionary) Get(word string) {
	definition := d.Entries[word]
	fmt.Printf("%s: %s\n", word, definition)
}

func (d *Dictionary) Remove(word string) {
	delete(d.Entries, word)
}

func (d *Dictionary) List() {
	for word, definition := range d.Entries {
		fmt.Printf("%s: %s\n", word, definition)
	}
}

func (d *Dictionary) SaveToFile(filename string) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func LoadFromFile(filename string) (*Dictionary, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var dictionary Dictionary
	err = json.Unmarshal(data, &dictionary)
	if err != nil {
		return nil, err
	}

	return &dictionary, nil
}

func main() {
	// Créer un dictionnaire vide
	dictionary := NewDictionary()

	// Ajout des mots
	dictionary.Add("1", "mot 1")
	dictionary.Add("2", "mot 2")
	dictionary.Add("3", "mot 3")

	// Enregistre dans le JSON
	err := dictionary.SaveToFile("dictionary.json")

	// Charger le dictionnaire depuis le fichier JSON
	loadedDictionary, err := LoadFromFile("dictionary.json")
	if err != nil {
		fmt.Println("Erreur", err)
		os.Exit(1)
	}

	// Utiliser le dictionnaire chargé
	loadedDictionary.Get("1")
	loadedDictionary.Remove("2")
	loadedDictionary.List()
}
