package main

import "fmt"

var (
	routineNum = 5
)

func goRoutine(ch *chan int) {
	fmt.Println("go_r")
	*ch <- 1
	return
}

func main() {

	var ch = make(chan int)
	defer func() {
		close(ch)
	}()
	for i := 0; i < routineNum; i++ {
		go goRoutine(&ch)
	}

	size := 0
	fmt.Println("size = ", size)

	for {
		<-ch
		size++
		fmt.Printf("%d of routine finished\n", size)
		if size >= routineNum {
			fmt.Println("all routine finished")
			break
		}
	}
}
