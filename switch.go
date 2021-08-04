package main

import "fmt"

func main() {
	name := "Budi"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hello Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	/**
	Pertanyaan golang apakah penggunaan switch tanpa variable
	dianjurkan atau dapat mempengaruhi peforma aplikasi
	 */
	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("Nama terlalu panjang")
	case length2 > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
