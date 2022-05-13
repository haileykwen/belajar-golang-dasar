package main

import ("fmt")

func main() {
	const greeting = "Hello World!"

	const (
		firstname = "Harli"
		lastame = "Fauzi Ramli"
	)

	fmt.Println(greeting)
	fmt.Println(firstname + " " + lastame)
}