package main

import "fmt"

func main() {
	var nilai32 int32 = 10000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Budi"
	var b byte = name[0]
	var bString = string(b)

	fmt.Println(name)
	fmt.Println(b)
	fmt.Println(bString)
}
