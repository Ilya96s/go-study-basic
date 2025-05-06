package main

import "fmt"

func main() {
	number := 5
	numberPointer := &number

	str := "hello"
	strPointer := &str

	fmt.Println("number до вызова foo:", number)
	foo(numberPointer)
	fmt.Println("number после вызова foo:", number)

	fmt.Println("number до вызова bar:", str)
	bar(strPointer)
	fmt.Println("number после вызова bar:", str)
}

func foo(ptr *int) {
	*ptr = 4444
}
func bar(ptr *string) {
	*ptr = "new string"
}
