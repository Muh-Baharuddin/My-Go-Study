package main

import "fmt"

//parameternya bisa lebih dari 1 tapi yang jadi slice hanya bisa 1 saja contoh :
//(nama string, numbers ...int) int
func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 5, 3, 20, 2)
	fmt.Println(total)

	slice := []int{10, 20, 30, 40}
	total = sumAll(slice...)
	// ... di belakang berguna untuk memberitahu bahwa itu adalah slice
	fmt.Println(total)
}