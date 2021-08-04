package main

import "fmt"

func main() {
	var months = [...]string {
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var sliceMonths = months[4:7]
	fmt.Println(sliceMonths)
	fmt.Println(len(sliceMonths))
	fmt.Println(cap(sliceMonths))

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Libur Lagi")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Budi"
	newSlice[1] = "Rahmawan"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
}
