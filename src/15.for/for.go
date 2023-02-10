package main

import "fmt"

func main() {
	counter := 1

	for counter <= 5 {
		fmt.Println("Looping ke - ", counter)
		counter++
	}

	fmt.Println("")

	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Looping ke - ", counter)
	}

	slice := []string{"Aldi", "Baharuddin", "Mansur"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range
	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}
	//gunakan _ agar golang mengetahui bahwa variable tersebut tidak digunakan
	for _, value := range slice {
		fmt.Println(value)
	}

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Someone Famous"

	for key, value := range book {
		fmt.Println(key, "=", value)
	}
}