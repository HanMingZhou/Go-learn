package main

import "fmt"

func main() {

	var numbers [10]int
	z := 10
	for z > 0 {
		for i := 0; i < 10; i++ {
			numbers[i] = z
			z--
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("numbers[%v]=%v\n", i, numbers[i])
	}

	arrayB := []float64{1: 222.001, 3: 334.4, 4: 234.4, 2: 345.008}
	fmt.Println("arrayB=", arrayB)

	// 二维数组
	arrayC := [][]int{}

	row1 := []int{1, 2, 3}
	row2 := []int{10, 20, 30}

	arrayC = append(arrayC, row1)
	arrayC = append(arrayC, row2)
	fmt.Println("row1")
	fmt.Println(arrayC[0])
	fmt.Println("row2")
	fmt.Println(arrayC[1])

	fmt.Println("arrayC[0][0]", arrayC[0][0])

	// 3X4二维数组
	arrayD := [3][4]int{
		{1, 2, 3, 3},
		{4, 5, 6, 6},
		{7, 8, 9, 9},
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(arrayD[i][j])
		}
		fmt.Println("")
	}

}
