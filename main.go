package main

import (
	"fmt"
	"dk.disys.phil/model"
)

func main() {
	left := model.NewFork()
	right := model.NewFork()
	p := model.NewPhilosopher(right, left)

	fmt.Println(p)
}
