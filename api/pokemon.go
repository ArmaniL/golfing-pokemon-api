package api

import (
	"encoding/json"
	"log"
	"os"
)

// Define structs to match the JSON structure.
// Field names must be capitalized to be exported and usable by the json package.
type PokemonInput struct {
	English string `json:"english"` // The tag maps the JSON key to the Go field
}

func loadFile() []PokemonInput {
	// 1. Read the entire file content
	// Use os.ReadFile for a simple, modern approach.
	jsonData, err := os.ReadFile("results.json")
	if err != nil {
		log.Fatalf("failed to read json file: %v", err)
	}

	// 2. Unmarshal the JSON data into a Go slice of structs
	var ListofPokemon []PokemonInput
	// Pass a pointer (address) to the variable where the data will be stored.
	err = json.Unmarshal(jsonData, &ListofPokemon)
	if err != nil {
		log.Fatalf("failed to unmarshal json data: %v", err)
	}
	return ListofPokemon
}
