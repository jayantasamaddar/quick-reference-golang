package main

import "fmt"

type myStruct struct {
    foo int
}

func main() {
	ms := &myStruct{foo: 40}
    fmt.Println(ms)                     // &{40}
    fmt.Println(ms.foo)					// 40
}