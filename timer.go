package main

import (
	"fmt"
	"time"
	"sync"
)

type mutexCounter struct {
	mu sync.Mutex
	x int64
	mas []int
}

func (c *mutexCounter) Increment(x int64, i int) {
	c.mu.Lock()
	c.x += x
	c.mas[i] += 1
	c.mu.Unlock()
}

func (c *mutexCounter) Value() (x int64){
	c.mu.Lock()
	x = c.x
	c.mu.Unlock()
	return
}
/*
type intCounter int64

func (c *intCounter) Add(x int64){
	*c++
}*/

func main() {
	counter := mutexCounter{}
	counter.mas = make([]int, 21)
	for i:= 0; i < 20; i++{
		//counter.mas = append(counter.mas, i)
		go func(no int){
			for j:= 0; j < 10000000; j++{
				counter.Increment(1, i)
			}
		}(i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println(counter.Value())
	fmt.Println(counter.mas)
}
