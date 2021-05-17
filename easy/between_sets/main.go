package main

import "fmt"

func main() {
	getTotalX([]int32{2, 4}, []int32{16, 32, 96})
}

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	//O(n log(n)) solution.
	// 1. find the LCM of all the integers of array A.
	// 2. find the GCD of all the integers of array B.
	// 3. Count the number of multiples of LCM that evenly divides the GCD.
	fmt.Println(a, b)
	fmt.Println(findLCM(a))
	fmt.Println(findGCD(b))
	count := 0
	c := findLCM(a)
	d := findGCD(b)
	if c < d {
		for i := c; i <= d; i++ {
			if i > 0 {
				fmt.Println(c, i, d, c*i, (c*i)%d == 0)
				if d%(c*i) == 0 {
					count++
				}
			}
		}
	} else {
		for i := d + 1; i <= c; i++ {
			if i > 0 {
				if (c*i)%d == 0 {
					count++
				}
			}
		}
	}
	fmt.Println(count)
	return int32(count)
}

func findLCM(arr []int32) int32 {
	product := int32(1)
	if len(arr) > 1 {
		for i := 0; i < len(arr)-1; i++ {
			product = arr[i] * arr[i+1]
		}
		return product / findGCD(arr)
	}
	return arr[0] / findGCD(arr)
}

func findGCD(arr []int32) int32 {
	if len(arr) > 1 {
		GCD := gcd(arr[0], arr[1])
		for i := 2; i < len(arr)-1; i++ {
			GCD = gcd(GCD, arr[i])
		}
		return GCD
	}
	return arr[0]
}

func gcd(a int32, b int32) int32 {
	if a < b {
		for i := a; int32(i) >= 2; i-- {
			if a%int32(i) == 0 && b%int32(i) == 0 {
				return int32(i)
			}
		}
	} else {
		for i := b; int32(i) >= 2; i-- {
			if a%int32(i) == 0 && b%int32(i) == 0 {
				return int32(i)
			}
		}
	}
	return int32(1)
}
