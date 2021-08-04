package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	budi := newMap("Budi")
	if budi == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(budi)
	}
}
