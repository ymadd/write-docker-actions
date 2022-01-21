package main

import (
	"fmt"
	"os"
)

func main() {

	firstGreeting := os.Getenv("INPUT_FIRST_GREETING")
	secondGreeting := os.Getenv("INPUT_SECOND_GREETING")
	thirdGreeting := os.Getenv("INPUT_THIRD_GREETING")

	fmt.Println("Hello" + firstGreeting)
	fmt.Println("Hello" + secondGreeting)

	if thirdGreeting != "" {
		fmt.Println("Hello" + thirdGreeting)
	}
}
