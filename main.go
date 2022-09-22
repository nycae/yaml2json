package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var indent = flag.String("t", "", "indent character")

func main() {
	flag.Parse()
	var data any
	if err := yaml.NewDecoder(os.Stdin).Decode(&data); err != nil {
		log.Fatalf("unable to unmarshal: %v", err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", *indent)
	if err := enc.Encode(&data); err != nil {
		log.Fatalf("unable to marshal: %v", err)
	}
}
