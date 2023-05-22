package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"os"
)

func main() {
	test()
}

func case1() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("result", s)
}

func test() {

	set1 := mapset.NewSet(1, 2, 3, 4)
	set2 := mapset.NewSet(4, 5, 6, 7)

	intersect := set1.Intersect(set2)
	fmt.Println(intersect)
	fmt.Println(set1)
	fmt.Println(set2)
}
