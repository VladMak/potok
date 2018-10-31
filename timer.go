package main

import (
	"fmt"
	"sync"
	"time"
)

type mutexCounter struct {
	mu  sync.Mutex
	x   int64
	mas []int
}

func (c *mutexCounter) Increment(x int64, i int) {
	c.mu.Lock()
	c.x += x
	c.mas[i] += 1
	c.mu.Unlock()
}

func (c *mutexCounter) Value() (x int64) {
	c.mu.Lock()
	x = c.x
	c.mu.Unlock()
	return
}

type nomutCounter struct {
	x int64
	mas []int
}

func (c *nomutCounter) Add(x int64, i int){
	c.x += x
	c.mas[i] += 1
}

func main() {
	//counter := nomutCounter{}
	counter := mutexCounter{}
	counter.mas = make([]int, 20)
	for i := 0; i < 20; i++ {
		go func(no int) {
			for j := 0; j < 10000000; j++ {
				counter.Increment(1, no)
				//counter.Add(1, no)
			}
		}(i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println(counter.x)
	fmt.Println(counter.mas)
}
