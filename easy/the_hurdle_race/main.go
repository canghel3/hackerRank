package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hurdleRace(4, []int32{1, 6, 3, 5, 2}))
}

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	result := findMax(height) - k
	if result > 0 {
		return result
	}
	return int32(0)
}

func findMax(h []int32) int32 {
	max := math.Inf(-1)
	for _, value := range h {
		if float64(value) > max {
			max = float64(value)
		}
	}
	return int32(max)
}
