package main

import "fmt"

func main() {
	var name = "Aldi"

	if name == "Aldi" {
		fmt.Println("Hallo ", name)
	} else if name == "Baharuddin" {
		fmt.Println("Hallo Baharuddin")
	} else {
		fmt.Println("Hallo, Boleh kenalan?")
	}

	//if short statement

	// var length = len(name)
	if length := len(name); length > 5 {
		fmt.Println("Namanya panjang yah")
	} else {
		fmt.Println("Namanya singkat yah")
	}
}