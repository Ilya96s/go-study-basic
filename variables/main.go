package main

import "fmt"

func main() {
	beerDelicious := true
	croutonsDelicious := true

	if beerDelicious || croutonsDelicious {
		fmt.Println("We're going for a walk")
	} else {
		fmt.Println("We're not going out")
	}
}
