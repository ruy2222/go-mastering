package main

import "fmt"

func main() {
	anArray := [4]int{3, 4, 2, 5}
	twoD := [4][4]int{{3, 2, 3, 4}, {4, 4, 4, 4}, {2, 5, 6, 6}, {2, 4, 5, 6}}
	threeD := [2][2][2]int{{{1, 6}, {2, 4}}, {{6, 2}, {2, 4}}}

	fmt.Println("The length of", anArray, "is", len(anArray))
	fmt.Println("The first element of", twoD, "is", twoD[0][0])
	fmt.Println("The length of", threeD, "is", len(threeD))
	for i := 0; i < len(threeD); i++ {
		v := threeD[i]
		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
		}
		fmt.Println()
	}
	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
		fmt.Println()
	}

}
