package main

import (
	"fmt"
	"time"
)

func timer(bol *bool) {
	tim := time.Tick(time.Second * 5)
	for _ = range tim {
		fmt.Println("TIMER STOP 1")
		*bol = false
	}
	fmt.Println("TIMER STOP")
	*bol = false
}

func thread(id int, c chan int, bol *bool){
	fmt.Println(id)
	for i := 0; ; i++ {
		if !*bol{
			fmt.Println("Thread STOP")
			fmt.Println(i)
			c <- i
			break
		}
	}
}

func main() {
	var bol bool = true
	var c chan int = make(chan int)
	go timer(&bol)

	for i := 0; i < 1; i++ {
		go thread(i, c, &bol)
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("Thats all")
	fmt.Println(<- c)
}
