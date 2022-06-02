package main;

import "fmt";

func main() {
	counter := 1;
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter);
		counter++;
	};





	// for loops dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan dengan statement ke", counter);
	};





	// for loops range
	// slice
	names := []string{"Harli", "Fauzi", "Ramli"};
	for index, name := range names {
		fmt.Println("index", index, "=", name);
	};

	// map
	person := make(map[string]string);
	person["name"] = "Harli";
	person["gender"] = "male";
	for key, value := range person {
		fmt.Println("key", key, "=", value);
	};

	// if index is not use
	for _, name := range names {
		fmt.Println(name);
	};
	for _, value := range person {
		fmt.Println(value);
	};
};