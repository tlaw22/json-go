package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color: "black",
			"has_dog": true
		},
		{
			"first_name": "Brue",
			"last_name": "Wayne",
			"hair_color: "black",
			"has_dog": false
		}
	]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println(err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)
}
