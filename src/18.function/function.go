package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

//with parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

//with return
func getHello(name string) string {
	return "hello " + name
}

//multiple reture
func getFullName() (string, string) {
	return "Aldi", "Mansur"
}

//named return value
// seandainya tipenya string semua cukup deklarasikan tipe nya sekali di akhir contoh :
// (nama, alamat, domisili string)
func getData() (namaUser string, umurUser int8, domisili string) {
	namaUser = "Baharuddin"
	umurUser = 22
	domisili = "Jakarta"
	// penulisan return nya bisa hanya return karena sudah dideklarasikan di parameternya
	// tetapi bisa juga tulis return namaUser, umurUser, Domisili
	return
}

func main() {
	sayHello()
	sayHelloTo("Muhammad", "Baharuddin")

	result := getHello("Aldi")
	fmt.Println(result)

	namaDepan, namaBelakang := getFullName()
	fmt.Println(namaDepan, namaBelakang)

	//menghiraukan return value yang tidak digunakan menggunakan _
	nama, _ := getFullName()
	fmt.Println(nama)

	a, b, c := getData()
	fmt.Println("nama :", a)
	fmt.Println("umur :", b)
	fmt.Println("domisili :", c)
}