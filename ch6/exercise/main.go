package main

func main() {
	// MakePerson("a", "b", 1)
	// MakePersonPointer("a", "b", 1)
	CreatePersons()
	CreatePersonsA()
}

// 6.1
// Create a struct named Person with three fields: FirstName and LastName of type string and Age of type int. Write a function called MakePerson that takes in firstName, lastName, and age and returns a Person. Write a second function MakePersonPointer that takes in firstName, lastName, and age and returns a *Person. Call both from main. Compile your program with go build -gcflags="-m". This both compiles your code and prints out which values escape to the heap. Are you surprised about what escapes?

// func MakePerson(f string, l string, a int) Person {
// 	return Person{f, l, a}
// }

// func MakePersonPointer(f string, l string, a int) *Person {
// 	return &Person{f, l, a}
// }

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// 6.2
// Write two functions. The UpdateSlice function takes in a []string and a string. It sets the last position in the passed-in slice to the passed-in string. At the end of UpdateSlice, print the slice after making the change. The GrowSlice function also takes in a []string and a string. It appends the string onto the slice. At the end of GrowSlice, print the slice after making the change. Call these functions from main. Print out the slice before each function is called and after each function is called. Do you understand why some changes are visible in main and why some changes are not?

// func UpdateSlice(s []string, a string) {
// 	s[len(s)-1] = a
// 	fmt.Println(s)
// }

// func GrowSlice(s []string, a string) {
// 	s = append(s, a)
// 	fmt.Println(s)
// }

// 6.3
// Write a program that builds a []Person with 10,000,000 entries (they could all be the same names and ages). See how long it takes to run. Change the value of GOGC and see how that affects the time it takes for the program to complete. Set the environment variable GODEBUG=gctrace=1 to see when garbage collections happen and see how changing GOGC changes the number of garbage collections. What happens if you create the slice with a capacity of 10,000,000?
// Pre-allocating the slice makes +GOGC=50+ faster than +GOGC=off+ when the slice grows over time. We have, at most, one GC cycle per run. If you know you are going to need a block of memory, it's best to allocate it at once and use it. If you can re-use it, all the better. That's the reason for the slice buffer pattern, also.
func CreatePersons() {
	p := make([]Person, 0, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		p = append(p, Person{"a", "a", 1})
	}
}

func CreatePersonsA() {
	p := []Person{}

	for i := 0; i < 1_000_000; i++ {
		p = append(p, Person{"a", "a", 1})
	}
}
