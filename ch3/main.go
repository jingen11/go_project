package main

import (
	"fmt"
)

func main() {
	// // slice literal
	// y := []int{1, 2, 3}
	// // nil zero value
	// var x []int
	// // make
	// z := make([]int, 5)

	// fmt.Println(x, y, z)
	// fmt.Println(x == nil)

	// play with capacity
	// x := make([]string, 0, 5)
	// x = append(x, "a", "b", "c", "d")
	// y := x[:2:2]
	// z := x[2:4:4]
	// fmt.Println(cap(x), cap(y), cap(z))
	// y = append(y, "i", "j", "k")
	// x = append(x, "x")
	// z = append(z, "y")
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	x := []int{1, 2}
	y := make([]int, 0)
	y = append(y, 2)
	y = append(y, 4)

	num := copy(y, x)
	fmt.Println(x, y, num)

	var m = make(map[string]int, 10)
	fmt.Println(m == nil)
	var a = make([]int, 0)
	fmt.Println(a)

	type firstPerson struct {
		name string
		age  int
	}

	type secondPerson struct {
		name string
		age  int
	}

	type thirdPerson struct {
		namee string
		age   int
	}

	fp := firstPerson{}
	sp := secondPerson{}
	// tp := thirdPerson{}

	fp = firstPerson(sp)
	// fp = firstPerson(tp)
	fmt.Println(fp)
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	// Write a program that defines a variable named greetings of type slice of strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт". Create a subslice containing the first two values; a second subslice with the second, third, and fourth values; and a third subslice with the fourth and fifth values. Print out all four slices.
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	first := greetings[0:2:2]
	second := make([]string, 3)
	copy(second, greetings[1:4])
	third := greetings[3:]
	fmt.Println(greetings, first, second, third)
}

func exercise2() {
	// Write a program that defines a string variable called message with the value "Hi 🧒🏽 and 👦🏽" and prints the fourth rune in it as a character, not a number.
	var message = "Hi 🧒 and 🧒"

	var runes = []rune(message)
	var s = string(runes[3])

	fmt.Println(s)
}

func exercise3() {
	// Write a program that defines a struct called Employee with three fields: firstName, lastName, and id. The first two fields are of type string, and the last field (id) is of type int. Create three instances of this struct using whatever values you’d like. Initialize the first one using the struct literal style without names, the second using the struct literal style with names, and the third with a var declaration. Use dot notation to populate the fields in the third struct. Print out all three structs.

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	e1 := Employee{
		id: 1,
	}

	e2 := Employee{
		firstName: "Jing En",
		lastName:  "Cheam",
	}

	var e3 = Employee{}

	e3.firstName = "Chia Wei"
	e3.lastName = "Tan"
	e3.id = 3

	fmt.Println(e1, e2, e3)
}
