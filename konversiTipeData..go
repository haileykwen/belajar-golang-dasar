package main

import ("fmt")

func main() {
	var nilai32 int32 = 1000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // terjadi integer overflow

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name string= "Harli"
	var r = name[2] // byte
	var rString string = string(r)
	
	fmt.Println(name)
	fmt.Println(r)
	fmt.Println(rString)
}