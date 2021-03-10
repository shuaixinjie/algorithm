package main

import (
	"algorithm/practice_1"
	"fmt"
)

func main() {
	var aa [][]string = [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}

	fmt.Println(practice_1.DestCity(aa))
}
