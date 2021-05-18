package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, "abc"))
}

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	max := math.Inf(-1)
	mapping := make(map[int32]int32)
	char_ := 'a'
	for _, height := range h {
		mapping[char_] = height
		if float64(height) > max {
			max = float64(height)
		}
		char_++
	}
	maxLetter := 'a'
	for _, letter := range word {
		if mapping[letter] > mapping[maxLetter] {
			maxLetter = letter
		}
	}
	//fmt.Println(mapping, mapping[maxLetter])
	return mapping[maxLetter] * int32(len(word))
}
