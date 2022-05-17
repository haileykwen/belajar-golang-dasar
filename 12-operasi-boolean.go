package main

import("fmt")

func main() {
	var nilaiAkhir = 90
	var absesnsi = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = absesnsi > 70
	var lulus = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}