package main

import "fmt"

func main() {
	var name string

	name = "Budi Rahmawan"
	fmt.Println(name)

	name = "Budi Wawan"
	fmt.Println(name)

	var (
		firstName = "Budi"
		lastName  = "Rahmawan"
	)
	fmt.Println(firstName, lastName)

	var age = 20
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)
}
