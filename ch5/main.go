package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div1,
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var myFuncVariable func(string) int = f1

	x := myFuncVariable("1201-31")
	fmt.Println(x)

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "3"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator: ", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}

	peoples := []Person{
		Person{FirstName: "Pat", LastName: "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		Person{"Fred", "Fredson", 18},
	}
	fmt.Println(peoples)

	sort.Slice(peoples, func(i, j int) bool {
		return peoples[i].LastName < peoples[j].LastName
	})
	fmt.Println(peoples)

	sort.Slice(peoples, func(i, j int) bool {
		return peoples[i].Age < peoples[j].Age
	})
	fmt.Println(peoples)
	result, err := div(10, 1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}

func f1(a string) int {
	return len(a)
}

func add(i, j int) int {
	return i + j
}

func sub(i, j int) int {
	return i - j
}

func mul(i, j int) int {
	return i * j
}

func div1(i, j int) int {
	return i / j
}

func anonymousFn() {
	f := func() {
		fmt.Println("I am an anonymous function")
	}

	f()

	//

	func() {
		fmt.Println("Inline anonymous function")
	}()
}

func div(num int, denom int) (int, error) {
	// The simple calculator program doesn’t handle one error case: division by zero. Change the function signature for the math operations to return both an int and an error. In the div function, if the divisor is 0, return errors.New("division by zero") for the error. In all other cases, return nil. Adjust the main function to check for this error.
	if denom == 0 {
		return 0, errors.New("division by zero")
	}

	return num / denom, nil
}
