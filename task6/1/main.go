package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	//вариант 1 через канал
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("горутина завершена")
				return
			default:
				fmt.Println("горутина работает")
			}

		}
	}()
	time.Sleep(1 * time.Second)
	ch <- "exit"
	wg.Wait()
}
