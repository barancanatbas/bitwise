package main

import "fmt"

func main() {
	printBitByNumber()
	andOperation()
	orOperation()
	xorOperation()
	leftShiftOperation()
	rightShiftOperation()
}

func printBitByNumber() {
	a := 4

	fmt.Printf("number: %d, binary: %b\n", a, a)
}

func andOperation() {
	a := 12
	b := 10

	result := a & b

	fmt.Printf("a: %b, b: %b, a & b: %b, result: %d\n", a, b, a&b, result)
}

func orOperation() {
	a := 12
	b := 10

	result := a | b

	fmt.Printf("a: %b, b: %b, a | b: %b, result: %d\n", a, b, a|b, result)
}

// xorOperation is a function to show the xor operation between two numbers
func xorOperation() {
	a := 12
	b := 10

	result := a ^ b

	fmt.Printf("a: %b, b: %b, a ^ b: %b, result: %d\n", a, b, a^b, result)
}

// leftShiftOperation is a shifting bit to the left way
// 13 -> 1101
// 13 << 1 -> 11010
// 13 << 2 -> 110100
// 13 << 3 -> 1101000
func leftShiftOperation() {
	a := 13

	result := a << 1

	fmt.Printf("left shift -> a: %b, a << 1: %b, result: %d\n", a, a<<1, result)
}

// rightShiftOperation is a shifting bit to the right way
// 13 -> 1101
// 13 >> 1 -> 110
// 13 >> 2 -> 11
// 13 >> 3 -> 1
func rightShiftOperation() {
	a := 13

	result := a >> 1

	fmt.Printf("right shift -> a: %b, a >> 1: %b, result: %d\n", a, a>>1, result)
}
