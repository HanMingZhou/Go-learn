package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		go say("hello")
		say("world")

		s := []int{1, 2, 3, 4, -1, -2, -3}
		fmt.Println("len(s)=", len(s))
		c := make(chan int)

		go sum(s[len(s)/2:], c)
		go sum(s[:len(s)/2], c)
		x, y := <-c, <-c
		fmt.Println("x =", x, "y=", y)

		channel_B := make(chan int, 5)
		for i := 0; i < 6; i++ {
			channel_B <- i
			fmt.Println("channel_B 发送成功，i=", i)
		}
	*/
	channel_C := make(chan int)
	go Channel_Write(channel_C)
	go Channel_Read(channel_C)
	time.Sleep(time.Second)

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func Channel_Write(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println("channel_c 写入成功，i=", i)
	}
}

func Channel_Read(c chan int) {
	//time.Sleep(100 * time.Millisecond)
	for i := 0; i < 10; i++ {
		j := <-c
		fmt.Println("channel_c 读取成功，i=", j)
	}
}
