package main

import (
	"fmt"
	"reflect"
)

func maps_structs() {
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

func main() {
	if_else_statements()
	// maps_structs()
}
