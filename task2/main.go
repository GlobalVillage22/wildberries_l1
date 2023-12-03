package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout.
*/

func main() {
	//вариант 1
	numbers := []int{2, 4, 6, 8, 10}  //Создаем слайс с числами
	c := make(chan int, len(numbers)) // Создаем канал
	wg := sync.WaitGroup{}            //Создаем WaitGroup
	for _, number := range numbers {
		wg.Add(1)
		number := number
		go func() {
			defer wg.Done()
			c <- number * number //Пишем в канал
		}()
	}
	go func() {
		wg.Wait()
		close(c) // Закрываем канал после завершения всех горутин
	}()
	for v := range c {
		fmt.Println(v)
	}
	//вариант 2
	fmt.Println("Вариант 2")
	wg2 := sync.WaitGroup{}
	for _, number := range numbers {
		wg2.Add(1)
		number := number
		go func() {
			defer wg2.Done()
			fmt.Println(number * number)
		}()
	}
	wg2.Wait()

}
