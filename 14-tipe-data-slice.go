package main

import("fmt")

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"Oktober",
		"November",
		"Desember",
	}

	var sliceMonths = months[4:7]
	fmt.Println(months)
	fmt.Println(sliceMonths)

	// months[0] = "Months Diubah"
	// fmt.Println(months)

	// sliceMonths[0] = "Slice Months Diubah"
	// fmt.Println(sliceMonths)

	fmt.Println(len(sliceMonths))
	fmt.Println(cap(sliceMonths))
}