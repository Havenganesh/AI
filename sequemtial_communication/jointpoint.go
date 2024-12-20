package main

import (
	"fmt"
	"sync"
)

func JointPoint() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("Hello")
	}
	wg.Add(1)
	sayHello()
	wg.Wait()
}
