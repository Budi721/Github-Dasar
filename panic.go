package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp2(err bool)  {
	defer endApp()
	if err {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp2(true)
	fmt.Println("Budi")
}
