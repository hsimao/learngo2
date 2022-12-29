package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	data := []float64{10, 25, 30, 50}

	copy(data, []float64{99, 100})

	// data = append(data[:0], []float64{1, 2, 3, 4, 5}...)

	// 不會指向到原本資料的 copy 用法
	// saved := make([]float64, len(data))
	// copy(saved, data)

	// 改用 append 也可達到相同效果, 更精簡
	saved := append([]float64(nil), data...)

	data[0] = 0

	s.Show("Probabilities (saved)", saved)
	s.Show("Probabilities (data)", data)

	fmt.Printf("Is it gonna rain? %.f%% change\n", (data[0]+data[1]+data[2]+data[3])/float64(len(data)))

}
