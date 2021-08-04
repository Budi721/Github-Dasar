package main

import (
	"fmt"
)

func random() interface{} {
	return "Ups"
}

func main() {
	result := random()

	switch value := result.(type) {
	case int:
		fmt.Println(value, "is integer")
	case string:
		fmt.Println(value, "is string")
	default:
		fmt.Println("Unknown type")
	}
}
