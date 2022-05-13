package main

import ("fmt")

func main() {
	type noKTP string
	type married bool

	var noKTPHarli noKTP = "321411"
	var marriedHarli married = true

	fmt.Println(noKTPHarli)
	fmt.Println(marriedHarli)
}