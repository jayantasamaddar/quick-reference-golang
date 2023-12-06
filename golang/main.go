package main

import "github.com/jayantasamaddar/quick-reference-golang/std"

type CustomMap[T comparable, V int | string] map[T]V

func main() {
	// user := make(CustomMap[string, int])
	// user["id"] = 123456
	// fmt.Println(user)

	// String Operations Demo
	// std.StringOperationsDemo()

	// String Conversions Demo
	// std.StringConversionsDemo()

	// Slice Operations Demo
	// std.SliceOperationsDemo()

	// Reflection Operations Demo
	// std.ReflectionOperationsDemo()

	// JSON Encoding and Decoding Demo
	// std.JSONOperationsDemo()

	// Flag Demo
	// std.FlagOperationsDemo()

	// Log Demo
	// std.LogOperationsDemo()

	// HTML Template Demo
	std.HTMLTemplatesDemo()
}
