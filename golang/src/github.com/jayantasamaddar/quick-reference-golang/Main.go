package main

import "fmt"

type CustomMap[T comparable, V int | string] map[T]V

func main() {
	user := make(CustomMap[string, int])
	user["id"] = 123456
	fmt.Println(user)
}
