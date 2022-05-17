package main

import("fmt")

func main() {
	var book map[string]string = map[string]string{
		"author": "harleykwen",
		"title": "how to bla bla",
	}

	var movie map[string]string = make(map[string]string)
	movie["title"] = "Endgame"
	movie["rating"] = "9"

	fmt.Println(book)
	fmt.Println(movie)
}