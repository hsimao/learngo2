package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	names := []string{}

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name)
			fmt.Println("names", names)
		}
	}
}