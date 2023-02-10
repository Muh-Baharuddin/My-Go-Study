package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string
	person := map[string]string {
		"name": "Aldi",
		"address": "Jakarta", 
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Someone Famous"
	book["wrong"] = "Oops"

	fmt.Println(book)
	delete(book, "wrong")

	fmt.Println(book)
}