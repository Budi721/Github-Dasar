package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello ", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Budi", "Rahmawan"
}

func getFullName2() (firstName, lastName string) {
	firstName = "Budi"
	lastName = "Rahmawan"
	return
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	sayHelloTo("Budi", "Rahmawan")

	result := getHello("Budi")
	fmt.Println(result)

	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	firstName2, lastName2 := getFullName2()
	fmt.Println(firstName2)
	fmt.Println(lastName2)

	total := sumAll(10, 10, 29, 21)
	fmt.Println(total)

	slice := []int{29, 23, 4, 2}
	total2 := sumAll(slice...)
	fmt.Println(total2)
}
