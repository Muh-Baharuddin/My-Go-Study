package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpAldi NoKTP = "12312512312524"
	fmt.Println(noKtpAldi)

	var marriedStatus Married = true
	fmt.Println(marriedStatus)

}