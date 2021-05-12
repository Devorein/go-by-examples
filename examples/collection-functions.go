package examples

import (
	"fmt"
	"strings"
)

func Index(s []string, t string) int {
	for i, v := range s {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(s []string, t string) bool {
	return Index(s, t) >= 0
}

func Any(s []string, fn func(s string) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}

func All(s []string, fn func(s string) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

func Filter(s []string, fn func(s string) bool) []string {
	filtered := make([]string, 0)

	for _, v := range s {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Map(s []string, fn func(s string) string) []string {
	mapped := make([]string, len(s))

	for _, v := range s {
		mapped = append(mapped, fn(v))
	}
	return mapped
}

func CollectionFunctions() {
	var fruits = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(fruits, "apple"))
	fmt.Println(Index(fruits, "mango"))

	fmt.Println(Include(fruits, "apple"))
	fmt.Println(Include(fruits, "mango"))

	fmt.Println(Any(fruits, func(s string) bool {
		return string(s[0]) == "a"
	}))

	fmt.Println(Any(fruits, func(s string) bool {
		return string(s[0]) == "b"
	}))

	fmt.Println(All(fruits, func(s string) bool {
		return strings.Contains(s, "p")
	}))

	fmt.Println(All(fruits, func(s string) bool {
		return string(s[0]) == "b"
	}))

	fmt.Println(Filter(fruits, func(s string) bool {
		return string(s[0]) == "p"
	}))

	fmt.Println(Map(fruits, strings.ToUpper))
}
