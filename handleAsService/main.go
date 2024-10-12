package main

import "fmt"

func main() {
	penguin := ServiceBird("penguin")
	peacock := ServiceBird("peacock")

	for i := 0; i < 10; i++ {
		select {
		case p := <-penguin:
			fmt.Println("bird received : ", p.name, p.count)
		case pea := <-peacock:
			fmt.Println("bird received : ", pea.name, pea.count)

		}

	}
}

func ServiceBird(name string) chan bird {
	b := make(chan bird)
	go func() {
		for i := 0; ; i++ {
			bird := bird{name: name, count: i}
			b <- bird
		}
	}()
	return b
}

type bird struct {
	name  string
	count int
}
