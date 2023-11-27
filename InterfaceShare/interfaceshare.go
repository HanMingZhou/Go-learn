package main

import "fmt"

func main() {

	var sharp Sharp
	sharp = Triangle{10, 2}
	fmt.Println(sharp.area())

	/*
		Triangle_a := Triangle{10, 2}
		fmt.Println(Triangle_a.area())
	*/

	/*
		Triangle_a := Triangle{10, 2}
		area(Triangle_a)
	*/
}

type Sharp interface {
	// 定义求area面积行为
	area() float32
}

type Triangle struct {
	length float32
	hight  float32
}

func (T Triangle) area() float32 {

	return (T.hight * T.length) / 2 // 实现了类的行为
}

type Circle struct {
	redius float32
	pri    float32
}

func (C Circle) area() float32 {
	return C.pri * C.redius * C.redius // 实现了类的行为
}

func area(sharp Sharp) {
	sharp.area()
}
