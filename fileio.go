package main

// -----
// Conversion of Todos structs into .json format, writing and reading from a file
// -----

import (
	"encoding/json"
	"os"
)

func SaveData(data TodosMap, fileName string) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

func LoadData(data *TodosMap, fileName string) error {
	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, data)
}
