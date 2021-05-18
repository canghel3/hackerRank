package main

import (
	"fmt"
	"math"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}

	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func main() {
	fmt.Println(encryption("haveaniceday"))
	fmt.Println(encryption("chillout"))
	fmt.Println(encryption("feedthedog"))
	fmt.Println(encryption("ifmanwasmeanttostayonthegroundgodwouldhavegivenusroots"))
	fmt.Println(encryption("wclwfoznbmyycxvaxagjhtexdkwjqhlojykopldsxesbbnezqmixfpujbssrbfhlgubvfhpfliimvmnny"))
}

func encryption(s string) string {
	// Write your code here
	l := len(s)
	rowCount := math.Floor(math.Sqrt(float64(l)))
	columnCount := math.Ceil(math.Sqrt(float64(l)))
	if int(rowCount)*int(columnCount) < l {
		rowCount, columnCount = increment(int(rowCount), int(columnCount), l)
	}
	grid := split(s, rowCount, columnCount)
	result := ""
	rememberI := 0
	//fmt.Println(columnCount)
	//fmt.Println(rowCount)
	if grid != nil {
		Block{
			Try: func() {
				for i := 0; i < int(columnCount); i++ {
					for j := 0; j < int(rowCount); j++ {
						current := grid[j] // this takes the current word, we need to take the first letter from first, first from second, from third and so on
						//fmt.Println(grid[j], j)
						word := current[0] //current[0] yields the actual word since current is an array containing only 1 element
						result += string(word[i])
						//fmt.Println("result = ", result)
						rememberI = i
					}
					result += " "
				}
				//Throw("error encountered")
			},
			Catch: func(e Exception) {
				fmt.Printf("Caught error %v\n", e)
			},
			Finally: func() {
				for i := rememberI + 1; i < int(columnCount); i++ {
					result += " "
					for j := 0; j < int(rowCount)-1; j++ {
						current := grid[j] // this takes the current word, we need to take the first letter from first, first from second, from third and so on
						//fmt.Println(grid[j], j)
						word := current[0] //current[0] yields the actual word since current is an array containing only 1 element
						result += string(word[i])
						//fmt.Println("result = ", result)
					}
				}
			},
		}.Do()
	}
	return result
}

func split(s string, rowCount float64, columnCount float64) [][]string {
	var grid [][]string
	remaining := ""
	for i := float64(0); i < rowCount; i++ {
		var newRow []string
		if int(columnCount) < len(s) {
			newRow = append(newRow, s[:int(columnCount)])
			s = s[int(columnCount):]
			remaining = s
		} else {
			newRow = append(newRow, s[:])
		}
		grid = append(grid, newRow)

	}
	var new []string
	new = append(new, remaining)
	grid = append(grid, new)
	grid = cleanUp(grid)
	return grid
}

func cleanUp(s [][]string) [][]string {
	if s[len(s)-1][0] == s[len(s)-2][0] {
		s = append(s[:len(s)-2], s[len(s)-1:]...)
	}
	return s
}

func increment(i int, j int, l int) (float64, float64) {
	for i*j < l {
		if i < j {
			i++
		} else {
			j++
		}
	}
	return float64(i), float64(j)
}
