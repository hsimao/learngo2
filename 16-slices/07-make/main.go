package main

import (
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 10

	tasks := []string{"jump", "run", "read"}

	// 使用 make 可以優化, 每次 append 的時候將不會重新創建新的記憶體位置
	upTasks := make([]string, 0, cap(tasks))

	// 每次 append 都會創建新的記憶體位置
	// upTasks := []string{}

	s.Show("upTasks", upTasks)

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}
}
