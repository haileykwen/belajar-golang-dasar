package main;

import "fmt";

func main() {
	var name = "Fauzi";

	if name == "Harli" {
		fmt.Println("Hello Harli");
	} else if name == "Fauzi" {
		fmt.Println("Hello Fauzi");
	} else {
		fmt.Println("Hello, boleh kenalan?");
	};





	// if dengan short statement

	// var length = len(name);

	if length := len(name); length > 5 {
		fmt.Println("Nama lebih dari 5 karakter");
	} else {
		fmt.Println("Nama kurang dari sama dengan 5 karakter");
	};
};