package practice_1

import "testing"

func TestDestCity(t *testing.T) {
	var arrtwo [][]string = [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	result := destCity(arrtwo)
	if result != "A" {
		t.Fatal("error")
	}

}
