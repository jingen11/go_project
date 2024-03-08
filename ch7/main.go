package main

import (
	"errors"
	"fmt"
	"net/http"
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

	// embedded field promotion
	m := Manager{
		Employee{"name", "12345"},
		[]Employee{},
	}

	fmt.Println(m.Name)
	fmt.Println(m.ID)

	o := Outer{
		Inner: Inner{
			A: 10},
		S: "Hello"}
	fmt.Println(o.Double()) // "Inner: 20"

	// var myStringer fmt.Stringer
	// var myIncrementer Incrementer
	// pointerCounter := &Counter{}
	// valueCounter := Counter{}

	// myStringer = pointerCounter
	// myStringer = valueCounter
	// myIncrementer = pointerCounter
	//  myIncrementer = valueCounter // compilation error

	var di DoubleInt = 10
	var di2 DoubleInt = 10
	var dis = DoubleIntSlice{1, 2, 3}
	var dis2 = DoubleIntSlice{1, 2, 3}

	DoubleCompare(&di, &di2)
	DoubleCompare(di, dis)
	DoubleCompare(dis, dis2)

	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
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

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

// Composition
type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}
func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

type Incrementer interface {
	Increment()
	Decrement()
}

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	return ""
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) program(data string) {
	c.L.Process(data)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

type Doubler interface {
	Double()
}

type DoubleInt int

func (d DoubleInt) Double() {
	d = d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

func DoubleCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UsernameForId(userId string) (string, bool) {
	username, ok := sds.userData[userId]

	return username, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Cheam",
			"2": "Jing",
			"3": "En",
		},
	}
}

type DataStore interface {
	UsernameForId(string) (string, bool)
}

type Logger interface {
	Log(string)
}

type LoggerAdapter func(string)

func (l LoggerAdapter) Log(message string) {
	l(message)
}

type SimpleLogic struct {
	d DataStore
	l Logger
}

func (sl SimpleLogic) SayHello(userId string) (string, error) {
	sl.l.Log("In SayHello for " + userId)

	name, ok := sl.d.UsernameForId(userId)

	if !ok {
		return "", errors.New("unknown user")
	}

	return "Hello " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userId string) (string, error) {
	sl.l.Log("In SayGoodbye for " + userId)

	name, ok := sl.d.UsernameForId(userId)

	if !ok {
		return "", errors.New("unknown user")
	}

	return "Good Bye " + name, nil

}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l: l,
		d: ds,
	}
}

type Logic1 interface {
	SayHello(userId string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic1
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userId := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic1) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}
