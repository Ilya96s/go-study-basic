package main

import "fmt"

func main() {
	number := 5
	numberPointer := &number

	foo(numberPointer)
}

func foo(ptr *int) {
	fmt.Println("Указатель", ptr)
	fmt.Println("Значение указателя:", *ptr)
}
