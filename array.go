package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Budi"
	names[1] = "Rahmawan"
	names[2] = "Wawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int { 3, 4, 5 }
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))
}
