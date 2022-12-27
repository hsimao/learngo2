package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

// 產出不重複的隨機數字
func main() {
	max, _ := strconv.Atoi(os.Args[1])
	rand.Seed(time.Now().UnixNano())

	var uniques []int

loop:
	for len(uniques) < max {
		n := rand.Intn(max) + 1

		for _, u := range uniques {
			if u == n {
				// 如果有重複就重新運行
				continue loop
			}
		}

		uniques = append(uniques, n)

	}

	fmt.Println("uniques: ", uniques)
	// 小到小排序
	sort.Ints(uniques)
	fmt.Println("sort uniques: ", uniques)

	nums := [5]int{5, 4, 3, 2, 1}
	sort.Ints(nums[:])
	fmt.Println("sort nums: ", nums)

}
