package main

import "fmt"

func main() {
	person := map[string]string {
		"name": "Budi",
		"age": "20 Tahun",
	}

	person["job"] = "Progammer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])

	var book = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Budi"

	delete(book, "author")
}
