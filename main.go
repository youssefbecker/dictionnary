package main

import (
	"fmt"
)

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// Get affiche la définition d'un mot spécifique.
func (d Dictionary) Get(word string) {
	definition := d[word]
		fmt.Printf("%s : %s\n", word, definition)
	
}

// Remove supprime un mot du dictionnaire.
func (d Dictionary) Remove(word string) {
	delete(d, word)
}

// List affiche la liste triée des mots et de leurs définitions.
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
