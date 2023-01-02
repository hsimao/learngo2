package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	const (
		width     = 50
		height    = 10
		cellEmpty = ' '
		cellBall  = '⚽'
		maxFrames = 1200
		speed     = time.Second / 20
	)

	var (
		px   int
		py   int
		vx   = 1
		vy   = 1
		cell rune
	)

	board := make([][]bool, width)

	for column := range board {
		board[column] = make([]bool, height)
	}

	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		// 碰到邊界, 反彈
		if px <= 0 || px >= width-1 {
			// 變成負數
			vx *= -1
		}

		// 碰到邊界, 反彈
		if py <= 0 || py >= height-1 {
			// 變成負數
			vy *= -1
		}

		// 清除球
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		// 設置當前位置球
		board[px][py] = true

		buf := make([]rune, 0, width*height)
		// 優化, 讓 buf 變成 slice特性, append 的時候不會重新再分配創新的記憶體位置
		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}

			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		fmt.Print(string(buf))

		time.Sleep(speed)
	}
}
