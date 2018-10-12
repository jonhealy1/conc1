package main

//source: github.com/soniakeys/LittleBookOfSemaphores/sem

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const M = 4
const nSavages = 6
const nPots = 3000

var pot int64
var wake = make(chan chan int)
var wg sync.WaitGroup

func savage(n int) {
	watch := make(chan int)
	var s int64
	for {
		//fmt.Println("savage ", n, " out hunting/gathering")
		//time.Sleep(time.Duration(rand.Intn(1e7)))
		//fmt.Println("savage ", n, " hungry, returns to camp")
		for {
			if s = atomic.AddInt64(&pot, -1); s >= 0 {
				break
			}
			//fmt.Println("savage ", n, " finds pot empty, yells for cook")
			wake <- watch
			// log.Print("savage ", n, " waiting to see full pot")
			<-watch
			// log.Print("savage ", n, " sees servings in pot")
		}
		fmt.Println("savage ", n, " eats (", s, " servings in pot)")
		//time.Sleep(1e6)
		wg.Done()
	}
}

func main() {
	start := time.Now()
	wg.Add(nPots * M)
	for i := 1; i <= nSavages; i++ {
		go savage(i)
	}
	awake := false
	waiting := make([]chan int, 0, nSavages)
	var stew <-chan time.Time
	for pots := 0; ; {
		select {
		case s := <-wake:
			// simulation code: "register" savage watching
			waiting = append(waiting, s)
			// cook wakes if not already awake
			if !awake {
				awake = true
				// cook must see pot empty
				if atomic.LoadInt64(&pot) > 0 {
					//fmt.Println("cook grumbles, goes back to sleep")
					awake = false
					// simulation code: savage must now notice same non-empty
					// pot that cook sees, and stop waiting.
					s <- 1
					waiting = waiting[:len(waiting)-1]
				} else {
					fmt.Println("cook awake, starts cooking")
					stew = time.After(1e5)
				}
			}
		case <-stew:
			pots++
			fmt.Println("cook puts", M, "servings in pot") //, pots, "pots cooked")
			atomic.StoreInt64(&pot, M)

			// simulation code: pot filling event observed by waiting savages
			for _, s := range waiting {
				s <- 1
			}
			waiting = waiting[:0]
			if pots == nPots {
				fmt.Println("cook leaves")
				wg.Wait() // simulation waits for savages to finish eating
				fmt.Println("simulation ends")
				//time.Sleep(100 * time.Millisecond)
				fmt.Println(time.Since(start))
				return
			}
			//log.Println("cook sleeps")
			awake = false
		}
		//time.Sleep(100 * time.Millisecond)

	}

}
