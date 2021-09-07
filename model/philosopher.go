package model

type philosopher struct {
	right fork 
	left fork
	eating bool
	timesEaten int 
	in chan int
	out chan int
}

func newPhilosopher (right fork, left fork) philosopher{
	p := philosopher{right: right, left: left}
	p.eating = false
	p.timesEaten = 0
	p.in = make(chan int)
	p.out = make(chan int)
	return p
}