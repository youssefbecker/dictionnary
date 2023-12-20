package main

import (
	"dico_go/dictionnary"
	"fmt"
)

func main() {
	// Create an empty dictionary
	dictionary := dictionnary.NewDictionary()

	// Add words
	dictionary.Add("1", "mot 1")
	dictionary.Add("2", "mot 2")
	dictionary.Add("3", "mot 3")

	// Save the dictionary to a file
	filename := "dictionary.json"
	err := dictionary.SaveToFile(filename)
	if err != nil {
		fmt.Println("Error saving dictionary to file:", err)
		return
	}

	// Load the dictionary from the file
	loadedDictionary, err := dictionnary.LoadFromFile(filename)
	if err != nil {
		fmt.Println("Error loading dictionary from file:", err)
		return
	}

	// Access and manipulate the loaded dictionary
	loadedDictionary.Get("1")
	loadedDictionary.Remove("2")
	loadedDictionary.List()
}
