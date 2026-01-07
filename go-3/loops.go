package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%20 == 0 {
			continue
		}
		if i == 95 {
			break
		}
		fmt.Println(i, " ")
	}
	fmt.Println("")
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Println(i, " ")
		i--
	}
	fmt.Println(" ")

	array := [5]int{1, 2, 3, 4, 5}
	for i, value := range array {
		if i < 10 {
			fmt.Println("index", i, "value", value)
		}
	}
}
