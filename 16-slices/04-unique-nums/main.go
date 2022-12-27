package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	max, _ := strconv.Atoi(os.Args[1])
	rand.Seed(time.Now().UnixNano())

	var uniques []int

loop:
	for found := 0; found < max; {
		n := rand.Intn(max) + 1

		for _, u := range uniques {
			if u == n {
				// 如果有重複就重新運行
				continue loop
			}
		}

		uniques = append(uniques, n)
		found++
	}

	fmt.Println("uniques: ", uniques)
}
