package main

import "fmt"

func main () {
	//membuat array bertipe string
	var name [3]string
	name[0] = "Aldi"
	name[1] = "Baharuddin"
	name[2] = "Muhammad Baharuddin"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int {
		90, 80, 95,
	}

	fmt.Println(values, "ini index ke 1", values[1])
	fmt.Println(len(values), "adalah panjang dari array values")
	values[2] = 75
	fmt.Println("ini index ke 2 setelah diubah", values[2])
}