package main

/*
Реализовать пересечение двух неупорядоченных множеств.
*/
import "fmt"

func main() {
	set1 := map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}}
	set2 := map[int]struct{}{3: {}, 4: {}, 5: {}, 6: {}, 7: {}}
	intersection := make(map[int]struct{}) //map для хранения пересечения

	for i := range set1 {
		if _, ok := set2[i]; ok {
			intersection[i] = struct{}{}
		}
	}
	fmt.Println(intersection)
}
