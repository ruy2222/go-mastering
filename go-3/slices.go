package main

import "fmt"

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func main() {
	sliceTest := make([]int, 20)
	for i := 1; i < len(sliceTest); i++ {
		fmt.Println(sliceTest[i], " ")
	}
	sliceTest = append(sliceTest, 50)
	fmt.Println(sliceTest)
	reslice := sliceTest[1:3]
	fmt.Println(reslice)
	reslice[0] = 40
	reslice[1] = 50
	fmt.Println(reslice)

	//cap
	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

	aSlice = append(aSlice, -2)
	aSlice = append(aSlice, -3)
	// aSlice = append(aSlice, -4)
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

	//byte slice
	s := make([]byte, 5)
	fmt.Println(s, " ")

	//copy function
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	copy(b4, b6)
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)

	//MultiDimensional Slice
	s1 := make([][]int, 4)
	fmt.Println(s1, " ")
}
