package dictionnary

import (
	"encoding/json"
	"fmt"
	"os"
)


type Entry struct {
	Nom  string `json:"nom"`
	Prenom string `json:"prenom"`
}

type Dictionnary struct {
	filePath string
}


func (dict *Dictionnary) List() ([]Entry, error){
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return nil, nil
	}

	return entries, nil
}

func (dict *Dictionnary) saveToFile(entries []Entry) error {
	jsonData, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(dict.filePath, jsonData, 644)
	if err != nil {
		return err
	}

	return nil	
}

func (dict *Dictionnary) loadFromFile() ([]Entry, error) {
	jsonData, err := os.ReadFile(dict.filePath)
	if err != nil {
		return nil, err
	}

	// Check if the JSON data is empty
	if len(jsonData) == 0 {
		return nil, nil
	}

	var entries []Entry
	err = json.Unmarshal(jsonData, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}


func (dict *Dictionnary) saving(entries []Entry) {
	if err := dict.saveToFile(entries); err != nil {
		fmt.Println("Error saving to file:", err)
	}
}

func NewDictionnary(filePath string) Dictionnary {
	return Dictionnary{
		filePath: filePath,
	}
}

func (dict *Dictionnary) Add(nom string, prenom string, action chan string) (Entry, error) {
	action <- "adding"
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return Entry{}, err
	}

	for i, entry := range entries {
		if entry.Nom == nom {
		
			fmt.Printf("Updating existing entry '%s': %s to %s\n", nom, entry.Prenom, prenom)
			entries[i].Prenom = prenom
			dict.saving(entries)
			return entry, nil
		}
	}

	entry := Entry{Nom: nom, Prenom: prenom}
	entries = append(entries, entry)
	dict.saving(entries)
	return entry, nil

}

func (dict *Dictionnary) Get(nom string) (Entry, error) {
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return Entry{}, err
	}

	for _, entry := range entries {
		fmt.Println("Found:", entry)
		if entry.Nom == nom {
			return entry, nil
		}
	}

	fmt.Println("Key not found:", nom)
	return Entry{}, nil
}

func (dict *Dictionnary) Remove(nom string) (Entry, error) {
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return Entry{}, err 
	}

	for i, entry := range entries {
		if entry.Nom == nom {
			fmt.Println("Removing:", entry)
			entries = append(entries[:i], entries[i+1:]...)
			dict.saving(entries)
			return entry, nil
		}
	}


	fmt.Println("Key not found:", nom)
	return Entry{}, nil
}