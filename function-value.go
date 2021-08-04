package main

import "fmt"

func getGoodBye(name string) string {
	return "Good, Bye " + name
}

func sayHelloFilter(name string, filter func(string)string) string {
	return "Hello " + filter(name)
}

func filterName(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

/**
Anonymous function
 */
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name){
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Budi"))

	fmt.Println(sayHelloFilter("Budi", filterName))
	fmt.Println(sayHelloFilter("Anjing", filterName))

	blacklist := func(name string) bool {
		return name == "root"
	}

	registerUser("root", blacklist)
	registerUser("budi", func(name string) bool {
		return name == "admin"
	})
}
