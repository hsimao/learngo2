package main

import s "github.com/inancgumus/prettyslice"

func main() {

	var todo []string

	todo = append(todo, "sing")
	todo = append(todo, "run", "code", "play")

	tomorrow := []string{"see mon", "learn go"}
	todo = append(todo, tomorrow...)

	s.Show("todo", todo)
}
