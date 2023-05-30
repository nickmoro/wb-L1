package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func Intersect(set1, set2 map[int]struct{}) map[int]struct{} {
	res := make(map[int]struct{})
	for elem := range set1 {
		if _, ok := set2[elem]; ok {
			res[elem] = struct{}{}
		}
	}
	return res
}

func main() {
	set1 := map[int]struct{}{9: {}, 3: {}, 7: {}, 5: {}, 1: {}}
	set2 := map[int]struct{}{4: {}, 8: {}, 6: {}, 2: {}, 3: {}}
	intersection := Intersect(set1, set2)

	fmt.Println(set1)
	fmt.Println(set2)
	fmt.Println(intersection)
}
