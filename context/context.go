package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second )
	defer cancel()

	//handle(ctx, 500 * time.Millisecond)
	handle(ctx, 1500 * time.Millisecond)
	select {
	case <- ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <- ctx.Done():
		fmt.Println("handle", ctx.Err())

	case <- time.After(duration) :
		fmt.Println("handle work finished in", duration)
	}
}