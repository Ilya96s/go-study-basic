package main

import "fmt"

func main() {
	square(2)
}
func square(x int) {
	fmt.Println("x:", x)
	fmt.Println("x в квадрате:", x*x)
}
