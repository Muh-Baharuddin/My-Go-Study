package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b

	fmt.Println(c)

	a += 10
	fmt.Println(a)

	a++
	fmt.Println(a)

	var negative = -100
	fmt.Println(negative)

	var name1 = "Aldi"
	var name2 = "Aldi"
	var name3 = "Baharuddin"
	//name3 lebih besar karena B diatas dari A (A, B, C ...)

	var hasil1 bool = name1 == name2
	var hasil2 bool = name1 > name3
	fmt.Println(hasil1)
	fmt.Println(hasil2)
}