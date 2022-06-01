package main;

import "fmt";

func main() {
	var name = "Harli";

	switch name {
		case "Harli":
			fmt.Println("Hello Harli");
		case "Fauzi":
			fmt.Println("Hello Fauzi");
		default:
			fmt.Println("Hello, boleh kenalan?");
	};





	// swith dengan short statement

	// var length = len(name);

	switch length := len(name); length > 5 {
		case true :
			fmt.Println("Nama lebih dari 5 karakter");
		case false:
			fmt.Println("Name kurang dari sama dengan 5 karakter");
	};





	// switch tanpa kondisi

	var anotherName = "Harley"

	switch {
		case anotherName == "Harley":
			fmt.Println("Hello Harley");
		case anotherName == "Kwen":
			fmt.Println("Hello Kwen");
		default:
			fmt.Println("Hello, boleh kenalan?");
	};
};