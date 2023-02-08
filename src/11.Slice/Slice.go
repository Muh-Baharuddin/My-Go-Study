package main

import "fmt"

func main() {
	//disarankan menggunakan [...] jika panjang dari array tidak diketahui
	var months = []string{
		"Januari", "Februari", "Maret",
		"April", "Mei", "Juni", "Juli",
		"Agustus", "September", "Oktober", "Novermber", "Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) //len untuk panjang
	fmt.Println(cap(slice1)) //cap untuk kapasitas

	//jika nilai dari array berubah maka nilai dari slice juga ikut berubah begitu juga sebaliknya
	months[5] = "bulannn"
	fmt.Println(slice1)

	slice1[0] = "Meiii"
	fmt.Println(months)
}