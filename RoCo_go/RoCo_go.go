package main

//source: github.com/soniakeys/LittleBookOfSemaphores/sem

import (
	"fmt"
	"sync"
	"time"

	"github.com/soniakeys/LittleBookOfSemaphores/sem"
)

var (
	mutex        = sem.NewChanSem(1, 1)
	mutex2       = sem.NewChanSem(1, 1)
	boarders     = 0
	unboarders   = 0
	boardQueue   = sem.NewChanSem(0, 1)
	unboardQueue = sem.NewChanSem(0, 1)
	allAboard    = sem.NewChanSem(0, 1)
	allAshore    = sem.NewChanSem(0, 1)
	//passNum      = 0
)

func car() {
	for {
		load()
		boardQueue.SignalN(C)
		allAboard.Wait()
		run()
		unload()
		unboardQueue.SignalN(C)
		allAshore.Wait()
	}
}

func load()   { fmt.Println("car ready to load") }
func run()    { fmt.Println("car runs") }
func unload() { fmt.Println("car ready to unload") }

func passenger() {
	boardQueue.Wait()
	board()
	mutex.Wait()
	boarders++
	if boarders == C {
		allAboard.Signal()
		boarders = 0
	}
	mutex.Signal()
	unboardQueue.Wait()
	unboard()
	mutex2.Wait()
	unboarders++
	if unboarders == C {
		allAshore.Signal()
		unboarders = 0
	}
	mutex2.Signal()
	wg.Done()
}

func board()   { fmt.Println("passenger boards") }
func unboard() { fmt.Println("passenger unboards") }

var wg sync.WaitGroup

const (
	C           = 4
	nPassengers = 12000
)

func main() {
	start := time.Now()
	go car()
	wg.Add(nPassengers)
	for i := 0; i < nPassengers; i++ {
		//passNum++
		go passenger()
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
