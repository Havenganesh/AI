package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go boring("ann", c)
	go boring("joe", c)
	for i := 0; i < 10; i++ {
		fmt.Println("received : ", <-c)
	}
	fmt.Println("exited")
}

func boring(name string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", name, i)
		time.Sleep(time.Second * 1)
	}
}
