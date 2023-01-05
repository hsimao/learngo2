package main

import (
	"fmt"
	"unsafe"
)

func main() {
	hello := "hello"
	dump(hello)
	dump("hello")
	dump("hello!")
	dump(hello[0:1])

	for i := range hello {
		dump(hello[i : i+1])
	}

	// 轉換單位會產生新的指針, 創建新的儲存位置
	dump(string([]byte(hello)))
	dump(string([]rune(hello)))
}

type StringHeader struct {
	pointer uintptr
	length  int
}

func dump(s string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&s))

	fmt.Printf("%q: %+v\n", s, ptr)
}
