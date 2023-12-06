# Builder Pattern

## Introduction

The motivation for the Builder Pattern is that some objects (`structs` in Go) are simple and can be created in a single constructor call or a Factory function or by initializing the fields.

In some situations, this works fine. In other situations, objects require some ceremony to create.

For e.g. Having a factory function with 10 arguments is not very productive.

We can work around it and make the construction process, a kind of a multi-stage process, i.e. we construct an object piece-wise than try to do everything in a single factory call.

The Builder Pattern solves a situation when object construction is complicated, by providing an API for constructing an object step-by-step succinctly.

## Buider Patterns in Go

### `strings.Builder`

The strings builder is an example of the Builder pattern implementation in the core language of Go. It needs to be imported from the `strings` module.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
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
}
```

This may seem a bit complicated to build a HTML element. Ideally we would want to simply add this to some sort of a Tree Data structure representing HTML nodes and then construct the HTML out of it.
