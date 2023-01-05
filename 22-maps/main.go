package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}

	query := args[0]

	dict := map[string]string{
		"good":    "qiyi",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel",
	}

	// 直接設定 nil 並不能完全清除 map 資料, 必須要 loop 逐一刪除所有屬性
	// dict = nil
	for k := range dict {
		delete(dict, k)
	}

	delete(dict, "awesome")

	turkish := make(map[string]string, len(dict))
	for k, v := range dict {
		turkish[v] = k
	}

	fmt.Printf("english: %q\n turkish: %q\n", dict, turkish)

	if value, ok := turkish[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)
	fmt.Printf("# of keys: %d\n", len(dict))

}
