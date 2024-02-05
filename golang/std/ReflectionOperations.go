package std

import (
	"fmt"
	"reflect"
)

// Person struct with nested Address struct
type Person struct {
	Name    string
	Age     int
	Address *Address
	Hobbies []string
}

// Address struct
type Address struct {
	City  string
	State string
}

type DeepCopyable interface {
	DeepCopy() DeepCopyable
}

// Attach
func (p *Person) DeepCopy() DeepCopyable {
	return DeepCopy(p).(DeepCopyable)
}

// DeepCopy recursively copies the given struct, slices, arrays, maps and pointers
func DeepCopy(src interface{}) interface{} {
	srcValue := reflect.ValueOf(src)

	switch srcValue.Kind() {
	case reflect.Slice, reflect.Array:
		// Handle slices and arrays
		dst := reflect.MakeSlice(srcValue.Type(), srcValue.Len(), srcValue.Len())
		for i := 0; i < srcValue.Len(); i++ {
			dst.Index(i).Set(reflect.ValueOf(DeepCopy(srcValue.Index(i).Interface())))
		}
		return dst.Interface()

	case reflect.Map:
		// Handle maps
		dst := reflect.MakeMap(srcValue.Type())
		for _, key := range srcValue.MapKeys() {
			dstKey := reflect.ValueOf(DeepCopy(key.Interface()))
			dstValue := reflect.ValueOf(DeepCopy(srcValue.MapIndex(key).Interface()))
			dst.SetMapIndex(dstKey, dstValue)
		}
		return dst.Interface()

	case reflect.Ptr:
		// Handle pointers
		if srcValue.IsNil() {
			return nil
		}
		// Create a new pointer and copy the pointed value
		dst := reflect.New(srcValue.Elem().Type())
		dst.Elem().Set(reflect.ValueOf(DeepCopy(srcValue.Elem().Interface())))
		return dst.Interface()

	case reflect.Struct:
		// Handle structs
		dst := reflect.New(srcValue.Type()).Elem()
		for i := 0; i < srcValue.NumField(); i++ {
			srcField := srcValue.Field(i)
			dstField := dst.Field(i)

			if srcField.Kind() == reflect.Struct || srcField.Kind() == reflect.Slice || srcField.Kind() == reflect.Array || srcField.Kind() == reflect.Map || srcField.Kind() == reflect.Ptr {
				dstField.Set(reflect.ValueOf(DeepCopy(srcField.Interface())))
			} else {
				dstField.Set(srcField)
			}
		}
		return dst.Interface()

	default:
		// For other types, return the original value
		return src
	}
}

type user struct {
	name string
	age  int
}

func ReflectionOperationsDemo() {
	/*********************************************************************************************************/
	// (1a)  You can go from interface value to reflection object
	/*********************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: (1a) From interface value to reflection object" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	var x float64 = 3.14
	var u user = user{name: "Jayanta Samaddar", age: 32}

	fmt.Printf("%v, Value: %v, Type: %s\n", reflect.ValueOf(x).String(), reflect.ValueOf(x), reflect.TypeOf(x).String()) // <float64 Value>, Value: 3.14, Type: float64
	fmt.Printf("%v, Value: %v, Type: %s\n", reflect.ValueOf(u).String(), reflect.ValueOf(u), reflect.TypeOf(u).String()) // <user Value>, Value: {Jayanta Samaddar 32}, Type: user

	/*********************************************************************************************************/
	// (1b)  You can go from reflection object to interface value
	/*********************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: (1b) From reflection object to interface value" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	newX := reflect.ValueOf(x).Interface().(float64)
	fmt.Printf("New X: %v, Type: %s\n", newX, reflect.TypeOf(newX).String()) // New X: 3.14, Type: float64

	newUser := reflect.ValueOf(u).Interface().(user)
	fmt.Printf("New User: %v, Type: %s\n", newUser, reflect.TypeOf(newUser).String()) // New User: {Jayanta Samaddar 32}, Type: user

	/*********************************************************************************************************/
	// (2a) Reflection-Based Deep Copying (Using Pointers)
	/*********************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: (2a) Reflection-Based Deep Copying (Using Pointers)" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	// Creating a sample Person instance
	person1 := &Person{
		Name: "John Doe",
		Age:  30,
		Address: &Address{
			City:  "New York",
			State: "NY",
		},
		Hobbies: []string{"Reading", "Traveling"},
	}

	// Creating a deep copy of the Person instance (using struct method that implements the standalone function)
	person1Copy := person1.DeepCopy().(*Person)

	// Modifying the original instance
	person1Copy.Name = "Jane Doe"
	person1Copy.Age = 25
	person1Copy.Address.City = "San Francisco"
	person1Copy.Hobbies = []string{"Dancing", "Gardening"}

	// Displaying the results
	fmt.Printf("Original Type (%s): %v\n", reflect.TypeOf(person1), person1)             // Original Type (*std.Person): &{John Doe 30 0x140001340c0 [Reading Traveling]}
	fmt.Printf("Deep Copy of Type (%s): %v\n", reflect.TypeOf(person1Copy), person1Copy) // Deep Copy of Type (*std.Person): &{Jane Doe 25 0x14000134100 [Dancing Gardening]}

	/*********************************************************************************************************/
	// (2b) Reflection-Based Deep Copying (Using Value + Pointer)
	/*********************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: (2b) Reflection-Based Deep Copying (Using Value + Pointer)" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	person2 := Person{
		Name: "Michael Jackson",
		Age:  51,
		Address: &Address{
			City:  "Houston",
			State: "TX",
		},
		Hobbies: []string{"Dancing", "Singing"},
	}
	// Creating a deep copy of the Person instance (using standalone function)
	person2Copy := DeepCopy(person2).(Person)

	// Modifying the original instance
	person2Copy.Name = "Michael Jordan"
	person2Copy.Age = 55
	person2Copy.Address.City = "Chicago"
	person2Copy.Address.State = "Illinois"
	person2Copy.Hobbies = []string{"Basketball", "Sneakers"}

	// Displaying the results
	fmt.Printf("Original Type (%s): %v\n", reflect.TypeOf(person2), person2)             // Original Type (std.Person): {Michael Jackson 51 0x140001341c0 [Dancing Singing]}
	fmt.Printf("Deep Copy of Type (%s): %v\n", reflect.TypeOf(person2Copy), person2Copy) // Deep Copy of Type (std.Person): {Michael Jordan 55 0x14000134200 [Basketball Sneakers]}
}
