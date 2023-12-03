package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// Создание структуры Human
type Human struct {
	name   string
	gender string
}

// Создание структуры Action
type Action struct {
	Human     //Встраивание структуры Human
	something string
}

// Метод для структуры Human
func (h *Human) SetName(name string) {
	h.name = name
}

func main() {
	human := Human{}
	action := Action{}
	human.SetName("Ivan")
	action.SetName(human.name) //Встраивание метода SetName
	fmt.Println(human)
	fmt.Println(action)
}
