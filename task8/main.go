package main

import "fmt"

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func changeBitTo0(n int64, i int64) int64 {
	return n &^ (1 << i) // установка i-го бита в 0
}
func changeBitTo1(n int64, i int64) int64 {
	return n | (1 << i) // установка i-го бита в 1
}

func main() {
	var n int64
	var i int64
	fmt.Scan(&n)
	fmt.Scan(&i)
	fmt.Println(changeBitTo0(n, i))
	fmt.Println(changeBitTo1(n, i))

}
