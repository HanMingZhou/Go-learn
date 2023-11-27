package main

import (
	"fmt"
)

type Node struct {
	num  int
	next *Node
}

func showNode(node Node) {
	if node.next != nil {
		fmt.Println(node.num)
		node = *node.next
		showNode(node)
	} else {
		return
	}
}

func main() {

	listA := &Node{}
	listB := &Node{}
	listC := &Node{}
	listD := &Node{}
	listF := &Node{}

	listA.num = 1
	listA.next = listB
	listB.num = 2
	listB.next = listC
	listC.num = 3
	listC.next = listD
	listD.num = 4
	listD.next = listF
	listF.num = 5
	listF.next = nil

	showNode(*listA)

	x := 100.0
	result := sqrt(x)
	fmt.Printf("%.2f 的平方根为 %.6f\n", x, result)
}

// 1、第一递归函数功能
func sqrtRecursive(x, guess, prevGuess, epsilon float64) float64 {
	//2、找出递归结束的条件
	if diff := guess*guess - x; diff < epsilon && -diff < epsilon {
		fmt.Println("diff", diff)
		return guess
	}
	//3要素、找出函数的等价关系式
	newGuess := (guess + x/guess) / 2
	fmt.Println("newGuess", newGuess)
	if newGuess == prevGuess {
		return guess
	}

	return sqrtRecursive(x, newGuess, guess, epsilon)
}
func sqrt(x float64) float64 {
	return sqrtRecursive(x, 1.0, 0.0, 1e-9)
}
