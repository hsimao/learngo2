package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	str := "Yūgen ☯ 💀"

	// 不能改字串中的某的字母, 若要改需要轉型, 且會重新產生記憶體位置, 不會改到原本
	bytes := []byte(str)
	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))

	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(bytes))
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))

	fmt.Println()

	for i, r := range str {
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}

	fmt.Println()
	fmt.Printf("1st byte : %c\n", str[0])
	fmt.Printf("2nd byte : %c\n", str[1])
	fmt.Printf("2nd rune : %s\n", str[1:3])

	runes := []rune(str)

	fmt.Printf("\t%d bytes\n", int(unsafe.Sizeof(runes[0]))*len(runes))

	fmt.Println("runes", string(runes))

}
