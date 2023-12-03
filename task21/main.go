package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

type Target interface {
	Print()
}

type Adaptee struct {
}

func (a *Adaptee) AdapteePrint() {
	fmt.Println("TargetFunc")
}

type NewAdapter struct {
	*Adaptee
}

func (a *NewAdapter) Print() {
	a.AdapteePrint()
}
func createTarget(a *Adaptee) Target {
	return &NewAdapter{a}
}
func main() {
	adapter := createTarget(&Adaptee{})
	adapter.Print()
}
