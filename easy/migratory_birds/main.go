package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
}

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	log := make(map[int32]int32)
	for _, value := range arr {
		log[value] = log[value] + 1
	}
	return findMax(log)
}

func findMax(l map[int32]int32) int32 {
	max := math.Inf(-1)
	result := int32(0)
	for key, value := range l {
		if float64(value) > max {
			max = float64(value)
			result = key
		}
	}
	return result
}
