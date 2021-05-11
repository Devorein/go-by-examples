package examples

import "fmt"

// Person struct with a name and age field
type person struct {
	name string
	age  int
}

// Constructs a new person and returns a reference to the created person
func newPerson(name string, age int) *person {
	p := person{name: name, age: age}
	return &p
}

func Struct() {
	fmt.Println(person{"Bob", 19})
	fmt.Println(person{name: "John", age: 19})
	fmt.Println(person{name: "Jane"})
	fmt.Println(person{age: 20})
	fmt.Println(&person{name: "Ann", age: 20})

	fmt.Println(newPerson("Jon", 29))

	s1 := newPerson("Alan", 10)

	// Struct are stored in reference
	s2 := s1

	s2.age = 19
	fmt.Println(s1.age)
}
