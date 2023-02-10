package main

import "fmt"

func main() {
	//disarankan menggunakan [...] jika [] akan dianggap slice. jika panjang dari array tidak diketahui
	var months = [...]string{
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

	//append slice
	days := [...]string{"senin", "selesa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur baru")
	daySlice2[0] = "Ups"
	daySlice2[1] = "bukan minggu"
	//dikarenakan kapasitas dari array sudah penuh maka slice akan membuat array baru
	//sehingga nilai dari Ups dan bukan minggu tidak ada terbaca di days dan daySlice1 yang menggunakan array pertama
	fmt.Println(daySlice2)
	fmt.Println(daySlice1)
	fmt.Println(days)

	//make slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Aldi"
	newSlice[1] = "Baharuddin"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}