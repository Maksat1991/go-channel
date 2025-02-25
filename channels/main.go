package main

import (
	"fmt"
	"time"
)

func main() {
	const bufSize = 3

	ch := make(chan int)
	//first - read from channel, first - close channel, third - break the range loop
	run(ch, bufSize)

	chBuf := make(chan int, bufSize)
	//first - close channel, second - read from closed channel, third - break the range loop
	run(chBuf, bufSize)

}

func run(ch chan int, n int) {
	fmt.Println("run started")
	go square(ch, n)

	//The loop is breaks when channel is closed AND don't have any data
	for val := range ch {
		fmt.Printf("val = %d\n", val)
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println("run stopped")
	fmt.Println()
}

func square(ch chan<- int, n int) {
	for i := 1; i <= n; i++ {
		ch <- i * i
	}
	close(ch)
	fmt.Println("channel closed")
}
