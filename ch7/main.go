package main

import (
	"fmt"
	"time"
)

func main() {
	p := Person{}
	fmt.Println(p.String())
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))

	// method value
	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	// method expression
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15))
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type IntTree struct {
	val   int
	left  *IntTree
	right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}
func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}
