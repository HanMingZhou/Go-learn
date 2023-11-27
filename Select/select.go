package main

import (
	"fmt"
	"time"
)

func main() {

	channelA := make(chan string)
	channelB := make(chan string)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			channelA <- "i am channelA"
		}

	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			channelB <- "i am channelB"
		}

	}()

	for {

		time.Sleep(8 * time.Second) //有多个 case 都可以运行，select 会随机公平地选出一个执行，其他不会执行

		select { //select 语句只能用于通道操作，每个 case 必须是一个通道操作，要么是发送要么是接收。

		case messageA := <-channelA:
			fmt.Println("Received messageA", messageA)
		case messageB := <-channelB:
			fmt.Println("Received messageB", messageB)
		}

	}

}
