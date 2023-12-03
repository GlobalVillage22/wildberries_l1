package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

*/

func main() {
	//вариант 1
	numbers := []int{2, 4, 6, 8, 10}  //Создаем слайс с числами
	c := make(chan int, len(numbers)) // Создаем канал
	sum := 0
	wg := sync.WaitGroup{} //Создаем WaitGroup
	for _, number := range numbers {
		wg.Add(1)
		number := number
		go func() {
			defer wg.Done()
			c <- number * number
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	for v := range c {
		sum += v
	}
	fmt.Println("sum = ", sum)
	//вариант 2
	fmt.Println("Вариант 2")
	sum2 := 0
	wg2 := sync.WaitGroup{}
	for _, number := range numbers {
		wg2.Add(1)
		number := number
		go func() {
			defer wg2.Done()
			sum2 += number * number
		}()
	}
	wg2.Wait()
	fmt.Println("sum = ", sum2)

}
