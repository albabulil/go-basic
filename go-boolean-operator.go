package main

import "fmt"

func main() {

	var ujian = 80
	var absen = 75

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absen >= 80

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	//OR
	fmt.Println(ujian >= 80 && absen >= 80)
}
