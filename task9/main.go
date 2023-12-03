package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ch1 := make(chan int) //канал для записи
	ch2 := make(chan int) //канал для вывода результата
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		for _, n := range nums {
			ch1 <- n //пишем в канал
		}
		close(ch1) // закрываем канал после записи всех чисел
	}()
	go func() {
		defer wg.Done()
		for n := range ch1 {
			ch2 <- n * 2 // читаем из канала ch1 и записываем в ch2
		}
		close(ch2)
	}()
	go func() {
		defer wg.Done()
		for n := range ch2 {
			fmt.Print(n, " ")
		}
	}()

	wg.Wait()
}
