package lib

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
)

type DeepCopyableSerialized interface {
	DeepCopySerialized() DeepCopyableSerialized
}

// DeepCopySerialized performs a deep copy using encoding/gob
func DeepCopySerialized(src interface{}) interface{} {
	var buf bytes.Buffer

	// Initialize Encoder and Decoder
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	// Encoding
	err := enc.Encode(src)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return nil
	}

	// Decoding into a new instance
	srcType := reflect.TypeOf(src)
	var dst interface{}

	// Create a new instance based on whether the input type is a pointer or not
	if srcType.Kind() == reflect.Ptr {
		dst = reflect.New(srcType.Elem()).Interface()
	} else {
		dst = reflect.New(srcType).Interface()
	}

	err = dec.Decode(dst)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return nil
	}

	return dst
}
