package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

//NOT AWARDED FULL POINTS

func main() {
	r := readFile("C:\\Users\\Cristian\\go\\src\\hackerRank\\medium\\climbing_the_leaderboard\\test_ranks")
	s := readFile("C:\\Users\\Cristian\\go\\src\\hackerRank\\medium\\climbing_the_leaderboard\\test_scores")
	var n []int32
	var m []int32
	for _, ra := range r {
		i, _ := strconv.ParseInt(ra, 10, 32)
		n = append(n, int32(i))
	}
	for _, ra := range s {
		i, _ := strconv.ParseInt(ra, 10, 32)
		m = append(m, int32(i))
	}
	climbingLeaderboard(n, m)
	r = readFile("t_rank_2")
	s = readFile("t_score_2")
	n = []int32{}
	m = []int32{}
	for _, ra := range r {
		i, _ := strconv.ParseInt(ra, 10, 32)
		n = append(n, int32(i))
	}
	for _, ra := range s {
		i, _ := strconv.ParseInt(ra, 10, 32)
		m = append(m, int32(i))
	}
	climbingLeaderboard(n, m)
	climbingLeaderboard([]int32{100, 90, 90, 80}, []int32{70, 80, 105})
	climbingLeaderboard([]int32{1}, []int32{1, 1})
	climbingLeaderboard([]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120})
	climbingLeaderboard([]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102})
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	var result []int32
	var toInt []int
	var toIntPlayer []int

	ranked = beautify(ranked)

	for _, rank := range ranked {
		toInt = append(toInt, int(rank))
	}

	for _, score := range player {
		toIntPlayer = append(toIntPlayer, int(score))
	}

	sort.Ints(toInt)
	fmt.Println(len(toInt), toInt)
	fmt.Println(len(toIntPlayer), toIntPlayer)

	for _, score := range toIntPlayer {
		j := sort.Search(len(toInt), func(i int) bool { return toInt[i] >= score })
		//fmt.Println("i:", i, score)
		result = append(result, int32(len(ranked))-int32(j))
	}
	fmt.Println("result:", result)
	return result
}
func readFile(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return strings.Split(string(bytes), " ")
}

func beautify(r []int32) []int32 {
	for i := 0; i < len(r)-1; i++ {
		if r[i] == r[i+1] {
			r = append(r[:i], r[i+1:]...)
		}
	}
	return r
}
