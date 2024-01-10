package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	AnimalsBiteForces = map[string]int{
		"Lion":         600,
		"Crocodile":    3500,
		"Great White":  4000,
		"Tiger":        1050,
		"Hyena":        1100,
		"Gorilla":      1300,
		"Kangaroo":     200,
		"Hippopotamus": 1825,
		"Polar Bear":   1235,
		"Wolf":         1200,
	}
)

type Animal struct {
	Name      string
	BiteForce int
}

func Sleep(a Animal, d time.Duration, ch chan Animal, wg *sync.WaitGroup) {
	time.Sleep(d)
	ch <- a
	wg.Done()
}

// A multi-threaded approach to sort a slice of items
// based on a provided sort key. In this case the BiteForce of an animal.
func SleepSortUnbuffered(animals []Animal) {
	fmt.Println("UnBuffered")
	sortCh := make(chan Animal)
	go func() {
		wg := &sync.WaitGroup{}
		for _, a := range animals {
			// Spawn a goroutine for every element
			wg.Add(1)
			go Sleep(a, time.Millisecond*time.Duration(a.BiteForce), sortCh, wg)
		}
		// Wait until all the goroutines are done with the task.
		wg.Wait()
		close(sortCh)
	}()

	// Items read off of the channel are always sorted
	// according to the sort key.
	for val := range sortCh {
		fmt.Println(val)
	}
	fmt.Println()
}

// A buffered channel implementation of the SleepSortUnbuffered
func SleepSortBuffered(animals []Animal) {
	fmt.Println("Buffered")
	sortCh := make(chan Animal, len(animals))
	// The following can now be blocking because we have a buffered
	// channel.
	wg := &sync.WaitGroup{}
	for _, a := range animals {
		// Spawn a goroutine for every element
		wg.Add(1)
		go Sleep(a, time.Millisecond*time.Duration(a.BiteForce), sortCh, wg)
	}
	// Wait until all the goroutines are done with the task.
	wg.Wait()
	close(sortCh)

	// Items read off of the channel are always sorted
	// according to the sort key.
	for val := range sortCh {
		fmt.Println(val)
	}
	fmt.Println()
}

func main() {
	animals := make([]Animal, 0)
	for name, biteForce := range AnimalsBiteForces {
		animals = append(animals, Animal{name, biteForce})
	}

	SleepSortUnbuffered(animals)
	SleepSortBuffered(animals)
}
