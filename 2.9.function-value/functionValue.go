package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Aldi"))

	result := goodbye("Baharuddin")
	fmt.Println(result)
}