package main

import "fmt"

func main() {
	fmt.Println(Greet("Banana"))
}

func Greet(input string) string {
	return "Hello, " + input
}
