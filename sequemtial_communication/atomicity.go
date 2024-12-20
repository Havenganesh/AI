package main

import "sync"

type Counter struct {
	Mu    sync.Mutex
	Value int
}

func (c *Counter) Increment() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Value++
}

// func main() {
// 	c := &Counter{}
// 	c.increment()
// }
