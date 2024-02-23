package main

import "fmt"

const c = 10 + 5 // can be evaluated at compile time, so true
// const cf = x + y // fase since value cannot be evaluated at compile time

// var x = 4
// var y = 9

func main() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z)

	var pointerToX2 *int
	pointerToX2 = &x
	fmt.Println(pointerToX2)

	var y = new(int)
	y = &x
	fmt.Println(y)

	f := &Foo{}
	fmt.Println(f)

	// wont compile
	// fa := &string
	// instead
	var fa string
	pfa := &fa

	fmt.Println(fa)
	fmt.Println(&fa)
	fmt.Println(*pfa)
	fa = "c"
	fmt.Println(fa)
	fmt.Println(*pfa)

	var mn = "perry"

	p := Person{
		FirstName:  "pat",
		MiddleName: &mn, // makePointer("perry")
		LastName:   "patterson",
	}

	fmt.Println(p)

	a1 := make([]int, 3, 5)
	p1 := &a1
	fmt.Println(p1)
	fmt.Printf("p1: %p\n", p1)
	fmt.Printf("%p\n", a1)

	a1 = append(a1, 4)
	p2 := &a1
	fmt.Printf("p2: %p\n", p2) // address of slice
	fmt.Printf("%p\n", a1)     // address of first element in the slice

	a1 = append(a1, 5)
	p3 := &a1
	fmt.Println(p3)
	fmt.Println(*p3)
	fmt.Printf("%p\n", a1)

	a1 = append(a1, 6)
	p4 := &a1
	fmt.Println(p4)
	fmt.Printf("%p\n", a1)

	fmt.Println(&a1)

	u := make([]int, 3, 5)
	fmt.Printf("outside: %p\n", &u)
	fmt.Printf("outside first el: %p\n", u)
	updateSlice(u)
	fmt.Println(u)

	u1 := make([]int, 4, 5)
	fmt.Printf("outside: %p\n", &u1)
	fmt.Printf("outside first el: %p\n", u1)
	updateSlice(u1)
	fmt.Println(u1)

	u2 := make([]int, 5)
	fmt.Printf("outside: %p\n", &u2)
	fmt.Printf("outside first el: %p\n", u2)
	updateSlice(u2)
	fmt.Println(u2)

}

func makePointer[T any](t T) *T {
	return &t
}

type Foo struct{}

type Person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func updateSlice(s []int) {
	s = append(s, 1)

	fmt.Printf("inside func: %p\n", &s)
	fmt.Printf("inside first el: %p\n", s)
	s[0] = 10
}
