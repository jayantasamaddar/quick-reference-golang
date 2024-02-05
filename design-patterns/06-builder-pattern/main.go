package main

import (
	"fmt"
	"strings"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/06-builder-pattern/lib"
)

/*********************************************************************************************************/
// Main Function
/*********************************************************************************************************/
func main() {
	// (1) Demonstration of `strings` builder
	hello := "Hello World!"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")

	fmt.Println(sb.String())

	words := []string{"Jack", "Paul", "Sam"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())
	sb.Reset()

	// (2a) Demonstration of HTMLBuilder
	b := lib.NewHTMLBuilder("ul")
	fruits := []string{"Apple", "Banana", "Mango"}
	for _, v := range fruits {
		b.AddChild("li", v)
	}
	fmt.Println(b.String())

	// (2b) Demonstration of HTMLBuilder with Fluent Interface (AddChildFluent)
	b.Clear()
	b.AddChildFluent("li", "One").AddChildFluent("li", "Two").AddChildFluent("li", "Three")
	fmt.Println(b.String())

	// (3) Demonstration of Builder Facets
	pb := lib.NewPersonBuilder("Sherlock Holmes")
	pb.Lives().At("221B Baker Street").In("London").Zip("NW16XE").Works().At("Science of Deduction").Designation("Private Investigator").AnnualIncome(200000)

	fmt.Println(pb.Build()) // &{Sherlock Holmes 221B Baker Street London NW16XE Science of Deduction Private Investigator 200000}

	// (4) Demonstration of Builder Parameters
	lib.SendEmail(func(b *lib.EmailBuilder) {
		b.From("foo@bar.com").To("bar@baz.com").Subject("Meeting!").Body("Hello do you want to meet?")
	})

	// (5) Demonstration of Functional Builder
	cb := lib.CountryBuilder{}
	fmt.Println(cb.Name("India").Currency("INR").Build()) // &{India INR}
}
