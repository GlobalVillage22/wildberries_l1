package main

import (
	"bytes"
	"fmt"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10) //создаем большую строку
	var buf bytes.Buffer
	buf.WriteString(v[:100])  //заполняем буфер
	justString = buf.String() //возвращаем строку
}

func createHugeString(size int) string { //функция создания большой строки
	var buf bytes.Buffer //создаем буфер
	for i := 0; i < size; i++ {
		buf.WriteByte('a') //заполняем буфер
	}
	return buf.String() //возвращаем строку
}

func main() {
	someFunc()
	fmt.Println(justString)
}
