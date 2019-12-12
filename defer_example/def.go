package main

import (
	"fmt"
	"time"
)

func wait(index int) {
	time.Sleep(time.Second * 10)
	fmt.Println("finished", index)
}

func main() {

	for i := 0; i < 10; i++ {
		w := i
		go wait(w)
		defer func() {
			fmt.Println(w)
		}()
	}
	fmt.Println("finish")
}
