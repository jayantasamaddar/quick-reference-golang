# Table of Contents

- [Table of Contents](#table-of-contents)
- [Singleton Pattern](#singleton-pattern)
  - [Introduction](#introduction)
  - [Construction](#construction)
  - [Problems with Singleton](#problems-with-singleton)
  - [Singleton and Dependency Inversion](#singleton-and-dependency-inversion)
- [Summary](#summary)

---

# Singleton Pattern

## Introduction

For some components it only makes sense to have one instance of it in the system.

Examples:

- Database repository
- Object factory

The need arises when the construction call is expensive. Thus we want to do it only once. We give everyone the same instance

With this setup, we want to prevent anyone from creating additional copies.

Another thing we need to take care of is: Lazy Instantiation.

---

## Construction

Imagine we have a database of capitals with their population that we want to load onto memory. Obviously, if you are loading onto memory, you only want to do it once. While doing so, we must also worry about thread safety because we do not want two threads to start initializing the database object at the same time.

We have two options to handle thread safety:

1. `sync.Once`
2. Package level `init()` function

However, we also want lazy initialization, i.e. we only read from memory only when there's a query for it. Laziness, however is not going to be guaranteed in the `init()` function, however it can be guaranteed in the `sync.Once`.

In **`lib/SingletonDatabase.go`**:

```go
package lib

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

// We want the database to have only one instance
var once sync.Once
var instance *singletonDatabase

type singletonDatabase struct {
	// Map storing capital against its population
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// Reads a file from a path and loads up a map from string to int
func readData(path string) (map[string]int, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fullPath := filepath.Join(filepath.Dir(ex), path)

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

func GetSingletonDatabase() *singletonDatabase {
	// Do calls the function if and only if Do is being called for the first time for this instance of Once.
	once.Do(func() {
		caps, err := readData("./capitals.txt")
		db := singletonDatabase{caps}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}
```

In **`main.go`**:

```go
package main

import (
	"fmt"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/09-singleton-pattern/lib"
)

func main() {
	db := lib.GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")

	fmt.Println("Population of Seoul:", pop) // Population of Seoul: 17500000
}
```

Run: We are using `os.Executable()` to get the path to the executable, which is the same directory as the `capitals.txt`, hence we need to build first.

```bash
go build
go run ./09-singleton-pattern
```

---

## Problems with Singleton

Singleton is often not the best idea as it quite often breaks the Dependency Inversion Principle.

Suppose you want to get the total population of cities (for our previous example), provided by a string slice.

In **`lib/SingletonDatabase.go`**:

```go
// Problems with Singleton
func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}
```

In **`main.go`**:

```go
package main

import (
	"fmt"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/09-singleton-pattern/lib"
)

func main() {
	// Problems with Singleton
	cities := []string{"Seoul", "Delhi", "New York"}
	tp := lib.GetTotalPopulation(cities)

    // Rudimentary Unit test to test whether this is working
	ok := tp == 17500000+14300000+17800000
	fmt.Println(tp, ok) // 49600000 true
}
```

While the test passes, the problem is that the test is dependent upon data from a real life database. In real life software engineering, you almost never test against the live database because the database can change at any time. This muddles our unit test because the data is dynamic and can change anytime. Also, we are actually doing an integration test because the database connection also has to be tested and so on.

One of the ideas of the Dependency Inversion Principle is that instead of depending upon concrete implementations, you want to depend upon abstractions, which typically implies depending upon an interface.

---

## Singleton and Dependency Inversion

What we can do is, we can slightly change the way that we calculate the `GetTotalPopulation` so that we now introduce a dependency.

Previously, the way we had implemented `GetTotalPopulation` is that it is hardcoded to use the Singleton Pattern and there was no way for us to substitute something instead of the Live Database. We need to use a `DummyDatabase` that can work predictably with a modified `GetTotalPopulationEx` function that initializes the `DummyDatabase` if it doesn't exist, and works on that.

We can write a better function for this:

In **`lib/SingletonDatabase.go`**

```go
package lib

// think of a module as a singleton
type Database interface {
	GetPopulation(name string) int
}

/************************************************************************/
// Singleton with Dependency Inversion
/************************************************************************/
type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3}
	}
	return d.dummyData[name]
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}
```

In **`main.go`**:

```go
package main

import (
	"fmt"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/09-singleton-pattern/lib"
)

func main() {
	/**********************************************************************************************************/
	// Problems with Singleton (Unit tests are a problem as they work on the Live Database)
	/**********************************************************************************************************/
	cities := []string{"Seoul", "Delhi", "New York"}
	tp := lib.GetTotalPopulation(cities)

	ok := tp == 17500000+14300000+17800000
	fmt.Println(tp, ok) // 49600000 true

	/**********************************************************************************************************/
	// With Dependency Inversion (Unit tests work on Dummy Data)
	/**********************************************************************************************************/
	names := []string{"alpha", "gamma"}
	total := lib.GetTotalPopulationEx(&lib.DummyDatabase{}, names)

	works := total == 4
	fmt.Println(total, works) // 4 true
}
```

Thus, the problem with singleton is not singleton itself, but depending directly upon the singleton as opposed to depending on some interface that the singleton implements, because if you depend on an interface, you can then substitute that interface. In order to able to substitute something, you need to provide an API where that something can be plugged in. In the above example, that is exactly what we did in the `GetTotalPopulationEx` function, which also takes the `Database` interface upon which it operates.

---

# Summary

- Lazy one-time initialization and thread safety using `sync.Once`
- Adhere to Dependency Inversion Principle: Depend on interfaces, not concrete types
