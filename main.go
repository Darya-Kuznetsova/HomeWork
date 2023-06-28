package main

import "fmt"

func main() {
	array := [4]int{6, 9, 8, 1}
	fmt.Println(even_multiplication(array[:]))
}

func even_multiplication(array []int) int {
	result := 1
	for _, v := range array {
		if v%2 == 0 {
			result *= v
		}
	}
	return result
}
