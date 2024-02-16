package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 10
	if x > 5 {
		x := 12
		fmt.Println(x)
		x, y := 5, 6
		fmt.Println(x, y)
	}

	fmt.Println(x)

	if n := rand.Intn(10); n == 0 {
	} else if n > 5 {
	} else {
	}

	exercise1()
	exercise2()

	// part of exercise3
	//Start a new program. In main, declare an int variable called total. Write a for loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive). The body of the for loop should be as follows:
	// total := total + i
	// fmt.Println(total)
	// After the for loop, print out the value of total. What is printed out? What is the likely bug in this code?
	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total) // 0, 1, 2, 3
	}

	fmt.Println(total) // 0

	// because of variable shadowing
}

func exercise1() []int {
	// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
	res := make([]int, 100)

	for i := 0; i < 100; i++ {
		res[i] = rand.Intn(100)
	}

	return res
}

func exercise2() {
	// Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
	// a. If the value is divisible by 2, print “Two!”
	// b. If the value is divisible by 3, print “Three!”
	// c. IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else. d. Otherwise, print “Never mind”.
	res := exercise1()

	for _, v := range res {
		if v%2 == 0 && v%3 == 0 {
			fmt.Println("Six!")
		} else if v%2 == 0 {
			fmt.Println("Two!")
		} else if v%3 == 0 {
			fmt.Println("Three!")
		} else {
			fmt.Println("Never mind.")
		}
	}
}
