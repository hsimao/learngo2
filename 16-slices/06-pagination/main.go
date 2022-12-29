package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	const pageSiae = 4

	l := len(items)

	for from := 0; from < l; from += pageSiae {
		to := from + pageSiae

		if to > l {
			to = l
		}

		currentPage := items[from:to]

		head := fmt.Sprintf("Page #%d", (from/pageSiae)+1)

		s.Show(head, currentPage)
	}

}
