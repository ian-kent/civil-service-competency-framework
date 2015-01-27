package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

func main() {
	specFile := "./civil-service-competency-framework.toml"
	b, err := ioutil.ReadFile(specFile)
	if err != nil {
		log.Fatal(err)
	}

	var spec map[string]interface{}
	if _, err := toml.Decode(string(b), &spec); err != nil {
		log.Fatal(err)
	}

	encSpec, err := json.MarshalIndent(&spec, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("civil-service-competency-framework.json", encSpec, 0660)
	if err != nil {
		log.Fatal(err)
	}
}
