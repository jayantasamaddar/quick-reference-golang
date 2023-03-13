# Table of Contents

- [Table of Contents](#table-of-contents)
- [The History of Go](#the-history-of-go)
- [Go - The Programming Language](#go---the-programming-language)
- [Workspace](#workspace)
- [Modules](#modules)
- [Running Go Applications](#running-go-applications)
	- [Running your first Go Application](#running-your-first-go-application)
	- [Live Reloading Go Applications](#live-reloading-go-applications)
- [Variables](#variables)
	- [Declaring Variables](#declaring-variables)
	- [Redeclaration and Shadowing](#redeclaration-and-shadowing)
	- [Visibility](#visibility)
	- [Naming Rules and Conventions](#naming-rules-and-conventions)
	- [Type Conversions](#type-conversions)
- [Primitives](#primitives)
	- [Boolean](#boolean)
	- [Numeric Types](#numeric-types)
		- [Integers](#integers)
			- [Integers: Types](#integers-types)
			- [Integers: Arithmetic Operations](#integers-arithmetic-operations)
			- [Integers: Bitwise Operators](#integers-bitwise-operators)
			- [Integers: Bitshift Operators](#integers-bitshift-operators)
		- [Floats](#floats)
			- [Floats: Types](#floats-types)
			- [Floats: Arithmetic Operations](#floats-arithmetic-operations)
		- [Complex Numbers](#complex-numbers)
			- [Complex Numbers: Types](#complex-numbers-types)
			- [Complex Numbers: Arithmetic Operations](#complex-numbers-arithmetic-operations)
	- [Text Types](#text-types)
		- [Strings](#strings)
			- [Strings: Overview](#strings-overview)
			- [Strings: Operations](#strings-operations)
		- [Runes](#runes)
			- [Runes: Overview](#runes-overview)
- [Constants](#constants)
	- [Typed Constants](#typed-constants)
	- [Untyped Constants](#untyped-constants)
	- [Enumerated Constants](#enumerated-constants)
	- [Enumeration Expressions](#enumeration-expressions)
			- [Bitshifting to enumerate by to the power](#bitshifting-to-enumerate-by-to-the-power)
		- [Bitshifting using Boolean flags](#bitshifting-using-boolean-flags)
- [Arrays and Slices](#arrays-and-slices)
	- [Arrays: Creation](#arrays-creation)
	- [Arrays: Built-in Functions](#arrays-built-in-functions)
	- [Arrays: Working with Arrays](#arrays-working-with-arrays)
	- [Slices: Creation](#slices-creation)
	- [Slices: Built-in Functions](#slices-built-in-functions)
	- [Slices: Working with Slices](#slices-working-with-slices)
- [Maps](#maps)
	- [Maps: Overview](#maps-overview)
	- [Maps: Creation](#maps-creation)
	- [Maps: Manipulation](#maps-manipulation)
- [Structs](#structs)
	- [Structs: Creation](#structs-creation)
	- [Structs: Naming Convention](#structs-naming-convention)
	- [Structs: Embedding](#structs-embedding)
	- [Structs: Tags](#structs-tags)
- [Conditionals](#conditionals)
	- [If Statements](#if-statements)
		- [If, If-else, else Statements](#if-if-else-else-statements)
		- [Using the Initializer Syntax](#using-the-initializer-syntax)
		- [Equality with Floating Point Numbers](#equality-with-floating-point-numbers)
	- [Switch Statements](#switch-statements)
		- [Simple Cases](#simple-cases)
		- [Cases with Multiple Tests](#cases-with-multiple-tests)
		- [Using an Initializer](#using-an-initializer)
		- [Tagless Switch Statement](#tagless-switch-statement)
		- [Falling Through](#falling-through)
		- [Type Switch](#type-switch)
- [Loops](#loops)
- [Defer, Panic and Recover](#defer-panic-and-recover)
	- [Defer](#defer)
	- [Panic](#panic)
	- [Recover](#recover)
- [Pointers](#pointers)
	- [Creating Pointers](#creating-pointers)
	- [Working with `nil`](#working-with-nil)

---

# The History of Go

Go was created at Google by Robert Griesemar, Rob Pike and Ken Thompson. But one of the questions that we need to ask is, "Why create a new programming language at all?".

At the time that Go was designed, there were three languages that were being used at Google - Python, Java and C/C++. However, the Go designers started to recognize that there were some limitations that Google was running into, that might not be able to be fixed, given the history and the designs of the existing programming languages.

For example,

- **Python** is very easy to use, but it's an interpreted language, hence, slow at Google scale.
- **Java** is very quick, but its type system has become increasingly complex over time. This is a natural trend a lot of languages go through - they start out simple but as additional use cases become common, additional features are layered on top of the language, it becomes increasingly more difficult to navigate.
- **C/C++** is fast but it suffers from a complex type system and additionally its compile times are notoriously slow. The type system however has been receiving a lot of attention lately in the C/C++ communities, however there is still the burden of needing to manage/migrate legacy code. Hence, just like Java, its very difficult to move past the history of what they have, because C++ applications written 10 years ago, still need to be compiled today. The slow compile times are another legacy issue that C/C++ have inherited.

When C/C++ were designed, computers didn't have nearly the memory they have today, so the decision was made to optimize the compilers to use a minimum amount of memory and one of the compromises that brought about was that compile times can be a bit sluggish. In addition, all three languages were created in a time when multi-threaded applications were extremely rare. Almost every application that was created, really focussed on a single thread at a time. So concurrency patterns built into these languages are patched in at best. So working in highly parallel, highly concurrent applications like Google often runs into can be a little bit challenging in these languages.

Enter Go.

---

# Go - The Programming Language

Go is strong and statically typed language similar to Java and C++.
Strongly typed means, the type of a variable cannot change over time. Statically typed means, all these variables have to be defined at compile time. Go does have features to go around its type system, but 99% of the time, you are going to be using Go's strong and statically typed environment.

**Key features:**

- Strong and Statically Typed
- Excellent Community
- Key principles
  - Simplicity
  - Extremely fast compile times
  - Garbage collection baked in
  - Built-in concurrency in the base language.
  - Compile to standalone binaries. Which means when you compile your Go application, everything is going to be bundled into that single binary that's related to the Go application itself. So, the Go runtime, any Go dependencies, all get bundled in. There is thus, no need to reach out to external libraries, DLLs and other things like that to make your application work. The benefit is that it makes version management at runtime becomes trivial.

---

# Workspace

It is preferred to create a workspace inside a src folder that looks like this:

```bash
├── bin
├── pkg
├── README.md
└── src
    └── github.com
        └── jayantasamaddar
            └── firstapp
                └── Main.go
```

---

# Modules

The naming convention when declaring a module is to name it according to the repository from which it would be downloaded. For e.g. you can initialize a module by typing the following:

```bash
go mod init github.com/[github-username]/[repository]
```

---

# Running Go Applications

## Running your first Go Application

A small Hello World application in Go:

```go
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```

**Running the Go application can be done in three ways**:

1. `go run [path]`
2. `go build [modulePath]` and then running the executable `./name`
3. `go install [moduePath]` and then accessing it via the bin directory like `GOPATH/bin/name`

---

## Live Reloading Go Applications

Live Reloading Go Applications can be done using any one of the following methods:

1. Using `Nodemon`: Those coming from a JavaScript background can use:

   ```s
   nodemon --exec go run [filename.go] --signal SIGTERM
   ```

   Nodemon also works for any executable file with the `--exec` flag.

   > **Note**: Nodemon needs to be installed globally using npm: `npm i -g nodemon`

---

# Variables

## Declaring Variables

Here are the rules when it comes to variables:

1. Every variable in Go has an initialization value.

Variables can be declared in the following ways:

```go
package main

import "fmt"

func main() {
	/** (1) Initializing a variable with a type and then assigning.
	 * Useful when not ready to use a variable but can declare it.
	*/
    var i int
    i = 40
	/** (2) Initializing a variable with a type and immediately assigning it.
	 * Useful when go doesn't have enough information to infer a type and assign it.
	*/
    var j float32 = 50
	/** (3) Initializing a variable and letting the Go compiler infer type.
	 * Simple, but there is no explicit way to set the type.
	 * Cannot be used to declare variables at the package level.
	 */
    k:= 60
	l:= 70.

	/** Variable blocks */
	var (
		first_name string = "Jayanta"
		last_name string = "Samaddar"
		age int = 31
	)

    /** Using the fmt module. Check: https://pkg.go.dev/fmt for usage  */
	fmt.Println(i)					// 40
    fmt.Printf("%v, %T\n", j, j) 	// 50, float32
	fmt.Printf("%v, %T\n", k, k) 	// 60, int
	fmt.Printf("%v, %T\n", l, l) 	// 70, float64 | Go infers the type as float64 and not float32
	fmt.Printf("%[2]v %[1]v, %[3]d yrs", first_name, last_name, age) // Samaddar Jayanta, 31 yrs
}
```

---

## Redeclaration and Shadowing

When it comes to redeclaration, these are the rules to remember:

1. Variables with the same name declared in the same scope cannot be redeclared.
2. Variables may however, be reassigned but to the same type which was declared.
3. Variables with the same name maybe declared more than once in different scopes - for e.g. once in the package scope and once within the function scope. In this case, the variable with the innermost scope takes precedence.
4. Variables always have to be used. Go throws a declared and not used compile time error to keep the program clean.

```go
package main

import "fmt"

var i int = 20

func main() {
    fmt.Println(i)          // 20
    var i int = 30
    fmt.Println(i)          // 30
    // i:= 50               // throws an error
    i = 40
    fmt.Println(i)          // 40
}
```

## Visibility

There are three levels of visibility of variables in Go.

1. **Package Scope**: If a variable is declared in lowercase at the package level, it is scoped to the package and any file in the same package can access it.

2. **Global Scope**: If a variable is declared in uppercase at the package, it is exported from the package and is globally visible.

3. **Block Scope**: When a variable is declared within a function block inside one of the files, that variable is scoped to that block and is not visible outside the block.

---

## Naming Rules and Conventions

**Rules**:

1. Lowercased variables at the package level are package scoped.
2. Uppercased variables at the package level are exported to the global scope.

**Conventions** (Not hard rules):

1. Follow `camelCase` or `PascalCase` naming convention.
2. Variable names that contain acronyms like `URL` should be uppercased.
3. The length of the variable name should ideally reflect the life of the variable.

---

## Type Conversions

Go doesn't do implicit type conversion found in interpreted languages like Python or JavaScript.

Instead, in Go, we use the expression `T(v)` to convert the value `v` to the type `T`.

```go
package main

import "fmt"

func main() {
    var i int = 40
	fmt.Printf("%v, %T\n", i, i)        // 40, int

	var j float32 = float32(i)
	fmt.Printf("%v, %T\n", j, j)        // 40, float32

    var k float32 = 50.5
	fmt.Printf("%v, %T\n", i, i)        // 50.5, float32

	var l int = int(k)
	fmt.Printf("%v, %T\n", l, l)        // 50, int
}
```

While this works well for number data types, when it comes to strings, it works differently.

```go
package main

import "fmt"

func main() {
    var i int = 65
	fmt.Printf("%v, %T\n", i, i)        // 65, int

	var j string = string(i)
	fmt.Printf("%v, %T\n", j, j)        // A, string
}
```

A string is just an alias for a stream of bytes. So what happens when we asked the function to convert the integer `65` into a string, it looks for what ASCII character is set to the value `65` and that happens to be the character `A` (capital).

To convert the number `65` into a string version of it, i.e. `"65"`, we need to use the `strconv` package and use its `Itoa` method.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var i int = 65
	fmt.Printf("%v, %T\n", i, i)        // 65, int

	var j string = strconv.Itoa(i)		// ItoA = I(integer) to A(ASCII string)
	fmt.Printf("%v, %T\n", j, j)        // 65, string

	k, err := strconv.Atoi(j)			// Atoi = A(ASCII string) to i(Integer)
	fmt.Printf("%v, %T\n", k, k)        // 65, int
	fmt.Printf("%v, %T\n", err, err)	// <nil>, <nil>

	/**
	 * The parse functions return the widest type (float64, int64, and uint64),
	 * but if the size argument specifies a narrower width the result can be converted
	 * to that narrower type without data loss:
	*/

	l, err := strconv.ParseInt(j, 10, 64)
	m:= int(l)
	fmt.Printf("%v, %T\n", l, l)        // 65, int64
	fmt.Printf("%v, %T\n", err, err)	// <nil>, <nil>
	fmt.Printf("%v, %T\n", m, m)        // 65, int
}
```

---

# Primitives

## Boolean

1. A Boolean variable declared but not assigned will have its default value to be 0, i.e. false.
2. A Boolean type is returned as the result of a logical test.

```go
package main

import (
	"fmt"
)

func main() {
    var a bool
	var b bool = true

    fmt.Printf("%v, %T\n", a, a)			// false, bool
	fmt.Printf("%v, %T\n", b, b)			// true, bool

	/** A boolean is returned as the result of a Logical Test */
	x := 30
	y := 40
	gt := x > y
	lt := x < y
	eq := x == y

	fmt.Printf("%v, %T\n", gt, gt)			// false, bool
	fmt.Printf("%v, %T\n", lt, lt)			// true, bool
	fmt.Printf("%v, %T\n", eq, eq)			// false, bool
}
```

---

## Numeric Types

### Integers

#### Integers: Types

The different types of Integers we can work with are:

1. **Signed Integers**

- `int`: The general integer type that is of minimum 32 bits regardless of environment but can stretch to a maximum of 128 bits depending on the system you are running on.
- `int8`: Integers between `-128` to `127`
- `int16`: Integers between `-32,768` to `32,767`
- `int32`: Integers between `-2,147,483,648` to `2,147,483,647`
- `int64`: Integers between `-9,223,372,036,854,775,808` to `-9,223,372,036,854,775,807`

2. **Unsigned Integers**

Unsigned Integers (often called `uints`) are just like integers (whole numbers) but have the property that they don't have a `+` or `-` sign associated with them. Thus they are always non-negative (zero or positive). We use `uints` when we know the value we are counting will always be non-negative. For example, if we are counting the number of players in a game, we could use a uint because there will always be 0 or more players.

> Note: it is almost always the case that you could use a regular integer variable in place of an unsigned integer. The advantage to using the unsigned version (when you know the values contained will be non-negative) is that sometimes the computer will spot errors for you (the program will "crash" when a negative value is assigned to the variable).

- `uint`: `0` to `255`
- `uint8`: `0` to `255`
- `byte`: `0` to `255` (alias for `uint8`)
- `uint16`: `0` to `65,536`
- `uint32`: `0` to `4,294,967,295`

---

#### Integers: Arithmetic Operations

1. Integer Operations always have to be done with the same integer type.
2. We can do all arithmetic operations with Integers of the same type. However an integer only operation returns an integer type.

```go
package main

import "fmt"

func main() {
    a:= 10
	b:= 3

	fmt.Println(a + b)			// 13
	fmt.Println(a - b)			// 7
	fmt.Println(a * b)			// 30
	fmt.Println(a / b)			// 3    => An integer divided by an integer returns an integer
	fmt.Println(a % b)			// 1    => Remainder
}
```

As you can see, the integer division did not give us a result where there are decimals as `10` divided by `3` is `3.3333333` but since we do an integer division, the remainder is dropped.

> **Note**: To get the decimal we can simply do a type conversion to `float32` for each value before we do the operation.

```go
package main

import "fmt"

func main() {
    a:= 10
    b:= 3
    fmt.Println(float32(a) / float32(b))        // 3.3333333
}
```

---

#### Integers: Bitwise Operators

Compares bits of the first number to the bits of the second number as per the operation.

```go
package main

import "fmt"

func main() {
    a:= 10						// Binary: 1010
	b:= 3						// Binary: 0011

	// Both have the bit set
	fmt.Println(a & b)			// 2  => Binary: 0010 		= 2
	// Either one has the bit set
	fmt.Println(a | b)			// 11 => Binary: 1011 		= 11
	// One or the other have the bit set, but not both
	fmt.Println(a ^ b)			// 9 => Binary: 1001 		= 9
	// Neither have the bit set
	fmt.Println(a &^ b)			// 8 => 0100 				= 8
}
```

---

#### Integers: Bitshift Operators

```go
package main

import "fmt"

func main() {
    a:= 8						// 2^3

	// Bitshifts left 3 places
	fmt.Println(a << 3)			// 2^3 * 2^3				= 2^6 = 64
	// Either one has the bit set
	fmt.Println(a >> 3)			// 2^3 / 2^3 				= 2^0  = 1
}
```

---

### Floats

In Go, Floating Point Numbers or Floats follow **IEEE-754 Standard**. We can use either decimal numbers or exponential notation.

#### Floats: Types

We have the following types:

1. `float32`: 32-bit Floating Point Numbers
2. `float64`: 64-bit Floating Point Numbers

```go
/** Initializing Floating Point Numbers */
package main

import "fmt"

func main() {
    a:= 3.14
	var b float64 = 13.7e72
	var c float32 = 2.1E14

	fmt.Printf("%v, %T\n", a, a)    // 3.14, float64 (float64 is default when initializer syntax)
	fmt.Printf("%v, %T\n", b, b)	// 1.37e+73, float64
	fmt.Printf("%v, %T\n", c, c)	// 2.1e+14, float32
}
```

---

#### Floats: Arithmetic Operations

1. Floating Point Operations always have to be done with the same floating point type. (i.e. float32 type cannot be added to a float64 type. Type conversion must be done prior to that.)
2. We can do all arithmetic operations with Floating Point numbers of the same type. The return type is always a floating point number type.
3. The Modulo operator (`%`), and Bitwise and Bitshift operators are not available. They are only available to Integer types.

```go
package main

import "fmt"

func main() {
    a:= 10.2
	b:= 3.7

	fmt.Println(a + b)			// 13.899999999999999
	fmt.Println(a - b)			// 6.499999999999999
	fmt.Println(a * b)			// 37.74
	fmt.Println(a / b)			// 2.7567567567567566
}
```

---

### Complex Numbers

It is rare for programming languages to treat complex numbers as a first class citizen. This opens Go to be used as a powerful language for Data Science.

#### Complex Numbers: Types

There are two types of Complex Numbers in Go.

- `complex64`: Takes a `float32` + `float32` for the real and imaginary parts
- `complex128`: Takes a `float64` + `float64` for the real and imaginary parts

> **Note**:
>
> - We can use wrap a complex number in `real()` or `imaginary()` to get the real and imaginary parts out respectively.
> - Similarly we can take two numbers (representing the real and imaginary part respectively) and convert them into a complex number using the `complex()` function that takes two numbers as arguments.
> - Go's parser understands the `i` as an imaginary number.

```go
package main

import "fmt"

func main() {
    var a complex64 = 1 + 2i
	var b complex64 = 2i

	fmt.Printf("%v, %T\n", a, a)					// (1+2i), complex64
	fmt.Printf("%v, %T\n", b, b)					// (0+2i), complex64
	fmt.Printf("%v, %T\n", real(a), real(a))		// 1, float32
	fmt.Printf("%v, %T\n", imag(b), imag(b))		// 2, float32
    fmt.Println(complex(1, 2))						// (1+2i)
}
```

---

#### Complex Numbers: Arithmetic Operations

```go
package main

import "fmt"

func main() {
    var a complex64 = 1 + 2i
	var b complex64 = 2 + 5.2i

	fmt.Println(a + b)			// (3+7.2i)
	fmt.Println(a - b)			// (-1-3.1999998i)
	fmt.Println(a * b)			// (-8.4+9.2i)
	fmt.Println(a / b)			// (0.39948454-0.03865979i)
}
```

---

## Text Types

### Strings

#### Strings: Overview

In Go,

- A String can be any UTF-8 character.
- We declare strings with double-quotes. `"This is a string"`
- Strings in Go are aliases for bytes
- Strings are immutable.

```go
package main

import "fmt"

func main() {
    s:= "This is a string"
    /** Get the third character of the string */
	s2:= s[2]
	// s[2] = "L"										// Unassignable operand. Immutable.
	fmt.Printf("%v, %T\n", s, s)						// This is a string, string
	fmt.Printf("%v, %T\n", s2, s2)						// 105, uint8
	fmt.Printf("%v, %T\n", string(s2), string(s2))		// i, string
}
```

---

#### Strings: Operations

```go
package main

import "fmt"

func main() {
    s:= "This is a string"

    /** Concatenating Strings */
	fmt.Println(s + ". " + "Okay!")						// This is a string. Okay!
    fmt.Printf("%v. Okay!\n", s)						// This is a string. Okay!

    /** Convert a string into a collection of bytes */
    b:= []byte(s)
    fmt.Printf("%v, %T\n", b, b)                        // [84 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8
}
```

> **Important**: A lot of the functions that we use in Go, work with byte slices. That makes them much more generic and much more flexible, than if we worked with hard coded strings. For example, if you want to send a string response to a web request, you can easily convert it to a collection of bytes. If you want to send a file back, even a file is just a collection of bytes too. Thus it allows transparency to work with strings without worrying about line endings and things like that. So while, we might be working with strings in our Go programs as strings, when we start sending them to other applications, we may just be sending them as byte slices.

---

### Runes

#### Runes: Overview

In Go,

- A Rune is UTF-32 character.
- We declare runes with single-quotes. `'This is a rune'`
- Runes in Go are aliases for `int32`.
- While strings can be converted back and forth between collection of bytes and string literal, runes in Go are a true type alias of `int32`. It is the same as a `int32`.

> **Note**: While a String type represents any UTF-8 character, a rune represents any UTF-32 character. UTF-32 is a little weird as while any character in UTF-32 can be 32-bits long, it doesn't have to be 32-bits long. Any UTF-8 character which is 8-bits long, is also a UTF-32 character.

```go
package main

import "fmt"

func main() {
    r:= 'r'

	fmt.Printf("%v, %T\n", r, r)			// 114, int32
}
```

> **Note**: Runes are useful when working with UTF-32 characters. Read more about runes in the [ReadRune](https://pkg.go.dev/strings#Reader.ReadRune) API documentation in the Go `strings` package.

---

# Constants

- All constants are declared with the `const` keyword.
- Has to be assigned at compile time and not runtime. So assigning a function return value to a constant is going to throw an error.
- Immutable and cannot be reassigned. Thus arrays which are mutable cannot be constants.
- If we come from other programming languages, we may think of declaring constants with all uppercase letters like: `const MY_CONST`. However with Go, as discussed earlier, all uppercase indicates global scope, hence we declare constant variable names just how we declare normal variables: either in camelCase if we didn't want to globally export it or PascalCase if we wanted to export it globally.
- Can be shadowed.
- Can run operations with constants or variables of the same type. The return value is a variable.
- Doesn't throw an error if unused unlike variables.

---

## Typed Constants

Typed constants are declared just like typed variables but with the `const` keyword instead of the `var` keyword. Typed constants offer type safety that restricts interoperability only to the same type unless there is a type conversion done.

```go
package main

import "fmt"

const a int32 = 10
const b float32 = 10
const c = a + int32(b)                  // Type conversion of the float32 type to int32
// const d = math.Sin(math.Pi/4)        // (value of type float64) is not a constant
const e int8 = 255
// const x                              // missing init expr for x

func main() {
    var v int32 = 30
    fmt.Printf("%v, %T\n", a, a)        // 10, int32
	fmt.Printf("%v, %T\n", b, b)        // 10, float32
    // fmt.Println(a + b)               // invalid operation:(mismatched types int32 and float32)
    fmt.Printf("%v, %T\n", sum, sum)    // 20, int32
    fmt.Printf("%v, %T\n", e, e)		// 255, uint8

	const e uint8 = 128                 // (Can be shadowed)
	fmt.Printf("%v, %T\n", e, e)		// 128, uint8
    fmt.Printf("%v, %T\n", a+v, a+v)	// 40, int32 (Valid operation using const and var)
}
```

---

## Untyped Constants

Untyped constants refer to constants whose type is determined by the Go compiler and not specified at declaration. Constants are replaced by their value at runtime by the Go compiler hence untyped constants can interoperate with similar types.

```go
package main

import "fmt"

const a = 10
var v int16 = 35

func main() {
    /** Constants are replaced by their values at runtime: e.g. (10 + v) in this case */
    fmt.Printf("%v, %T\n", a + v, a + v)		// 45, int16
    fmt.Printf("%v, %T\n", 10 + v, 10 + v)		// 45, int16 (Same as above)
}
```

---

## Enumerated Constants

In Go, enumerated constants are created using the `iota` enumerator. Since `iota` can be part of an expression and expressions can be implicitly repeated, it is easy to build intricate sets of values. E.g. we can use `iota` as a counter.

Here are some other characteristics:

- The value of `iota` is block scoped.
- Initial value of `iota` is `0`.
- If we do not assign the value `iota`, the compiler is going to try to infer the value of the pattern.

```go
package main

import "fmt"

const i = iota
const (
	a = iota
	b = iota
	c = iota
    d
    e
)
const (
    x = iota + 10                       // All integer operations are permitted
    y
    z
)

func main() {
    const j int = y
    fmt.Printf("%v, %T\n", i, i)		// 0, int
    /** Using `iota` as a counter */
    fmt.Println(a)						// 0
	fmt.Println(b)						// 1
	fmt.Println(c)						// 2
	fmt.Println(d)						// 3
	fmt.Println(e)						// 4
    fmt.Println(x, y, z)				// 10 11 12
    fmt.Println(j == y)					// true
}
```

> **Note**: Often a situation requires you to compare an iota value to an integer. However, an unassigned integer value is also `0`. To avoid this situation, we can determine the first iota value to be an error value.
>
> Assigning a `_` name to the `0` value of the enumerated constant assigns it to a blank identifier

```go
package main

import "fmt"

const (
	_  = iota                           // ignore first value by assigning to blank identifier
    x
    y
    z
)

func main() {
	var j int
	fmt.Println(x, y, z)				// 1 2 3
	fmt.Println(j == x)					// false
}
```

---

## Enumeration Expressions

#### Bitshifting to enumerate by to the power

Bitshifting enumerated constants, is one of the ways to use Enumerated constants. The reason in because we don't have the ability to raise to power of an enumerated constant, as raising to the power in Go, is a function in the `math` package, so we cannot do that in our constant enumeration. We can overcome this by using Bitshift operators.

```go
package main

import "fmt"

const (
    _  = iota 							// ignore first value by assigning to blank identifier
    KB = 1 << (10 * iota)				// 2^0 * 2^(10 * 1) = 1024
    MB									// 2^0 * 2^(10 * 2) = 1048576
    GB									// 2^0 * 2^(10 * 3) = 1073741824
    TB									// 2^0 * 2^(10 * 4) = 1099511627776
    PB									// 2^0 * 2^(10 * 5) = 1125899906842624
    EB									// 2^0 * 2^(10 * 5) = 1152921504606846976
)


func main() {
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)			// 3.73GB
	fmt.Println(KB)								// 1024
	fmt.Println(MB)								// 1048576
	fmt.Println(GB)								// 1073741824
	fmt.Println(TB)								// 1099511627776
	fmt.Println(PB)								// 1125899906842624
	fmt.Println(EB)								// 1152921504606846976
}
```

---

### Bitshifting using Boolean flags

We can use Bitshifting to set boolean flags inside of a single byte. This makes memory efficient code. In the code below, we can store 8 access roles for users into a single byte of data.

```go
package main

import "fmt"

const (
    isAdmin = 1 << iota					// 2^0 * 2^0 	= 1
    isHeadquarters						// 2^0 * 2^1 	= 2
    canSeeFinancials					// 2^0 * 2^2 	= 4

    canSeeAfrica						// 2^0 * 2^3 	= 8
    canSeeAsia							// 2^0 * 2^4 	= 16
    canSeeEurope						// 2^0 * 2^5 	= 32
    canSeeNorthAmerica					// 2^0 * 2^6 	= 64
    canSeeSouthAmerica					// 2^0 * 2^7 	= 128
)

func main() {
	/**
	 *  isAdmin 														= 00000001
	 *	canSeeFinancials 												= 00000100
	 *	canSeeEurope													= 00100000
	 *  isAdmin | canSeeFinancials | canSeeEurope 						= 00100101
	*/
    var roles byte = isAdmin | canSeeFinancials | canSeeEurope
    fmt.Printf("%b\n", roles)											// 100101
	fmt.Printf("isAdmin: %v\n", isAdmin & roles == isAdmin)				// isAdmin: true
	fmt.Printf("canSeeAsia: %v\n", canSeeAsia & roles == canSeeAsia)	// canSeeAsia: false
}
```

---

# Arrays and Slices

Arrays and Slices are two collection types available in Go.

## Arrays: Creation

With Arrays, by the design of the language, the elements are contiguous in memory. So collecting our data together in arrays, not only is it easier to work with, but also makes our applications fast.

In Go, unlike C

1. Arrays are values. Assigning one array to another copies all the elements.
2. In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
3. The size of an array is part of its type. The types [10]int and [20]int are distinct.

```go
package main

import "fmt"

func main() {
	var empty [3]string
	nums := [3]int{100, 200, 300}
	// The dots mean, create an array just long enough to hold the elements passed
	scores := [...]int{100, 200, 300, 400, 500}

	fmt.Printf("%v, %T\n", empty, empty)        // [  ], [3]string
    fmt.Printf("%v, %T\n", nums, nums)          // [100 200 300], [3]int
	fmt.Printf("%v, %T\n", scores, scores)		// [100 200 300 400 500], [5]int
    /** Assigning to arrays: Assigns the value to index 1. Indexes start at 0.  */
    scores[1] = 150
    fmt.Printf("%v, %T\n", scores, scores)		// [100 150 300 400 500], [5]int
}
```

---

## Arrays: Built-in Functions

```go
package main

import "fmt"

func main() {
	scores := [...]int{100, 200, 300, 400, 500}

    fmt.Printf("%v, %T\n", scores, scores)		// [100 200 300 400 500], [5]int
	/** Get the length of the array */
	fmt.Println(len(scores))					// 5
}
```

> **Note**: Use the index `len(a)-1` to access the last element of a slice or array `a`. Go doesn't have negative indexing like Python does.

---

## Arrays: Working with Arrays

Here are some other ways we can work with arrays:

```go
package main

import "fmt"

func main() {
	/** A 2-dimensional array */
	var matrix = [3][3]int{ {1, 0, 0}, {0, 1, 0}, {0, 0, 1} }
    fmt.Printf("%v, %T\n", matrix, matrix)		// [[1 0 0] [0 1 0] [0 0 1]], [3][3]int

	/** Array passed as a new copy */
	a := [...]int{1, 2, 3}
	b := a
	b[2] = 4
	fmt.Printf("a: %v\n", a)					// a: [1 2 3]
	fmt.Printf("b: %v\n", b)					// b: [1 2 4]   => Original array is not mutated

	/** Array passed as a pointer */
	c := [...]int{1, 2, 3}
	d := &c
	d[2] = 5
	fmt.Printf("c: %v\n", c)					// c: [1 2 5]
	fmt.Printf("d: %v\n", d)					// d: &[1 2 5]  => Original array is mutated
}
```

Arrays are very useful and there are use cases for arrays. However the fact that their size is limited and must be known at compile time, limits their usefulness. In Go, the most common use case for an array is to back a **Slice**.

---

## Slices: Creation

Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimension such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.

There are three ways to create a Slice:

1. Using the empty square brackets to denote a slice at initialization
2. Using the `:` syntax
3. Using a built-in-function: `make`

```go
package main

import "fmt"

func main() {
    // A slice is initialized with the empty square brackets
	var empty []string
	nums := []int{100, 200, 300}
	s := []int{100, 200, 300, 400, 500}

	fmt.Printf("%v, %T\n", empty, empty)        // [], []string
    fmt.Printf("%v, %T\n", nums, nums)          // [100 200 300], []int
	fmt.Printf("%v, %T\n", s, s)		        // [100 200 300 400 500], []int
    /** Assigning to slices: Assigns the value to index 1. Indexes start at 0.  */
    s[1] = 150
    fmt.Printf("%v, %T\n", s, s)		        // [100 150 300 400 500], []int

	/**
	 * Another way of creating a slice following the syntax: s[indexIncluding:indexExcluding]
	*/
    s1 := s[:]                                  // slice of all elements
    s2 := s[3:]                                 // slice from 4th element to end
    s3 := s[:4]                                 // slice first 5 elements
    s4 := s[3:5]                                // slice the 4th, 5th
    s5 := s[1:len(s) - 1]                       // slice all elements leaving the first and last

    fmt.Println(s1)                             // [100 150 300 400 500]
    fmt.Println(s2)                             // [400 500]
    fmt.Println(s3)                             // [100 150 300 400]
    fmt.Println(s4)                             // [400 500]
    fmt.Println(s5)                             // [150 300 400]
}
```

---

## Slices: Built-in Functions

```go
package main

import "fmt"

func main() {
	scores := []int{100, 200, 300, 400, 500}

    fmt.Printf("%v, %T\n", scores, scores)		// [100 200 300 400 500], []int
	/** Get the length of the slice */
	fmt.Println(len(scores))					// 5
    /** Get the capacity of the slice which is the length of the underlying array.
     * The number of elements of the slice doesn't necessarily match the size of the
     * array backing it. Capacity refers to the length of the underlying array.
    */
    fmt.Println(cap(scores))                    // 5 => Underlying array is same size

    /** Push a new element to the end of a slice */
    scores = append(scores, 600)
    fmt.Printf("%v, %T\n", scores, scores)      // [100 200 300 400 500 600], []int
    fmt.Printf("Len: %v, Cap: %v\n", len(scores), cap(scores))  // Len: 6, Cap: 10

    /** Slice whose capacity remains same as its length */
    var names []string
    fmt.Printf("Len: %v, Cap: %v\n", len(names), cap(names))    // Len: 0, Cap: 0

    names = append(names, "Tom")
    fmt.Printf("Len: %v, Cap: %v\n", len(names), cap(names))    // Len: 1, Cap: 1

    names = append(names, "Jerry", "Johnny", "Bravo")
    fmt.Printf("Len: %v, Cap: %v\n", len(names), cap(names))    // Len: 4, Cap: 4
}
```

---

## Slices: Working with Slices

- Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.
- If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer to the underlying array.
- A Read function can therefore accept a slice argument rather than a pointer and a count; the length within the slice sets an upper limit of how much data to read.
- Use the index `len(a)-1` to access the last element of a slice or array `a`. Go doesn't have negative indexing like Python does.
- One has to be careful when mutating slices as the underlying array will be mutated too as slices hold reference not value.

```go
package main

import "fmt"

func main() {
    /**
     * Creating a slice using `make`:
     * Syntax: make(type, length, capacity)
    */
    a := make([]int, 3)

    fmt.Println(a)                                      // [0 0 0]
    fmt.Printf("Len: %d, Cap: %d\n", len(a), cap(a))    // Len: 3, Cap: 3

    b := make([]int, 3, 6)
    fmt.Println(b)                                      // [0 0 0]
    fmt.Printf("Len: %d, Cap: %d\n", len(b), cap(b))    // Len: 3, Cap: 6

	/** Concatenating slices */
    x := []int{1, 2, 3, 4, 5}
    y := []int{11, 12, 13, 14, 15}

    z := append(x, y...)

    fmt.Println(z)                          // [1 2 3 4 5 11 12 13 14 15]

	/** Treating the slice as a stack */
	x = x[1:]                               // Shift operation (remove first element)
    fmt.Println(x)                          // [2 3 4 5]

    x = x[:len(x) - 1]                      // Pop operation (remove last element)
    fmt.Println(x)                          // [2 3 4]

    /** Removing elements from the middle */
    x = []int{1, 2, 3, 4, 5}
    x = append(x[:1], x[len(x) - 2:]... )   // Remove elements between first and second last

    fmt.Println(x)                          // [1 4 5]
}
```

---

# Maps

## Maps: Overview

**Maps** are a convenient and powerful built-in data structure that associate values of one type (the key) with values of another type (the element or value). They can be created via literals or via `make` function.

Here are some rules when it comes to Maps

1. The key can be of any type for which the equality operator is defined, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays.
2. Slices and other maps cannot be used as map keys, because equality is not defined on them.
3. Like slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.
4. The return order of a Map is not guaranteed.

---

## Maps: Creation

Within a single map, there can be multiple keys of one Key type that can be mapped to their respectively values of only one value type.

```go
package main

import "fmt"

func main() {
    var timeZone = map[string]int{
        "UTC":  0*60*60,
        "EST": -5*60*60,
        "CST": -6*60*60,
        "MST": -7*60*60,
        "PST": -8*60*60,
    }

    fmt.Println(timeZone)           // map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
}
```

---

## Maps: Manipulation

```go
package main

import "fmt"

func main() {
    timeZone := map[string]int{
        "UTC":  0*60*60,
        "EST": -5*60*60,
        "CST": -6*60*60,
        "MST": -7*60*60,
        "PST": -8*60*60,
    }

    fmt.Println(timeZone)	// map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]

    /** Accessing the value of a key */
    fmt.Printf("UTC: %v\n", timeZone["UTC"])    // UTC: 0

	/** Adding a new key value */
	timeZone["IST"] = 5*60*60 + 0.5*60*60
	fmt.Println(timeZone)	// map[CST:-21600 EST:-18000 IST:19800 MST:-25200 PST:-28800 UTC:0]

	/** Delete a map entry. Safe to use even if key is absent. */
	delete(timeZone, "MST")
	fmt.Println(timeZone)   // map[CST:-21600 EST:-18000 IST:19800 PST:-28800 UTC:0]

	/** "comma ok" idiom to check if a value exists */
	pop := timeZone["GST"]
	fmt.Println(pop)	// 0 (If not found, always returns 0)

    // Non-existing key returns ok as `false`
	key, ok := timeZone["GST"]
	fmt.Println(key, ok)	// 0 false

    // Existing key returns ok as `true`
	ist, ok := timeZone["IST"]
	fmt.Println(ist, ok)	// 19800 true

	/** Check the number of key-value pairs */
    fmt.Println(len(timeZone))  // 5

	/** Test for presence in the map without worrying about the actual value */
	_, present := timeZone["GST"]
}
```

---

# Structs

## Structs: Creation

- Structs can be declared either as a type or as an anonymous struct.
- Can leave out optional fields.
- Structs are stored as values and not references unlike Maps.

**Declare a struct as a type**:

```go
package main

import "fmt"

type Person struct {
    id int
    name string
    friends []string
}

func main() {
    person := Person {
        id: 1,
        name: "Jayanta",
        friends: []string {
            "Adraha",
            "Rohit",
            "Ravi",
        },
    }

    fmt.Println(person)             // {1 Jayanta [Adraha Rohit Ravi]}
	fmt.Println(person.friends[1])  // Rohit
}
```

**Declare an anonymous struct**:

```go
package main

import "fmt"

func main() {
    person := struct {
        id int
        name string
        friends []string
        }{
            id: 1,
            name: "Jayanta",
            friends: []string{
                "Adraha",
                "Rohit",
                "Ravi",
            },
        }

    fmt.Println(person)             // {1 Jayanta [Adraha Rohit Ravi]}
	fmt.Println(person.friends[1])  // Rohit
}
```

**Structs as references**:

```go
package main

import "fmt"

func main() {
    person := struct {
        id int
        name string
        friends []string
        }{
            id: 1,
            name: "Jayanta",
            friends: []string{
                "Adraha",
                "Rohit",
                "Ravi",
            },
        }

    fmt.Println(person)             // {1 Jayanta [Adraha Rohit Ravi]}

	person2 := person;				// copies the struct into a new struct
	person3 := &person;				// points the person struct to person3 struct

	person2.name = "Bhargav"
	person3.name = "Prince"

    fmt.Println(person)             // {1 Prince [Adraha Rohit Ravi]}
	fmt.Println(person2)  			// {1 Bhargav [Adraha Rohit Ravi]}
	fmt.Println(person3)  			// &{1 Prince [Adraha Rohit Ravi]}
}
```

---

## Structs: Naming Convention

The naming conventions are same as variables in Go. If a struct field starts with an uppercase letter, it will be exported to be globally available or else it will be internal to the package.

---

## Structs: Embedding

Go does not provide the typical, type-driven notion of subclassing, i.e. inheritance model, but it does have the ability to "borrow" pieces of an implementation by embedding types within a struct or interface using a composition model called embedding.

- In a composition relationship a struct that embeds another struct, is still its own type.
- When we are talking of modelling behaviour, embedding is not the right choice. Embedding allows methods to carry through to the type that has the embedding. However the fact that we cannot use them interchangeably is a severe limitation. Generally, it's much better to use interfaces, when we are to describe common behaviour (have common methods).
- Embedding is a good idea, when we just want to get some basic behaviour into a custom type. For e.g. when authoring a library or a web framework. In that case, maybe we want consumers of the library to embed the base controller into the custom controller to get functionality out of it.

```go
package main

import "fmt"

type Animal struct {
    name string
    origin string
}

type Bird struct {
    Animal
    wingspan float32
    canFly bool
}

/** Declaring and manipulating from the outside */
func main() {
    bird := Bird{
        wingspan: 20.25,
        canFly: true,
    }
    bird.name = "Bald Eagle"
    bird.origin = "North America"

	/** {{Bald Eagle North America} 20.25 true}  */
    fmt.Println(bird)
	fmt.Println(bird.name) 				// "Bald Eagle"

	/** Declaring using the literal syntax */
	bird2 := Bird{
        Animal: Animal{
            name: "Bald Eagle",
            origin: "North America",
        },
        wingspan: 20.25,
        canFly: true,
    }

	/** {{Bald Eagle North America} 20.25 true}  */
    fmt.Println(bird2)
	fmt.Println(bird.name)              // Bald Eagle
}
```

---

## Structs: Tags

- Tags describe specific metadata about a field of a struct.
- All tags do is provide a string of text. It's upto the validation framework to do something with tags. By themselves tag bring in no functionality.

```go
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
    id int `required:"true"`
    name string `required:"true"`

}

func main() {
    person := Person{
        id: 1,
        name: "Jayanta",
    }
    t := reflect.TypeOf(person)
    field, _ := t.FieldByName("name")

    fmt.Println(person)                     // {1 Jayanta}
    fmt.Println(person.name)                // Jayanta
    fmt.Println(field.Tag)                  // required:"true"
    fmt.Println(field.Tag.Get("required"))  // true
}
```

---

# Conditionals

## If Statements

- The curly braces are a necessary Go syntax for if statements
- Variables declared within an if block are block-scoped to that particular block.
- First brackets are permissible to be used for the logic part of the if statement. `if a > b` and `if (a > b)` are both valid ways of declaring an if statement.

### If, If-else, else Statements

```go
package main

import (
	"fmt"
)

func main() {
    a := 12
    b := 20

    if a > b {
        fmt.Println("'a' is greater than 'b'")
    } else if a == b {
        fmt.Println("'a' is equal to 'b'")
    } else {
        fmt.Println("'b' is greater than 'a'")
    }
}
```

### Using the Initializer Syntax

- Syntax: `if initializer; boolean {}`

```go
package main

import (
	"fmt"
)

func main() {
    timeZone := map[string]int{
        "UTC":  0*60*60,
        "IST":  5*60*60 + 0.5*60*60,
        "EST": -5*60*60,
        "CST": -6*60*60,
        "MST": -7*60*60,
        "PST": -8*60*60,
    }

    if time, ok := timeZone["IST"]; ok {
        fmt.Println(time)								// executes this line
    } else {
        fmt.Println("IST not present")
    }
}
```

---

### Equality with Floating Point Numbers

When working with floating point numbers, you may run into an issue with equality because the return value of floating point operations are approximations of decimal values and not exact.

Example:

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	num = 0.123

	if num == math.Pow(math.Sqrt(num), 2) {
		fmt.Println("These are the same!")
	} else {
		fmt.Println("These are different!")				// executes this line
	}
}
```

Hence when we are doing comparison operations with decimal values, a better approach is to have a small error margin when doing the comparison.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	num := 0.1234141794174194

	if math.Abs(math.Pow(math.Sqrt(num), 2)) - 1 < 0.0001 {
		fmt.Println("These are the same!")					// executes this line
	} else {
		fmt.Println("These are different!")
	}
}
```

This isn't a perfect solution but the decimal error margin can be further tweaked to fine-tune the error margin.

---

## Switch Statements

Go's switch is more general than C's. The expressions need not be constants or even integers, the cases are evaluated top to bottom until a match is found, and if the switch has no expression it switches on true. It's therefore possible—and idiomatic—to write an if-else-if-else chain as a switch.

Some characteristics to keep in mind with Switch statements

- There can be multiple tests in a single case
- Instead of a simple tag, can use an initializer like the if statement.
- Can use a tagless syntax in which there is no value evaluated at the switch statement and the cases behave like if-else-if-else blocks.
- There cannot be overlapping cases EXCEPT when using a tagless syntax. Will throw a syntax error.
- The `break` after each case statement is implicit in Go, as a design decision, instead of the default fallthrough. However if we want to exit a single case statement early for some reason we can use the `break` keyword explicitly.
- If we do intend a fallthrough behaviour for a particular case, we can specify the keyword `fallthrough`. The fallthrough will override any case logic in the next case and execute the code directly.

---

### Simple Cases

```go
package main

import "fmt"

func main() {
	num := 2
	switch num {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")						// executes this line
		default:
			fmt.Println("other number")
	}
}
```

---

### Cases with Multiple Tests

```go
package main

import "fmt"

package main

import "fmt"

func main() {
	num := 6
	switch num {
		case 0:
			fmt.Println("Number must be greater than 0")
		case 1, 3, 5, 7, 9:
			fmt.Println("odd")
		case 2, 4, 6, 8:
			fmt.Println("even")						// executes this line
		default:
			fmt.Println("Number must be between 1 and 10")
	}
}
```

---

### Using an Initializer

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	const num = -1
	switch i := math.Max(num, 0); i {
		case 0:
			fmt.Println("Number must be greater than 0")    // executes this line
		case 1, 3, 5, 7, 9:
			fmt.Println("odd")
		case 2, 4, 6, 8:
			fmt.Println("even")
		default:
			fmt.Println("Number must be between 1 and 10")
	}
}
```

---

### Tagless Switch Statement

```go
package main

import "fmt"

func main() {
	num := 10
	switch {
		case num <= 0:
			fmt.Println("Number must be greater than 0")
        case num > 10:
            fmt.Println("Number must be between 1 and 10")
		case num % 2 == 1:
			fmt.Println("odd")
		default:
			fmt.Println("even")             // executes this line
	}
}
```

---

### Falling Through

We can use the keyword `fallthrough` to make a case statement not break when it executes but fall through to the next case statement. `fallthrough` does not evaluate any logic, and executes the next case regardless of the case logic in that particular case. It is upto the developer to implement the control flow.

```go
package main

import "fmt"

func main() {
	num := 10
	switch {
		case num <= 10:
			fmt.Println("Number is less than or equal to 10")   // executes this line
            fallthrough
        case num <= 20:
            fmt.Println("Number is less than or equal to 20")   // executes this line
		default:
			fmt.Println("Number is greater than 20")
	}
}
```

---

### Type Switch

A switch can also be used to discover the dynamic type of an interface variable. Such a type switch uses the syntax of a type assertion with the keyword type inside the parentheses. If the switch declares a variable in the expression, the variable will have the corresponding type in each clause. It's also idiomatic to reuse the name in such cases, in effect declaring a new variable with the same name but a different type in each case.

```go
package main

import "fmt"

func main() {
	var i interface{} = 10
	switch i.(type) {
		case int:
			fmt.Printf("i is int %v\n", i)                  // i is int 10
            if (i == 10) {
                fmt.Println("i is 10")                      // executes this line
                break										// breaks here
            }
            fmt.Println("Will not execute this line")       // Doesn't execute
        case bool:
            fmt.Printf("i is boolean %v\n", i)
        case string:
            fmt.Printf("i is int %v\n", i)
		default:
			fmt.Printf("i is another type: %T %v\n", i, i)
	}
}
```

---

# Loops

Things to remember while working with Loops in Go:

- The increment/decrement operation in Go is not an expression, it is a statement on its own.
- The traditional `do-while` loop can be implemented by running a for-loop with only the conditional and the counter initialized outside the for-loop and controlled from within.
- You can manually short-circuit out of a for-loop by using `break`.
- To end the particular iteration of the loop and go to the next without running the code below, we can use `continue`.

```go
package main

import "fmt"

func main() {
	/** Basic for-loop */
	for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

	for i := 0; i < 10; i+=2 {
        fmt.Println(i)
    }

	/** Looping using multiple variables */
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	/** Leaving the first statement out. Note: `i` is scoped outside of the for-loop */
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	/** This is how Go implements do-while loops */
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	/** Running loops that require complex logic and manual short-circuit */
	i := 0
	for {
        if (i == 10) {
			break
		}
        i++
		if(i % 2 == 1) {
			continue
		}
		fmt.Println(i)		// Prints the even numbers from 0 - 10, i.e. 2 4 6 8 10
	}

	/** Nested for-loops with labels */
	Loop:
		for i := 1; i <= 10; i++ {
			for j := 1; j <= 3; j++ {
				fmt.Println(i * j)
				if (i * j >= 3) {
					break Loop				// breaks the outer-loop defined as the label `Loop`
				}
			}
		}

	/** Working with collection types */
	s := []int{1, 2, 3, 4, 5}
	var timeZone = map[string]int{
        "UTC":  0*60*60,
        "EST": -5*60*60,
        "CST": -6*60*60,
        "MST": -7*60*60,
        "PST": -8*60*60,
    }

	/** Method 1: Works for strings, arrays and slices */
    for i := 0; i < len(s); i++ {
        fmt.Println(s[i])					// Prints each element out
    }

	/** Method 2: Works with strings, arrays, slices, maps and channels */

	// On slices, arrays
	for index, val := range s {
		fmt.Println(index, val)
	}

	// On maps
	for key, val := range timeZone {
        fmt.Println(key, val)				// Prints the key and value out
    }

	/** A situation where we need only the value */
	for _, val := range timeZone {
        fmt.Println(val)					// Prints the key and value out
    }
}
```

We can also range over another data type called **Channels**. Channels are used for concurrent programming in Go and a topic for a future section.

---

# Defer, Panic and Recover

## Defer

In a normal Go application, control flows from the top to the bottom of any function that we call.
With `defer` we can invoke a function but delay its execution time to a future point in time.

Here are some important points to remember about deferred functions:

- Takes a function call (not a function itself)
- The way `defer` works is that it executes any functions that are passed into it, after the function finishes its final statement but before it returns.
- Deferred functions are executed in LIFO (Last-in-first-out) order.
- Deferred functions take arguments at the time of the deferred function is called, not at the time the function is executed.
- Care should be exercised when it comes to loops.

```go
package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("middle")			// second last defer function executes second
	defer fmt.Println("middle2")		// last defer function executes first
	fmt.Println("end")
}

/** Prints

start
end
middle2
middle
*/
```

**Deferring functions whose arguments maybe modified later**

```go
package main

import "fmt"

func main() {
	a := "start"
	a = "middle"
	defer fmt.Println(a)		// defer prints "middle" as it takes the argument at the time of defer and not during execution
	a = "end"
}
```

**Use Case**: Associate the opening and the closing of a resource right next to each other.

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.google.com/robots.txt")
    if (err != nil) {
        log.Fatal(err)
    }
    defer res.Body.Close()
    robots, err := io.ReadAll(res.Body)
    if(err != nil) {
        log.Fatal(err)
    }
    fmt.Printf("%s", robots)
}
```

---

## Panic

In Go, we don't have exceptions, like how many programming languages have, because a lot of what is considered exceptions in other languages are considered normal in Go. For example: If you try to open a file that doesn't exist, Go returns an error value but doesn't throw an exception.

However, sometimes, a Go application can enter a state where it can no longer continue to run. This can be triggered by the Go runtime but we can also trigger this on our own.

Here are some characteristics of `panic` statements:

1. Can be manually triggered or triggered by the Go runtime.
2. Panics happen after deferred statements are executed.

```go
package main

import "fmt"

func main() {
	fmt.Println("start")					// This is executed first
	defer fmt.Println("This is deferred!")	// This is executed third
	fmt.Println("middle")					// This is executed second
	panic("something bad happened")
	fmt.Println("end")						// doesn't show as application panics before.
}
```

**Using panic manually when running a http server**

```go
package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello Go!"))
    })
    err := http.ListenAndServe(":8080", nil)
    if(err != nil) {
        panic(err.Error())
    }
}
```

---

## Recover

When the application starts to panic, ideally there should be some way to recover.
When panic is called, including implicitly for run-time errors such as indexing a slice out of bounds or failing a type assertion, it immediately stops execution of the current function and begins unwinding the stack of the goroutine, running any deferred functions along the way. If that unwinding reaches the top of the goroutine's stack, the program dies. However, it is possible to use the built-in function recover to regain control of the goroutine and resume normal execution.

A call to recover stops the unwinding and returns the argument passed to panic because the only code that runs while unwinding is inside deferred functions, recover is only useful inside deferred functions.

One application of recover is to shut down a failing goroutine inside a server without killing the other executing goroutines.

- The proper place to use a recover function is inside a deferred function as while panic stops executing any other statements, it will execute deferred functions.

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")					// This is executed first
	panicker()
    fmt.Println("end")                      // This is executed last
}

func panicker() {
    fmt.Println("About to panic")           // This is executed second
    defer func() {
        if err := recover(); err != nil {
            log.Println("Error:", err)		// This is executed third
			// can rethrow the panic
        }
    }()
    panic("Something bad happened")
    fmt.Println("Done panicking")			// This does not execute
}
```

---

# Pointers

Pointers hold a reference in memory about another variable. By default all primitive types are value types.

The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.

## Creating Pointers

Pointers can be created by adding a `*` to the type during declaration (adds a pointer to the data of the type) and assigning another variable with the address-of (`&`) operator.

The `*` operator can also be used to dereference a numerical representation of a memory location to get it's actual value by prefixing a pointer with it.

```go
package main

import "fmt"

func main() {
	var a int = 42
    var b *int = &a // numeric representation for the memory address that is holding the location of a

    fmt.Println(a, b)   // 42 0xc0000a2000
    fmt.Println(&a, b)  // 0xc0000a2000 0xc0000a2000
    fmt.Println(a, *b)  // 42 42	(dereferenced)

	/** Modify all Pointers by changing the original value directly */
	a = 30
    fmt.Println(a, *b)  // 30 30

	/** Changing the deferenced b changes the original value */
	*b = 50
    fmt.Println(a, *b)  // 50 50
}
```

Now, if you come from a background in languages that allow you to work with pointers as variables, then you might be able to do something called Pointer Arithmetic. Go does not allow Math to be done in the same way.

If you've come from C or C++, you are probably aware of the tremendous performance benefits and advantages if you're allowed to do Pointer Arithmetic because you can jump around memory pretty quickly and get performance benefits in certain applications. However, when you're getting into Pointer Arithmetic, you're getting into some fairly complex code and since one of Go's core design concerns is Simplicity, the decision was made to leave Pointer Arithmetic out of the Go language.

```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
    d := &a[1]
	// e := &a[1] - 8							// will throw an error

	fmt.Printf("%v %p %p %p\n", a, b, c, d)	// [1 2 3] 0xc00001a1e0 0xc00001a1e8 0xc00001a1f0
}
```

We can notice that there's difference of 8 between the memory locations.
Since `a` is an array of integers and integers in this version of the runtime are 8 bytes long, each element in the array are 8 bytes apart.

> **Note**: If at all you need to do something like Pointer Arithmetic, we can use the [`unsafe`] (https://pkg.go.dev/unsafe) package for those advanced scenarios. This package contains operations that step around the type safety of Go programs and is not checked by the Go runtime.

So far, we've always declared the underlying type first. But that's often not necessary in Go, because often only want to work with the Pointers and don't really care where the underlying data is stored, you just need the ability to point to it wherever that is.

```go
package main

import "fmt"

type myStruct struct {
    foo int
}

func main() {
	var ms *myStruct
    ms = &myStruct{
        foo: 40,
    }
    fmt.Println(ms)						// &{40}

	/** Alternative declaration */
	ms2 := &myStruct{foo: 40}
    fmt.Println(ms2)                    // &{40}
    fmt.Println(ms2.foo)				// 40
}
```

---

## Working with `nil`

A Pointer that you didn't initialize will be initialized for you and is going to hold the value `nil`. When accepting Pointers as arguments, it is best practice to check if the Pointer is a `nil` pointer or not, because if it is, we have to handle it in a different way. For e.g. we have a Pointer to a struct and we want to drill through and get to a field, but if the Pointer is `nil`, we are going to get a runtime exception.

**Example of a `nil` pointer**:

```go
package main

import "fmt"

type myStruct struct {
    foo int
}

func main() {
	var ms *myStruct
	fmt.Println(ms)						// <nil>
    ms = new(myStruct)					// Initialize a Pointer but without any fields at the same time
    fmt.Println(ms)						// &{0}
}
```

**Working with a pointer that is `nil`**:

```go
package main

import "fmt"

type myStruct struct {
    foo int
}

func main() {
	var ms *myStruct = new(myStruct)
    fmt.Println(ms)                         // &{0}
	(*ms).foo = 40                          // Deferencing operator has a lower precedence than the dot operator, hence we need to ensure we dereference the struct and not the struct.field
    fmt.Println((*ms).foo)					// 40
}
```

Go however has a syntactic sugar that makes it easy to work with pointers. Instead of constantly dereferencing we can do the same operation as above by getting rid of the dereferencing operator completely:

```go
package main

import "fmt"

type myStruct struct {
    foo int
}

func main() {
	var ms *myStruct = new(myStruct)
    fmt.Println(ms)                     // &{0}
	ms.foo = 40
    fmt.Println(ms.foo)					// 40
}
```

This might be freaking out those that are coming from a C/C++ background because the Pointer `ms` doesn't actually have a field `foo` on it. The pointer `ms` is pointing to a `struct` that has a field `foo`. So how does this work? In Go, this is syntactic sugar where the Go compiler understands that we are not actually trying to access the `foo` field on the pointer but implying that we want to access the underlying `struct`.

In the next sections, we will get into the benefits of using Pointers.

---
