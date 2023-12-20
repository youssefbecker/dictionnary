package dictionnary

import (
	"encoding/json"
	"fmt"
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

	return os.WriteFile(filename, data, 0644)
}

func LoadFromFile(filename string) (*Dictionary, error) {
	data, err := os.ReadFile(filename)
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
