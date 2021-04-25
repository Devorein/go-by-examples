package main

import (
	"fmt"
	"reflect"
	"sync"
)

var wg = sync.WaitGroup{}

func maps_struct() {
	type Human struct {
		weight float32
		height float32
		age    int
	}

	type Fighter struct {
		Human
		division string `required max:"100"`
		wins     int
		losses   int
	}

	gsp := Fighter{
		wins:     20,
		losses:   2,
		division: "Welterweight",
	}
	gsp.height = 5.11
	gsp.weight = 176
	gsp.age = 36

	khabib := Fighter{
		Human:    Human{weight: 155, height: 5.10, age: 30},
		wins:     29,
		losses:   0,
		division: "Lightweight",
	}
	fmt.Println(gsp)
	fmt.Println(khabib)
	t := reflect.TypeOf(Fighter{})
	field, _ := t.FieldByName("division")
	fmt.Println(field.Tag)
}

func if_else_statements_example1() {
	type mma_record struct {
		wins   uint8
		losses uint8
		nc     uint8
	}

	ufcFightersRecord := map[string]mma_record{
		"Kamaru Usman":        mma_record{wins: 19, losses: 1, nc: 0},
		"Khabib Nurmagomedov": mma_record{wins: 29, losses: 0, nc: 0},
		"Isreal Adesanya":     mma_record{wins: 20, losses: 1, nc: 0},
	}

	if mma_record, ok := ufcFightersRecord["Kamaru Usman"]; ok {
		fmt.Println(mma_record.wins)
	}
}

func if_else_statements_example2() {
	number := 20
	guess := 21
	if guess < 1 {
		fmt.Println("Guess must be between 1 and 100")
	} else if guess > 100 {
		fmt.Println("Guess is greater than half of number")
	} else {
		if guess < number {
			fmt.Println("guess is too low")
		}
		if guess > number {
			fmt.Println("guess is too high")
		}
		if guess == number {
			fmt.Println("Correct guess")
		}
	}
}

func switch_statements() {
	switch i := 3 + 2; i {
	case 1, 3, 5:
		fmt.Println("odd")
	case 2, 4, 6:
		fmt.Println("even")
	default:
		fmt.Println("complex")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i >= 11 && i <= 20:
		fmt.Println("i is greater than 10 and less than 20")
	default:
		fmt.Println("i is greater than 20")
	}

	var s interface{} = [2]int{1, 2}
	switch s.(type) {
	case int:
		fmt.Println("s is an int")
	case float32:
		fmt.Println("s is an float32")
	case float64:
		fmt.Println("s is an float64")
	case [2]int:
		fmt.Println("s is 2 length int array")
	}
}

func loops() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i <= 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	i := 0
	for ; i <= 5; i++ {
		fmt.Println(i)
	}

	j := 0
	for {
		j++

		if j == 2 {
			continue
		}

		if j == 3 {
			break
		}
		fmt.Println(j)
	}

	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Println(i, j)
		}
	}

	// Breakout with label
Loop:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Println(i, j)
			if j == 3 {
				break Loop
			}
		}
	}
	fmt.Println("Breakout of nested loop")

	// Range keyword
	divisions := []string{"FEW", "BMW", "FLW", "LGW", "WW", "MW", "LHW", "HW"}
	for k, v := range divisions {
		fmt.Println(k, v)
	}

	// Looping over string
	for k, v := range "Hello" {
		fmt.Println(k, v)
	}

	// Looping over map
	for k, v := range map[string]string{"a": "A", "b": "B"} {
		fmt.Println(k, v)
	}
}

func defer_examples() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Three")

	a := "a"
	defer fmt.Println(a)
	a = "b"

	b := map[int]string{1: "One", 2: "Two", 3: "Three"}
	defer fmt.Println(b)
	b[1] = "Two"
}

func panic_examples() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error", err)
		}
	}()
	panic("Panicked")
}

func pointer_examples() {
	var a int = 32
	var b *int = &a
	fmt.Println(a, &a, b, *b)
	a = 25
	fmt.Println(*b, a)
	*b = 24
	fmt.Println(*b, a)

	i := []int{1, 2, 3}
	j := &i[0]
	k := &i[1]
	fmt.Println(i, j, k, *j, *k)

	type my_struct struct {
		a int
	}

	var x *my_struct = &my_struct{a: 1}
	x.a = 2
	(*x).a = 3
	fmt.Println(x, *x, (*x).a, x.a)
}

func pass_by_value(text string) {
	text = "New Value"
}

func pass_by_pointer(text *string) {
	*text = "New Value"
}

func sum(msg string, values ...int) {
	var sum int = 0
	for _, v := range values {
		sum += v
	}
	fmt.Println(msg, sum)
}

func return_pointer_sum(values ...int) *int {
	var sum int = 0
	for _, v := range values {
		sum += v
	}
	return &sum
}

func return_divide(a float32, b float32) (float32, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func function_examples() {
	// Pass by value example
	a := "Hello"
	pass_by_value(a)
	fmt.Println(a)

	// Pass by pointers example
	b := "Hello"
	pass_by_pointer(&b)
	fmt.Println(b)

	sum("The sum is", 1, 2, 3, 4)

	s := return_pointer_sum(1, 2, 3, 4, 5)
	fmt.Println(*s)

	r, err := return_divide(5.0, 1.0)
	if err != nil {
		fmt.Println("error occurred", err)
		return
	}
	fmt.Println(r)

	// anonymous function
	d := func() {
		fmt.Println("Anonymous function")
	}
	d()

	// Anonymous function within a loop
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	divide()
}

// An example of function declaration in a variable
func divide() {
	var d func(a, b float64) (float64, error) = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Divide by zero error")
		}
		return a / b, nil
	}

	v, err := d(1, 2)
	if err != nil {
		fmt.Println("An error occurred", err)
		return
	}
	fmt.Println(v)

	g := greeter{
		greeting: "Hello there",
		name:     "Safwan",
	}
	g.greet()
	fmt.Println(g.name)
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "new name"
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	return fmt.Println(string(data))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	(*ic)++
	return int(*ic)
}

func interface_examples() {
	// struct interface example
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World"))

	// int interface example
	ic := IntCounter(0)
	var inc Incrementer = &ic
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

func goroutines_examples_1() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "World"
	wg.Wait()
}

var counter = 0
var wg1 = sync.WaitGroup{}
var m = sync.RWMutex{}

func goroutines_examples_2() {
	for i := 0; i < 10; i++ {
		wg1.Add(2)
		m.RLock()
		go printer()
		m.Lock()
		go incrementer()
	}

	wg1.Wait()
}

func printer() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg1.Done()
}

func incrementer() {
	counter++
	m.Unlock()
	wg1.Done()
}

var wg3 = sync.WaitGroup{}

func channels_examples() {
	ch := make(chan int)
	wg3.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg3.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 22
		close(ch)
		wg3.Done()
	}(ch)

	wg3.Wait()
}

func main() {
	// if_else_statements_example1()
	// if_else_statements_example2()
	// switch_statements()
	// loops()
	// maps_struct()
	// defer_examples()
	// panic_examples()
	// pointer_examples()
	// function_examples()
	// interface_examples()
	// goroutines_examples_1()
	// goroutines_examples_2()
	channels_examples()
}
