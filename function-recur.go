package main

import (
	"fmt"
)

func factorialLoop(num int) int {
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecur(num int) int {
	if (num == 1){
		return 1
	} else {
		return num * factorialRecur(num-1)
	}
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)

	fmt.Println(factorialRecur(5))
}
