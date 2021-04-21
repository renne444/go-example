package main

import "fmt"

func main() {
	a, b := make(chan int), make(chan int)
	go func() {
		b <- 7
	}()
	select {
	case c := <-a:
		fmt.Println(fmt.Sprintf("c<-a %d", c))
	case c := <-b:
		fmt.Println(fmt.Sprintf("c<-b %d", c))
	default:
		fmt.Println("nothing in chan")
	}
}
