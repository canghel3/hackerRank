package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(biggerIsGreater("ab"))
	fmt.Println(biggerIsGreater("bb"))
	fmt.Println(biggerIsGreater("dcba"))

	fmt.Println(biggerIsGreater("hefg"))
	fmt.Println(biggerIsGreater("dhck"))
	fmt.Println(biggerIsGreater("dkhc"))
}

func biggerIsGreater(w string) string {
	// Write your code here
	wSlice := split(w)
	i, a := findLongestSuffix(w)
	if i == 0 {
		return "no answer"
	}
	aSlice := split(a)
	pivot := ""
	pivot = string(w[i-1])
	_, index := swapIndex(pivot, a)
	temp := wSlice[i-1]
	wSlice[i-1] = aSlice[index]
	aSlice[index] = temp
	reverse(aSlice)
	wSlice = append(wSlice[:i], aSlice...)
	return strings.Join(wSlice, "")
}

func findLongestSuffix(w string) (i int, a string) {
	if len(w) > 1 {
		for j := 1; j < len(w); j++ {
			if w[j] > w[j-1] {
				i = j
			}
		}
	}
	return i, w[i:]
}

func swapIndex(pivot string, a string) (i string, index int) {
	for i := len(a) - 1; i >= 0; i-- {
		if string(a[i]) > pivot {
			return string(a[i]), i
		}
	}
	return "", 0
}

func split(w string) []string {
	sl := make([]string, len(w))
	for index, v := range w {
		sl[index] = string(v)
	}
	return sl
}

func reverse(a []string) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
