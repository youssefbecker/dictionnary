package main

import (
	"fmt"
)

type Dictionary map[string]string

// Ajout 
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// Get
func (d Dictionary) Get(word string) {
	definition := d[word]
		fmt.Printf("%s : %s\n", word, definition)
	
}

// Remove
func (d Dictionary) Remove(word string) {
	delete(d, word)
}

// List
func (d Dictionary) List() {
	var words []string
	for word := range d {
		words = append(words, word)
	}

	for _, word := range words {
		fmt.Printf("%s : %s\n", word, d[word])
	}
}

func main() {
	//dictionnaire vide
	dictionary := make(Dictionary)

	// Ajout des mots
	dictionary.Add("1", "mot 1")
	dictionary.Add("2", "mot 2")
	dictionary.Add("3", "mot 3")

	//dictionary.Get("1")
	dictionary.Remove("2")
	dictionary.List()
}
