package main

import "fmt"

func main() {
	var ujian = 90
	var absensi = 80

	// var lulusUjian bool = ujian > 80
	// var lulusAbsensi bool = absensi > 80
	
	// var lulus bool = lulusUjian && lulusAbsensi
	// fmt.Println(lulus)

	//simple way
	fmt.Println(ujian > 80 && absensi > 80)
}