package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

// 數字佔位空間
type Placeholder [5]string

type Digits [10]Placeholder

type Clock [8]Placeholder

func printClock(digits Digits) {
	now := time.Now()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()

	// 冒號 :
	colon := Placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	// 時鐘數字
	clock := Clock{
		// hour 第一位號碼
		digits[hour/10],
		// hour 第二位號碼
		digits[hour%10],
		// 冒號
		colon,
		// min 第一位號碼
		digits[min/10],
		// min 第二位號碼
		digits[min%10],
		// 冒號
		colon,
		// sec 第一位號碼
		digits[sec/10],
		// sec 第二位號碼
		digits[sec%10],
	}

	// 五行
	for line := range clock[0] {
		// 每一行印出每位數字的當前 line 字串
		for index, digit := range clock {
			next := clock[index][line]

			if digit == colon && sec%2 == 0 {
				next = "   "
			}
			fmt.Print(next, "  ")
		}
		fmt.Println()
	}
	fmt.Println()

	time.Sleep(time.Second)
}

func main() {

	// 0~9 擺放數字 placeholder 陣列空間
	zero := Placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := Placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two := Placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := Placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := Placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := Placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := Placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := Placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := Placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := Placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := Digits{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()
	for {
		screen.MoveTopLeft()
		printClock(digits)
	}

}
