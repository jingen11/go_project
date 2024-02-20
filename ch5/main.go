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
	"/": div2,
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
		Person{FirstName: "Pat", LastName: "Patterson", Age: 37},
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

	// if len(os.Args) < 2 {
	// 	log.Fatal("no file specified")
	// }
	// f, err := os.Open(os.Args[0])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// data := make([]byte, 2048)
	// for {
	// 	_, err := f.Read(data)
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.Fatal(err)
	// 		}
	// 		break
	// 	}
	// }

	// f, closer, err := getFile(os.Args[1])

	// defer closer()

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modeMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)

	result, err := div(10, 1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)

	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob")) // should print Hello Bob
	fmt.Println(helloPrefix("Maria"))
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

func div2(i, j int) int {
	// The simple calculator program doesn’t handle one error case: division by zero. Change the function signature for the math operations to return both an int and an error. In the div function, if the divisor is 0, return errors.New("division by zero") for the error. In all other cases, return nil. Adjust the main function to check for this error.

	return i / j
}

func div1(i, j int) (int, error) {
	// The simple calculator program doesn’t handle one error case: division by zero. Change the function signature for the math operations to return both an int and an error. In the div function, if the divisor is 0, return errors.New("division by zero") for the error. In all other cases, return nil. Adjust the main function to check for this error.
	if j == 0 {
		return 0, errors.New("division by zero")
	}

	return i / j, nil
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

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}

func modeMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func div(num int, denom int) (int, error) {
	// The simple calculator program doesn’t handle one error case: division by zero. Change the function signature for the math operations to return both an int and an error. In the div function, if the divisor is 0, return errors.New("division by zero") for the error. In all other cases, return nil. Adjust the main function to check for this error.
	if denom == 0 {
		return 0, errors.New("division by zero")
	}

	return num / denom, nil
}

// exercise2
func fileLen(s string) (int, error) {
	// Write a function called fileLen that has an input parameter of type string and returns an int and an error. The function takes in a filename and returns the number of bytes in the file. If there is an error reading the file, return the error. Use defer to make sure the file is closed properly.

	f, error := os.Open(s)

	if error != nil {
		return 0, error
	}

	b := make([]byte, 2048)

	count, error := f.Read(b)

	if error != nil {
		return 0, error
	}

	defer func() {
		f.Close()
	}()

	return count, nil
}

func prefixer(greet string) func(string) string {
	// Write a function called prefixer that has an input parameter of type string and returns a function that has an input parameter of type string and returns a string. The returned function should prefix its input with the string passed into prefixer. Use the following main function to test prefixer:
	// func main() {
	// 	helloPrefix := prefixer("Hello")
	//  fmt.Println(helloPrefix("Bob")) // should print Hello Bob
	// fmt.Println(helloPrefix("Maria")) // should print Hello Maria
	// 	}

	return func(name string) string {
		return greet + " " + name
	}
}
