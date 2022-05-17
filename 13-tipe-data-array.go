package main

import("fmt")

func main() {
	var names [3]string

	names[0] = "Harli"
	names[1] = "Fauzi"
	names[2] = "Ramli"

	fmt.Println(names)

	var angka = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(angka)

	fmt.Println(len(angka))
	angka[2] = 222

	fmt.Println(angka)
}