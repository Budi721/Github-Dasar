package main

import "fmt"

func ups(num int) interface{} {
	if num == 1 {
		return 1
	} else if num == 2 {
		return "2"
	} else {
		return true
	}
}

func main() {
	var ups = ups(2)
	fmt.Println(ups)
}
