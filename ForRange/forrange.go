package main

import (
	"fmt"
	"os"
)

func main() {
	sliceA := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for _, V := range sliceA {
		fmt.Println("sliceA =", V)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	fmt.Println(len(os.Args))
	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	quantity := 10
	fruits := "apples"
	text := fmt.Sprint("I have ", quantity, fruits)
	fmt.Printf("Value stored in text is: '%v'\n", text)
	text = fmt.Sprintf("%v,%v", quantity, fruits)
	fmt.Println(text)
}
