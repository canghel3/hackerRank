package main

import "fmt"

func main() {
	fmt.Println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	result := int32(0)
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				result++
			}
		}
	}
	return result
}
