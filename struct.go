package main

import "fmt"

type Customer struct {
	Name, Address string
	Age 		  int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello,", name, "my name is", customer.Name)
}

func main() {
	var budi Customer
	budi.Name = "Budi"
	budi.Address = "Pekalongan"
	budi.Age = 20

	budi.sayHi("Eko")
	fmt.Println(budi)

	var joko = Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     39,
	}
	fmt.Println(joko)
}
