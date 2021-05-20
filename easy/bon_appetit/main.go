package main

import "fmt"

func main() {
	bonAppetit([]int32{3, 10, 2, 9}, 1, 12)
}

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	didNotEat := bill[k]
	correctSum := float64(sumExcept(didNotEat, bill)) / 2
	if correctSum == float64(b) {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(int(float64(b) - correctSum))
	}
}

func sumExcept(i int32, bill []int32) (sum int32) {
	for _, v := range bill {
		if v != i {
			sum += v
		}
	}
	return
}
