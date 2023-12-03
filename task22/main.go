package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
значение которых > 2^20.
*/

func main() {
	var a, b big.Int
	fmt.Scan(&a, &b)

	fmt.Println(a.Mul(&a, &b)) //a*b
	fmt.Println(a.Div(&a, &b)) //   a/b
	fmt.Println(a.Add(&a, &b)) //   a+b
	fmt.Println(a.Sub(&a, &b)) //   a-b
}
