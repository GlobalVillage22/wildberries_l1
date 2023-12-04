package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	var a interface{}
	var ch chan int
	a = ch
	switch a.(type) { // определяем тип переменной через switch
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel")
	}
	fmt.Println(reflect.TypeOf(a)) //определяем тип переменной через reflect
	fmt.Println(reflect.ValueOf(a).Type())
	fmt.Println(reflect.ValueOf(a).Kind())
	fmt.Printf("%T\n", a)

}
