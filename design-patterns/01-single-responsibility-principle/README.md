# Single Responsibility Principle

The Single Responsibility Principle states that a function should have a single primary responsibility and have only one reason to change, that reason being related to its primary responsibility.

The anti-pattern for Single Responsibility Principle is the God Object. A God Object is when you take everything inside a single package. Thus, there need to be a separation of concerns.

**Example**: Below is given an example of how we would implement a Journal where one can `AddEntry` or `RemoveEntry`. There are also two other functions `Count` and `String` to display the number of entries and stringify the entries with a separator, respectively. However, Persistence is a separate concern and we will create the Separation of Concerns, thereby implementing the Single Responsibility Principle.

In `lib/Journal.go`

```go
package lib

import (
	"fmt"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

// Add an entry to the Journal
func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// Remove an entry from the Journal
func (j *Journal) RemoveEntry(index int) bool {
	if index >= j.Count() {
		return false
	} else if index == j.Count()-1 {
		// Remove last index
		j.entries = append(j.entries[:index])
	} else {
		j.entries = append(j.entries[:index], j.entries[:index+1]...)
	}
	entryCount = j.Count()
	return true
}

// Display the count
func (j *Journal) Count() int {
	return len(j.entries)
}

// Stringify
func (j *Journal) String(sep string) string {
	return strings.Join(j.entries, sep)
}
```

In `lib/Persistence.go`

```go
package lib

import (
	"log"
	"os"
)

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	var sep string
	if p.LineSeparator != "" {
		sep = p.LineSeparator
	} else {
		sep = "\n"
	}
	err := os.WriteFile(filename, []byte(j.String(sep)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

```

In `main.go`

```go
package main

import (
	"fmt"
	"log"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/lib"
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
```

Thus we have clear separation of concerns:

1. **Journal**: Focusses on Creation, Removal, Parsing and Displaying metadata of journal entries (Journal focussed functions).
2. **Persistence**: Focusses on Saving File to Disk. Later we can even allow it to save other types than Journals, however it's function remains the same - i.e. Saving File to Disk