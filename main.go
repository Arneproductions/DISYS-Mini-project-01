package main

import (
	"fmt"
	"time"

	"dk.disys.phil/model"
)

func main() {
	forks := initForks(5)
	phils := initPhils(forks)

	for {
		for _, val := range phils {
			fmt.Println(val.GetStatus())
		}

		fmt.Println("==Forks==")
		for _, val := range forks {
			fmt.Println(val.GetStatus())
		}
		fmt.Println("=============")

	time.Sleep(400 * time.Millisecond)
	}
}

func initForks(amount int) []*model.Fork {
	forks := make([]*model.Fork, 0)

	for i := 0; i < amount; i++ {
		forks = append(forks, model.NewFork(i))
	}

	return forks
}

func initPhils(forks []*model.Fork) []model.Philosopher {
	phils := make([]model.Philosopher, 0)

	for i := 0; i < len(forks); i++ {
		if i == 0 {
			phils = append(phils, model.NewPhilosopher(i, forks[len(forks)-1], forks[i]))
		} else {
			phils = append(phils, model.NewPhilosopher(i, forks[i], forks[i-1]))
		}
	}

	return phils
}
