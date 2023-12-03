package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return strings.Join(arr, " ")

}
func main() {
	s := "snow dog sun"
	fmt.Println(reverseWords(s))
}
