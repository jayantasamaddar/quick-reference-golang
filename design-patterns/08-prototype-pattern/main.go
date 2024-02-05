package main

import (
	"fmt"
	"reflect"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/08-prototype-pattern/lib"
	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
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

// Attach
func (p *Person) DeepCopy() lib.DeepCopyable {
	return lib.DeepCopy(p).(lib.DeepCopyable)
}

func (p *Person) DeepCopySerialized() lib.DeepCopyableSerialized {
	return lib.DeepCopySerialized(p).(lib.DeepCopyableSerialized)
}

/*********************************************************************************************************/
// Main Function
/*********************************************************************************************************/
func main() {
	/*********************************************************************************************************/
	// (1a) Reflection-Based Deep Copying (Using Pointers)
	/*********************************************************************************************************/
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
	fmt.Println(std.Purple + "\tDemo: Reflection-Based Deep Copying (Using Pointers)" + std.Reset)
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
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
	fmt.Printf("Original Type (%s): %v\n", reflect.TypeOf(person1), person1)             // Original Type (*main.Person): &{John Doe 30 0x140001420e0 [Reading Traveling]}
	fmt.Printf("Deep Copy of Type (%s): %v\n", reflect.TypeOf(person1Copy), person1Copy) // Deep Copy of Type (*main.Person): &{Jane Doe 25 0x14000142120 [Dancing Gardening]}

	/*********************************************************************************************************/
	// (1b) Reflection-Based Deep Copying (Using Value + Pointer)
	/*********************************************************************************************************/
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
	fmt.Println(std.Purple + "\tDemo: Reflection-Based Deep Copying (Using Value + Pointer)" + std.Reset)
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
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
	person2Copy := lib.DeepCopy(person2).(Person)

	// Modifying the original instance
	person2Copy.Name = "Michael Jordan"
	person2Copy.Age = 55
	person2Copy.Address.City = "Chicago"
	person2Copy.Address.State = "Illinois"
	person2Copy.Hobbies = []string{"Basketball", "Sneakers"}

	// Displaying the results
	fmt.Printf("Original Type (%s): %v\n", reflect.TypeOf(person2), person2)             // Original Type (main.Person): {Michael Jackson 51 0x140001421e0 [Dancing Singing]}
	fmt.Printf("Deep Copy of Type (%s): %v\n", reflect.TypeOf(person2Copy), person2Copy) // Deep Copy of Type (main.Person): {Michael Jordan 55 0x14000142220 [Basketball Sneakers]}

	/*********************************************************************************************************/
	// (2a) Serialized-Based Deep Copying (Using Pointers)
	/*********************************************************************************************************/
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
	fmt.Println(std.Purple + "\tDemo: Serialization-Based Deep Copying (Using Pointers)" + std.Reset)
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
	person1SerializedCopy := person1.DeepCopySerialized().(*Person)
	person1SerializedCopy.Name = "Nelson Mandela"
	person1SerializedCopy.Age = 100
	person1SerializedCopy.Address.City = "Cape Town"
	person1SerializedCopy.Address.State = "South Africa"
	person1SerializedCopy.Hobbies = []string{"Activism", "Writing"}
	// Displaying the results
	fmt.Printf("Original Type (%s): %v\n", reflect.TypeOf(person1), person1)                                 // Original Type (*main.Person): &{John Doe 30 0x140001400e0 [Reading Traveling]}
	fmt.Printf("Deep Copy of Type (%s): %v\n", reflect.TypeOf(person1SerializedCopy), person1SerializedCopy) // Deep Copy of Type (*main.Person): &{Nelson Mandela 100 0x14000140500 [Activism Writing]}

	/*********************************************************************************************************************/
	// (2b) Serialized-Based Deep Copying (Using Value + Pointer)
	// NOTE: Always returns a pointer which has to be dereferenced
	/*********************************************************************************************************************/
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
	fmt.Println(std.Purple + "\tDemo: Serialization-Based Deep Copying (Using Value + Pointer)" + std.Reset)
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
	person2SerializedCopy := *lib.DeepCopySerialized(person2).(*Person) // Pointer dereferenced to ensure the copied output is the same type
	person2SerializedCopy.Name = "Nelson Mandela"
	person2SerializedCopy.Age = 100
	person2SerializedCopy.Address.City = "Cape Town"
	person2SerializedCopy.Address.State = "South Africa"
	person2SerializedCopy.Hobbies = []string{"Activism", "Writing"}
	// Displaying the results
	fmt.Printf("Original Type (%s): %v\n", reflect.TypeOf(person2), person2)                                 // Original Type (main.Person): {Michael Jackson 51 0x14000090200 [Dancing Singing]}
	fmt.Printf("Deep Copy of Type (%s): %v\n", reflect.TypeOf(person2SerializedCopy), person2SerializedCopy) // Deep Copy of Type (main.Person): {Nelson Mandela 100 0x14000090740 [Activism Writing]}
}
