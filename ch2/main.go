package main

import "fmt"

const w = 10

func main() {
	test := "this is a \x61 multiline string"

	fmt.Println(test)

	var x int64 = 10
	var y float64 = 30.2

	var sum1 int64 = x + int64(y)
	var sum2 float64 = float64(x) + y

	fmt.Println(sum1, sum2)

	var a int = 3
	var b byte = 6

	var sum3 int = a + int(b)
	var sum4 byte = byte(a) + b

	fmt.Println(sum3, sum4)

	// var d = 10
	// var e int = d
	// var f int64 = int64(d)

	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	// Write a program that declares an integer variable called i with the value 20. Assign i to a floating-point variable named f. Print out i and f.

	var i = 20
	var f = float64(i)

	fmt.Println(i, f)
}

func exercise2() {
	// Write a program that declares a constant called value that can be assigned to both an integer and a floating-point variable. Assign it to an integer called i and a floating-point variable called f. Print out i and f.

	const c = 20

	i := c
	var f float64 = c

	fmt.Println(i, f)
}

func exercise3() {
	// Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64. Assign each variable the maximum legal value for its type; then add 1 to each variable. Print out their values.

	var b byte
	var smallI int32
	var bigI uint64

	b = 255
	smallI = 2147483647
	bigI = 18446744073709551615

	fmt.Println(b+1, smallI+1, bigI+1)
}
