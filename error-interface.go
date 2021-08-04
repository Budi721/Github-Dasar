package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Tidak bisa dibagi nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagi(10, 0)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Hasil", hasil)
	}
}
