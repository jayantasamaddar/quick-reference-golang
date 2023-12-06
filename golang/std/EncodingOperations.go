package std

import (
	"encoding/json"
	"fmt"
	"log"
)

type customer struct {
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	IsActive bool   `json:"is_active"`
}

func JSONOperationsDemo() {
	/********************************************************************************************************************/
	// JSON Encoding: Marshal
	// JavaScript equivalent of: JSON.stringify
	/********************************************************************************************************************/
	var c *customer = &customer{
		Name:     "Jayanta Samaddar",
		Age:      32,
		IsActive: true,
	}
	bytes, err := json.Marshal(&c)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		jsonstring := string(bytes)
		fmt.Printf("Value: %v, Type: %T\n", jsonstring, jsonstring)
	}

	/********************************************************************************************************************/
	// JSON Decoding: Unmarshal. Inverse of Marshal
	// JavaScript equivalent of: JSON.parse
	/********************************************************************************************************************/
	var c2 *customer
	err = json.Unmarshal(bytes, &c2)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		c2.IsActive = false
		// New c2 with changes
		fmt.Printf("Value: %v, Type: %T\n", c2, c2)

		// Original
		fmt.Printf("Value: %v, Type: %T\n", c, c)
	}
}
