package main

import (
	"fmt"
	"time"
)

func main() {
	// вариант 3 с использованием флага
	stop := false
	go func() {
		for {
			if stop {
				fmt.Println("exit")
				return
			}
			fmt.Println("working")
		}
	}()
	time.Sleep(1 * time.Second)
	stop = true
}
