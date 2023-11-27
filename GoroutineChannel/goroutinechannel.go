package main

import (
	"fmt"
	"time"
)

func main() {

	// c := make(chan int, 2)
	// c <- 1
	// c <- 2
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// c1 := make(chan int, 10)
	// go Febonacci(cap(c1), c1)

	// for v := range c1 {
	// 	fmt.Println("out", time.Now())
	// 	time.Sleep(10 * time.Millisecond)
	// 	fmt.Println(v)
	// }

	// channel := make(chan string)
	// go say("hello")
	// go say2("world", channel)

	// fmt.Println(<-channel)

	Closechanenl := make(chan int)
	go CloseChanenl(Closechanenl)

	/* 		after close channel 1s, what is happing?
	time.Sleep(1* time.Second)
	 		after close,the channel canot read anymore.
			if the channle has buffer, after closing, can still read. */
	fmt.Println(<-Closechanenl)

}

func Febonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println("in", time.Now())
		time.Sleep(10 * time.Millisecond)
		c <- x
		x, y = y, x+y
	}
	// 此时 channel还未关闭
	close(c)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, "*", i+1)
	}
}

func say2(s string, c chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, "*", i+1)
	}

	c <- "done..."
	close(c)
}

func CloseChanenl(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	close(c)
}
