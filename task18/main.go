package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
// структура счетчика
type counter struct {
	val int
	mu  sync.Mutex
}

// функция увеличения счетчика
func (c *counter) increment(x int, wg *sync.WaitGroup) {
	c.mu.Lock()
	defer wg.Done()
	c.val += x
	c.mu.Unlock()
}

// функция получения значения счетчика
func (c *counter) get() int {
	return c.val
}

// функция создания счетчика
func newCounter() *counter {
	return &counter{
		val: 0,
	}
}

func main() {
	c := newCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go c.increment(rand.Intn(10), &wg)
	}
	wg.Wait()
	fmt.Println(c.get())

}
