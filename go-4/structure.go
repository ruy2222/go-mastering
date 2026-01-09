package main

import (
	"fmt"
)

type mystruct struct {
	name   string
	age    int
	height int
}

func pStruct(a string, b, h int) *mystruct {
	if h > 300 {
		h = 0
	}
	return &mystruct{a, b, h}
}

func cStruct(a string, b, h int) mystruct {
	if h > 300 {
		h = 0
	}
	return mystruct{a, b, h}
}

func main() {
	aMan := pStruct("aba", 44, 175)
	bMan := cStruct("aaa", 44, 175)
	fmt.Println(*aMan, " ")
	fmt.Println(bMan, " ")
}
