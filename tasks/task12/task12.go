package main

import "fmt"

// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree").
// Создать для нее собственное множество.

func MakeStringSet(strs []string) StringSet {
	sset := StringSet{set: make(map[string]struct{})}

	for _, str := range strs {
		if _, ok := sset.set[str]; !ok {
			sset.set[str] = struct{}{}
		}
	}
	return sset
}

type StringSet struct {
	set map[string]struct{}
}

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}

	set := MakeStringSet(strs)

	fmt.Println(strs)
	fmt.Println(set)
}
