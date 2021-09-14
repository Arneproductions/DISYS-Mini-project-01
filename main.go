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
		for idx, val := range phils {
			fmt.Printf("Philosopher %d - %s\n", (idx + 1), val.GetStatus())
		}

		fmt.Println("=============")
		time.Sleep(400 * time.Millisecond)
	}
}

func initForks(amount int) []model.Fork {
	forks := make([]model.Fork, 0)

	for i := 0; i < amount; i++ {
		forks = append(forks, model.NewFork())
	}

	return forks
}

func initPhils(forks []model.Fork) []model.Philosopher {
	phils := make([]model.Philosopher, 0)

	for i := 0; i < len(forks); i++ {
		if i%len(forks) == 0 {
			phils = append(phils, model.NewPhilosopher(forks[len(forks)-1], forks[i]))
		} else {
			phils = append(phils, model.NewPhilosopher(forks[len(forks)-1], forks[i]))
		}
	}

	return phils
}
