package main

import ("fmt")

func main() {
	var a = 100
	var b = 50

	// operasi matematika
	var tambah = a + b
	var kurang = a - b
	var bagi = a / b
	var kali = a * b
	var sisa = a % b

	fmt.Println(tambah)
	fmt.Println(kurang)
	fmt.Println(bagi)
	fmt.Println(kali)
	fmt.Println(sisa)

	// augmented operation
	// a += a equal a = a + a
	// a -= a equal a = a - a
	// a *= a equal a = a * a
	// a /= a equal a = a / a
	// a %= a equal a = a % a

	// unary operation
	// ++	ex a++
	// --	ex a--
	// +	ex +a
	// -	ex -a
	// !	ex !false	equal true
}