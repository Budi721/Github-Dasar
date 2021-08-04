package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp(value int) {
	defer logging()

	fmt.Println("Running app")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApp(10)
}
