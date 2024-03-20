package std

import (
	b "bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type customer struct {
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	IsActive bool   `json:"is_active"`
}

func JSONOperationsDemo() {
	/********************************************************************************************************************/
	// (1a) JSON Encoding: Marshal
	// JavaScript equivalent of: JSON.stringify
	/********************************************************************************************************************/
	PrintHeader("JSON Encoding: Marshal")
	var c customer = customer{
		Name:     "Jayanta Samaddar",
		Age:      32,
		IsActive: true,
	}
	bytes, err := json.Marshal(&c)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		jsonstring := string(bytes)
		fmt.Printf("Marshalled Value: %v, Type: %T\n", jsonstring, jsonstring)
	}

	/********************************************************************************************************************/
	// (1b) JSON Decoding: Unmarshal. Inverse of Marshal
	// JavaScript equivalent of: JSON.parse
	/********************************************************************************************************************/
	PrintHeader("JSON Decoding: Unmarshal")
	var c2 customer
	err = json.Unmarshal(bytes, &c2)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		c2.IsActive = false
		// New c2 with changes
		fmt.Printf("Unmarshalled Value: %v, Type: %T\n", c2, c2)

		// Original
		fmt.Printf("Original: %v, Type: %T\n", c, c)
	}

	/********************************************************************************************************************/
	// (2a) Gob (Go Binary) Encoding
	/********************************************************************************************************************/
	PrintHeader("Gob (Go Binary) Encoding")
	var buffer b.Buffer
	encoder := gob.NewEncoder(&buffer)
	err = encoder.Encode(&c)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
	fmt.Printf("Encoded Data (gob): %v, Type: %T\n", buffer, buffer)

	/********************************************************************************************************************/
	// (2b) Gob (Go Binary) Decoding
	/********************************************************************************************************************/
	PrintHeader("Gob (Go Binary) Decoding")
	// Create a decoder and decode data from the buffer
	var decodedData customer
	decoder := gob.NewDecoder(&buffer)
	err = decoder.Decode(&decodedData)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}
	// Print decoded data
	fmt.Printf("Decoded Data (gob): %v, Type: %T\n", decodedData, decodedData)

	/********************************************************************************************************************/
	// (3a) Binary Encoding
	/********************************************************************************************************************/
	PrintHeader("Binary Encoding")
	// resetting buffer
	buffer = b.Buffer{}
	err = binary.Write(&buffer, binary.LittleEndian, &c)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		if strings.Contains(err.Error(), "some values are not fixed-sized") {
			fmt.Printf("binary.Write requires fixed-size values, and the %T struct contains a field that is variable-sized.\n", &c)
			fmt.Println("Variable types are: Strings, Slice or Array, Map, Structs with Variable-Length fields, Interfaces, Pointers")
		}
		return
	}
	fmt.Printf("Encoded Data (binary): %v, Type: %T\n", buffer, buffer)

	/********************************************************************************************************************/
	// (3b) Binary Decoding
	/********************************************************************************************************************/
	// PrintHeader("Binary Decoding")
	// // Create a buffer reader
	// bufferReader := b.NewReader(buffer.Bytes())
	// var c3 customer
	// err = binary.Read(&buffer, binary.LittleEndian, &c3)
	// if err != nil {
	// 	fmt.Println("Error decoding data:", err)
	// 	return
	// }
	// // Print decoded data
	// fmt.Printf("Decoded Data (binary): %v, Type: %T\n", c3, c3)
}
