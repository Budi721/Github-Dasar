package main

import "fmt"

func main() {
	type noKTP string
	type married bool

	var noKtpBudi noKTP = "3326160702010001"
	var marriedStatus married = true

	fmt.Println(noKtpBudi)
	fmt.Println(marriedStatus)
}
