package main

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{}) //map для хранения множества с пустой структурой в качестве значений

	for _, s := range sequence {
		set[s] = struct{}{}
	}

	fmt.Println(set)
}
