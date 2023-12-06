package main

import (
	"fmt"
	"log"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/01-single-responsibility-principle/lib"
)

func main() {
	j := lib.Journal{}

	// Add some Journal entries
	j.AddEntry("I became more determined today!")
	j.AddEntry("Visited a friend")
	j.AddEntry("Finished office work")
	fmt.Println("Added Journal entries")
	fmt.Println(j.String("\n"))
	fmt.Println("Total entries:", j.Count())

	// Remove the last Journal entry
	result := j.RemoveEntry(j.Count() - 1)
	if result {
		fmt.Println("Removed Journal entries")
		fmt.Println(j.String("\n"))
		fmt.Println("Total entries:", j.Count())
	} else {
		log.Println("Could not remove entry")
		fmt.Println(j.String("\n"))
	}

	// Persist to file
	var persist *lib.Persistence = new(lib.Persistence)
	persist.LineSeparator = "\n"

	persist.SaveToFile(&j, "./journal.txt")
	log.Println("Persisted to file")
}
