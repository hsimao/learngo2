package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type a search word.")
		return
	}

	query := args[0]

	rx := regexp.MustCompile(`[^a-z]+`)

	in := bufio.NewScanner(os.Stdin)
	// 依據每個單字 Scan
	in.Split(bufio.ScanWords)

	// 字典
	words := make(map[string]bool)

	// 每次掃描成功執行
	for in.Scan() {
		word := strings.ToLower(in.Text())

		// 移除特殊符號, 只取 a~z字符
		word = rx.ReplaceAllString(word, "")

		if len(word) > 2 {
			words[word] = true
		}

		fmt.Println(word)

	}

	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}

	fmt.Printf("The input does not contain %q.\n", query)
}

// DEMO:
// 抓取網路檔案來判斷
// curl -s https://www.rfc-editor.org/rfc/rfc20.txt | go run main.go "inclusive"
// 抓取本地檔案來判斷
// go run main.go night < shakespeare.txt
