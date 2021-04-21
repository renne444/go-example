package main

import (
	"fmt"
	"time"
)

func run(w int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if w % 5 == 0 {
		panic(w)
	} else {
		fmt.Println(w)
	}

}

func main() {
	for i := 0; i < 10 ;i++ {
		go run(i)
	}
	time.Sleep(time.Second)
}

