package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(remove(arr, 2))

}
