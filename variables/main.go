package main

import "fmt"

func main() {

	x := 25
	y := "string"

	fmt.Println("Результат вызова функции foo(): ", foo(x))
	i, s := bar(x, y)
	fmt.Println("Результат вызова функции bar():", i, s)
}
func foo(x int) int {
	return x * x
}

func bar(x int, y string) (int, string) {
	return x, y
}
