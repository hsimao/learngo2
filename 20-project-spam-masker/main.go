package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println("args", args)

	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	if len(args) != 1 {
		fmt.Println("gimme something to mask!")
		return
	}

	var (
		text = args[0]
		size = len(text)
		// buf  = make([]byte, 0, size)
		buf = []byte(text)

		in bool
	)

	// fmt.Println("buf", buf)
	fmt.Println("size", size)

	for i := 0; i < size; i++ {
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true
			// buf = append(buf, link...)
			i += nlink
		}

		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			// c = mask
			buf[i] = mask
		}

		// buf = append(buf, c)
	}
	fmt.Println(string(buf))
}
