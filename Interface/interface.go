package main

import "fmt"

/*理解 Go 接口要记住一点，接口是一组方法的集合，这句话非常重要，
理解了这句话，再去理解 Go 的其他知识，
比如类型、多态、空接口、反射、类型检查与断言等就会容易很多*/
// duck
type Duck interface {
	Quack()
	DuckGo()
	isAnimal()
}

func DoDuck(d Duck) {
	d.Quack()
	d.DuckGo()
	d.isAnimal()
}

/*
因为小鸡实现了鸭子的所有方法，所以小鸡也是鸭。
那么在 main 函数中就可以这么写了。*/
type Chicken struct {
}

func (c Chicken) isAnimal() {
	fmt.Println("我是小鸡")
}
func (c Chicken) Quack() {
	fmt.Println("嘎嘎")
}
func (c Chicken) DuckGo() {
	fmt.Println("大摇大摆的走")
}
func main() {
	c := Chicken{}
	DoDuck(c)
}
