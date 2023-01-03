package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var start, stop int

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])

	}

	if start == 0 || stop == 0 {
		start, stop = 'A', 'Z'
	}
	fmt.Println("Start", start, stop)

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n", "literal", "dec", "hex", "encoded", strings.Repeat("-", 45))

	for n := start; n <= stop; n++ {
		// %d => 10進制數字
		// %[1]x => hex 16進制數字
		// %x => utf8
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(rune(n)))

	}

	// unicode code 表情符號是 1~4 byte 的大小, 因此不能用 byte 來配置, 須改用 rune
	// var char byte = '🤩'
	var char rune = '🤩'
	fmt.Println("char", char)
}
