package main

import "fmt"

func main() {
	input := []int{1, 2, 100, 4, 5, 100}

	sum := GetHundreds(input)

	fmt.Println(sum)
}

func GetHundreds(numbers []int) int {
	var sum int

	for _, v := range numbers {
		if v == 100 {
			sum++
		}
	}

	return sum
}
