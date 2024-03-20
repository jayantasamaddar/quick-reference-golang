package main

import (
	"fmt"

	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
)

func reverseArray(a []int32) []int32 {
	reverse := []int32{}
	for i := len(a) - 1; i >= 0; i-- {
		reverse = append(reverse, a[i])
	}
	return reverse
}

func reverseArrayInPlace(a []int32) []int32 {
	for i := 0; i < len(a)/2; i++ {
		temp := a[i]
		a[i] = a[len(a)-1-i]
		a[len(a)-1-i] = temp
	}
	return a
}

func main() {

	// 231: Power of Two: Given an integer n, return true if it is a power of two. Otherwise, return false.
	std.PrintHeader("1979: Power of Two: Given an integer n, return true if it is a power of two. Otherwise, return false")
	fmt.Println(std.PrintC(std.Green, fmt.Sprintf("%v", isPowerOfTwo(1))))  // true
	fmt.Println(std.PrintC(std.Green, fmt.Sprintf("%v", isPowerOfTwo(16)))) // true
	fmt.Println(std.PrintC(std.Green, fmt.Sprintf("%v", isPowerOfTwo(3))))  // false

	// 1979: Greatest Common Divisor of a list of integers
	std.PrintHeader("1979: Greatest Common Divisor of a list of integers")
	fmt.Println("The greatest common divisor of [7, 5, 6, 8, 3] is:", std.PrintC(std.Green, fmt.Sprintf("%d", findGCD([]int{2, 5, 6, 9, 10})))) // 2
	fmt.Println("The greatest common divisor of [7, 5, 6, 8, 3] is:", std.PrintC(std.Green, fmt.Sprintf("%d", findGCD([]int{7, 5, 6, 8, 3}))))  // 1
	fmt.Println("The greatest common divisor of [3, 3] is:", std.PrintC(std.Green, fmt.Sprintf("%d", findGCD([]int{3, 3}))))                    // 3

	arr := []int32{1, 4, 3, 2}

	fmt.Println(reverseArray(arr))
	fmt.Println(reverseArrayInPlace(arr))
	fmt.Println(arr)
}
