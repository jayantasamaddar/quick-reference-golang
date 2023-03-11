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
