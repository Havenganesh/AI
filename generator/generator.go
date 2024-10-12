package main

import "fmt"

func main() {
	val := generator("ann")
	for i := 0; i < 10; i++ {
		fmt.Println("you say : ", <-val)
	}
	fmt.Println("exited")
}

func generator(name string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", name, i)
		}
	}()
	return c
}
