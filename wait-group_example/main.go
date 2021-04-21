package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0;i < 2;i++ {
		go func() {
			wg.Done()
		}()
	}
	for i := 0;i < 2;i++ {
		go func() {
			wg.Wait()
		}()
	}
	time.Sleep(2 * time.Second)
}
