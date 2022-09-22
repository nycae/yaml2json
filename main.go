package main

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	var data any
	if err := yaml.NewDecoder(os.Stdin).Decode(&data); err != nil {
		log.Fatalf("unable to unmarshal: %v", err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(&data); err != nil {
		log.Fatalf("unable to marshal: %v", err)
	}
}
