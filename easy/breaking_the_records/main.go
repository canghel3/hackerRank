package main

import (
	"fmt"
	"math"
)

func main() {
	breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1})
}

func breakingRecords(scores []int32) []int32 {
	min := math.Inf(0)
	max := math.Inf(-1)
	brokeMax := -1
	brokeMin := -1
	for _, score := range scores {
		if float64(score) > max {
			max = float64(score)
			brokeMax++
		}
		if float64(score) < min {
			min = float64(score)
			brokeMin++
		}
	}
	fmt.Println([]int32{int32(brokeMax), int32(brokeMin)})
	return []int32{int32(brokeMax), int32(brokeMin)}
}
