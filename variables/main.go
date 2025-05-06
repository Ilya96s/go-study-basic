package main

import "fmt"

func main() {
	score := 0

	fmt.Println("Get ready")
	fmt.Println("Score:", score)
	fmt.Println("")

	for i := 1; i <= 10; i++ {
		fmt.Println("----------------")
		fmt.Println("Вы подлетаете к трубе", i)
		fmt.Println("")

		fmt.Println("Вы пролетаете трубу", i)
		fmt.Println("")

		fmt.Println("Вы пролетели трубу", i)
		fmt.Println("")

		score++

		fmt.Println("Score:", score)
		fmt.Println("")
	}
}
