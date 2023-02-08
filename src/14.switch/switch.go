package main

import "fmt"

func main() {
	name := "Aldi"

	switch name {
	case "Aldi":
		fmt.Println("Hallo ", name)
		fmt.Println("senang bisa berkenalan dengan mu")
	case "Baharuddin":
		fmt.Println("Hallo ", name)
		fmt.Println("senang bisa berkenalan dengan mu")
	default:
		fmt.Println("Hallo, boleh kenalan")
	}

	//switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Namanya panjang yah")
	case false:
		fmt.Println("Namanya singkat yah")
	}

	//switch without condition
	panjang := len(name)
	switch {
	case panjang > 10:
		fmt.Println("Namanya terlalu panjang yah")
	case panjang > 5:
		fmt.Println("Namanya panjang yah")
	default:
		fmt.Println("Namanya cukup singkat yah")
	}
}