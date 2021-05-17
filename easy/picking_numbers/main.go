package main

import "fmt"

//n=int(input())
//l=[int(i) for i in input().split()]
//maximum=0
//for i in l:
//	c=l.count(i)
//	d=l.count(i-1)
//	c=c+d
//	if c > maximum:
//		maximum=c
//print(maximum)

func main() {
	pickingNumbers([]int32{4, 6, 5, 3, 3, 1})
}

func pickingNumbers(a []int32) int32 {
	// Write your code here
	currentCount := 0
	previousCount := 0
	max := 0
	for _, n := range a {
		currentCount = int(count(n, a))
		previousCount = int(count(n-1, a))
		currentCount = currentCount + previousCount
		if currentCount > max {
			max = currentCount
		}
	}
	fmt.Println(max)
	return int32(max)
}

func count(i int32, a []int32) int32 {
	parsed := i
	count := 0
	for _, n := range a {
		if n == parsed {
			count++
		}
	}
	return int32(count)
}
