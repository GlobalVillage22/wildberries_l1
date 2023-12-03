package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan int) {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}(ch)

	t := 5 * time.Second
	for {
		select {
		case <-time.After(t):
			fmt.Println("exit")
			close(ch)
			wg.Wait()
			return
		default:
			ch <- rand.Intn(100)
		}
	}
}
