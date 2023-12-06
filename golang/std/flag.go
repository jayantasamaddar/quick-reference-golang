package std

// Package flag implements command-line flag parsing.

import (
	"flag"
	"fmt"
	"strings"
)

type CustomFlag []string

// String returns a string representation of the CustomFlag type.
func (c *CustomFlag) String() string {
	return strings.Join(*c, ", ")
}

// Set parses and sets the value of the CustomFlag.
func (c *CustomFlag) Set(value string) error {
	*c = strings.Split(value, ",")
	return nil
}

func FlagOperationsDemo() {
	// Declares an integer flag, -n, stored in the pointer nFlag, with type *int:
	var nFlag = flag.Int("n", 1234, "help message for flag n")

	// Bind the flag to a variable using the Var() functions.
	var dFlag int
	flag.IntVar(&dFlag, "d", 1234, "help message for flag d")

	// Create custom flags that satisfy the Value interface (with pointer receivers) and couple them to flag parsing.
	// flag.Var is a function that enables you to define custom flag types by implementing the flag.Value interface.
	// Check the flag.Value interface here: https://pkg.go.dev/flag@go1.21.1#Value or `run go doc flag.Value`
	// This is useful when you want to parse and handle command-line arguments in a specific way that's not supported by the built-in flag types.
	// (e.g., list of comma separated items).
	// Source Code: https://cs.opensource.google/go/go/+/refs/tags/go1.21.1:src/flag/flag.go;l=1010
	var customFlag CustomFlag
	flag.Var(&customFlag, "x", "A custom flag x that takes a comma-separated list of values.")

	// After all flags are defined, call flag.Parse() to parse command line arguments
	flag.Parse()

	/*******************************************************************************************************/
	// Printing the flags (Run the following command): `go run main.go -d=1000 -n=2000 -x=apple,bat,cat`
	/*******************************************************************************************************/
	fmt.Printf("nFlag [value, type]: %d, %T\n", *nFlag, *nFlag) // nFlag [value, type]: 2000, int
	fmt.Printf("dFlag [value, type]: %d, %T\n", dFlag, dFlag)   // dFlag [value, type]: 1000, int

	// Use the custom flag value. (Run `go run main.go -x=value1,value2,value3....`)
	fmt.Printf("Custom flag value and type: %v, %T\n", customFlag, customFlag) // [apple bat cat], std.CustomFlag

	groceries := make(map[string]float32)
	groceries["Milk"] = 2.30
	groceries["Fish"] = 15.25
	groceries["Eggs"] = 1.99

	fmt.Println(groceries)
}
