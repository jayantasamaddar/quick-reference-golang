package std

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

var x float64 = 3.14
var u user = user{name: "Jayanta Samaddar", age: 32}

func ReflectionOperationsDemo() {
	// (1) You can go from interface value to reflection object
	fmt.Printf("%v, Value: %v, Type: %s\n", reflect.ValueOf(x).String(), reflect.ValueOf(x), reflect.TypeOf(x).String())
	fmt.Printf("%v, Value: %v, Type: %s\n", reflect.ValueOf(u).String(), reflect.ValueOf(u), reflect.TypeOf(u).String())

	// (2) You can go from reflection object to interface value
}
