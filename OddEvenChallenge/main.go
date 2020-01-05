package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, index := range numbers {
		if index%2 == 0 {
			fmt.Println(index, " is even")
		} else {
			fmt.Println(index, " is odd")
		}
	}
}
