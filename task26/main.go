package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func unique(s string) bool {
	s = strings.ToLower(s)         // приводим к нижнему регистру
	set := make(map[rune]struct{}) // создаем мапу символов
	for _, r := range s {
		if _, ok := set[r]; ok {
			return false // если символ уже есть в мапе, то возвращаем false
		}
		set[r] = struct{}{} // если нет, то добавляем
	}
	return true
}

func main() {
	fmt.Println(unique("abcd"))
	fmt.Println(unique("abCdefAaf"))
	fmt.Println(unique("aabcd"))
}
