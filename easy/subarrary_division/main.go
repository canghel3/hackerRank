package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []int32
	d := readFile("C:\\Users\\Cristian\\go\\src\\hackerRank\\easy\\subarrary_division\\test_input")
	for _, a := range d {
		i, _ := strconv.ParseInt(a, 10, 32)
		input = append(input, int32(i))
	}
	fmt.Println(birthday(input, 26, 8))
	fmt.Println(birthday([]int32{1, 2, 1, 3, 2}, 3, 2))
	fmt.Println(birthday([]int32{4}, 4, 1))
}

func birthday(s []int32, d int32, m int32) int32 {
	result := int32(0)
	for i := 0; i <= len(s)-int(m); i++ {
		//fmt.Println(s, i, i + int(m), sum(s[i:i+int(m)]))
		if sum(s[i:i+int(m)]) == d {
			result++
		}
	}
	return result
}

func sum(s []int32) (sum int32) {
	for _, e := range s {
		sum += e
	}
	return
}

func readFile(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return strings.Split(string(bytes), " ")
}
