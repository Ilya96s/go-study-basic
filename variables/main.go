package main

import "fmt"

func main() {
	x := 25
	y := "string"

	fmt.Println("x до:", x)
	fmt.Println("y до:", y)

	foo(x, y)

	fmt.Println("x после:", x)
	fmt.Println("y после:", y)
}

// Значения буду изменены только в рамках метода, потому что в Go параметры метода передаются по значению,
// это значит что при передаче переменной создается ее копия
func foo(x int, y string) {
	x = 10
	y = "new string"
}
