package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	os.Stdin.Close()

	in := bufio.NewScanner(os.Stdin)

	var lines int

	// 等待用戶輸入文字
	for in.Scan() {
		lines++
		// 輸入完後馬上印出
		// fmt.Println("Scanned Text :", in.Text())
	}

	fmt.Printf("There are %d line(s)\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}

}
