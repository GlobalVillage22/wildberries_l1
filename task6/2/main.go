package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// вариант 2 с использованием контекста
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("exit")
				return
			default:
				fmt.Println("working")
			}
		}
	}()
	time.Sleep(1 * time.Second)
	cancel()
}
