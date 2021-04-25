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

func main() {
	// if_else_statements_example1()
	// if_else_statements_example2()
	// switch_statements()
	// loops()
	// maps_structs()
}
