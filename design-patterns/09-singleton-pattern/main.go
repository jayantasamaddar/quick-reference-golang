package main

import (
	"fmt"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/09-singleton-pattern/lib"
)

func main() {
	/************************************************************************/
	// (1) Demonstration of Singleton
	/************************************************************************/
	db := lib.GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")

	fmt.Println("Population of Seoul:", pop) // Population of Seoul: 17500000

	/**********************************************************************************************************/
	// (2) Problems with Singleton (Unit tests are a problem as they work on the Live Database)
	/**********************************************************************************************************/
	cities := []string{"Seoul", "Delhi", "New York"}
	tp := lib.GetTotalPopulation(cities)

	ok := tp == 17500000+14300000+17800000
	fmt.Println(tp, ok) // 49600000 true

	/**********************************************************************************************************/
	// (3) With Dependency Inversion (Unit tests work on Dummy Data)
	/**********************************************************************************************************/
	names := []string{"alpha", "gamma"}
	total := lib.GetTotalPopulationEx(&lib.DummyDatabase{}, names)

	works := total == 4
	fmt.Println(total, works) // 4 true
}
