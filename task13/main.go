package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	//вариант 1
	a, b := 1, 2
	a, b = b, a
	fmt.Println(a, b)
	//вариант 2
	a, b = 1, 2
	swap := func(a, b int) (int, int) {
		return b, a
	}
	a, b = swap(a, b)
	fmt.Println(a, b)
	//вариант 3
	a, b = 1, 2
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
	//вариант 4
	a, b = 1, 2
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)

}
