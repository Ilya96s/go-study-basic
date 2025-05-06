package main

import "fmt"

func main() {
	age := 25

	if age >= 30 {
		fmt.Println("Age is greater than 30")
	} else if age <= 12 {
		fmt.Println("Age is less than 12")
	} else {
		fmt.Println("Age is over 12 but under 30")
	}
}
