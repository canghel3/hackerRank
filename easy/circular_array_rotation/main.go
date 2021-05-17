package main

func main() {
	circularArrayRotation([]int32{3, 4, 5}, 2, []int32{})
}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	newA := make([]int32, len(a))
	for index := range a {
		newA[(index+int(k))%len(a)] = a[index]
	}
	toReturn := make([]int32, len(queries))
	for index, value := range queries {
		toReturn[index] = newA[value]
	}
	return toReturn
}
