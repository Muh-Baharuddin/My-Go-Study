package main

import "fmt"

func main() {
	var name string

	name = "Aldi"
	fmt.Println(name)

	name = "Muh. Baharuddin"
	fmt.Println(name)

	var friendName = "Zakiul"
	fmt.Println(friendName)

	var age int8 = 30
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	country = "Jepang"
	fmt.Println(country)

	var (
		firstName = "Muhammad"
		lastName = "Baharuddin"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

	//const
	const namaDepan string = "Muhammad const"
	const namaBelakang = "Baharuddin const"
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)

	const (
		namaTeman1 string = "Iqsyal"
		namaTeman2 = "Zakiul"
	)
	fmt.Println(namaTeman1)
	fmt.Println(namaTeman2)
}
