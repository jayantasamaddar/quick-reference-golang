package main

import (
	"fmt"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/07-factory-pattern/lib"
)

type Person struct {
	Name, City, Country string
}

// Factory Function: A standalone function that returns an instance of the struct that you want to create
func NewPerson(name, country string) *Person {
	p := &Person{Name: name, Country: country}
	switch p.Country {
	case "India":
		p.City = "Kolkata"
	default:
		p.City = ""
	}
	return p
}

func main() {
	// (1) Factory Function with default value
	fmt.Println(NewPerson("Michael", "Australia")) // &{Michael  Australia}
	fmt.Println(NewPerson("Jayanta", "India"))     // &{Jayanta Kolkata India}

	// (2) Demonstration of Interface Factory
	p := lib.NewPerson("Tom", 30)
	p.SayHello()

	// (3a) Demonstration of Factory Generators: Functional Approach
	engineerGenerator := lib.NewEmployeeFactory("Engineer")
	e1 := engineerGenerator("Amarpreet Singh", "Senior DevOps Engineer", 6000000)
	e2 := engineerGenerator("Jayanta Samaddar", "Senior Software Engineer", 4000000)
	fmt.Println(e1, e2) // &{Amarpreet Singh Engineer Senior DevOps Engineer 6000000} &{Jayanta Samaddar Engineer Senior Software Engineer 4000000}

	// (3b) Demonstration of Factory Generators: Structural Approach
	employeeFactoryGenerator := lib.NewEmployeeFactoryStructural("Engineer")
	es1 := employeeFactoryGenerator.Create("Amarpreet Singh", "Senior DevOps Engineer", 6000000)
	es2 := employeeFactoryGenerator.Create("Jayanta Samaddar", "Senior Software Engineer", 4000000)
	fmt.Println(es1, es2) // &{Amarpreet Singh Engineer Senior DevOps Engineer 6000000} &{Jayanta Samaddar Engineer Senior Software Engineer 4000000}

	employeeFactoryGenerator.Position = "Manager"
	es3 := employeeFactoryGenerator.Create("Maulik Vinchi", "Engineering Manager", 3500000)
	fmt.Println(es3) // &{Maulik Vinchi Manager Engineering Manager 3500000}

	// (4a) Prototype Factory (Cloning a Factory)
	mr := lib.NewCandidate("Bryan Adams", lib.Male, 56, true)
	ms := mr.CloneInto(lib.Female, false).(*lib.Candidate)
	ms.Name = "Adele"
	ms.Age = 42
	mrs := ms.CloneInto(lib.Female, true).(*lib.Candidate)
	mrs.Name = "Priyanka Chopra"
	mrs.Age = 36
	fmt.Println(mr, ms, mrs) // &{Bryan Adams Mr. true 56} &{Adele Ms. false 42} &{Priyanka Chopra Mrs. true 36}
}
