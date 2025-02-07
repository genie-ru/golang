package main

import "fmt"

func main() {
	emptyMap := make(map[string]int)
	mapWithElements :=map[string]int{"a" : 1, "b" : 2}
	fmt.Println("Maps - map1 ", emptyMap, " and map2 ", mapWithElements)

	latte := Coffee{Name: "Latte", Size: 3}
	fmt.Println("Coffee instance ", latte)
}

type Coffee struct {
	Name string
	Size int
}