package main

import(
	"fmt"
	"time"
)

func main() {
	var c chan int = make(chan int)
	go thread(0, c)
	go writeChan(c)
	var input int
	fmt.Scanln(&input)
}

func writeChan(c chan int) {
	for{
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second *1)
	}
}

func thread(i int, c chan int) {
	var count = 0
	for{
		count = incr(count)
	}
}

func incr(count int) int{
	count++
	return count
}
