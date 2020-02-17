package main

import (
	"fmt"
	"sync"
)

// ChopS represent chopstick
type ChopS struct{ sync.Mutex }

// Philo represents Philosophers
type Philo struct {
	number          int
	leftCS, rightCS *ChopS
}

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("starting to eating %d \n", p.number)
		fmt.Printf("finishes eating %d \n", p.number)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	Csticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		Csticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 4; i++ {
		philos[i] = &Philo{
			i,
			Csticks[i],
			Csticks[(i+1)%5]}
	}

	philos[4] = &Philo{
		4,
		Csticks[0],
		Csticks[4]}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg)
	}
	wg.Wait()

}
