package main

import (
	"fmt"
	"time"
)

func main() {
	go boring()
	fmt.Println("exited")
}

func boring() {
	for i := 0; ; i++ {
		fmt.Println("it is boring : ", i)
		time.Sleep(time.Second * 1)
	}
}
