package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke :", counter)
		counter++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke :", i)
	}

	slice := []string{"Budi", "Yayang", "Kiki", "Anas"}

	for i, _ := range slice {
		fmt.Println(slice[i])
	}
}
