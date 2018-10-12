package main

//source: github.com/soniakeys/LittleBookOfSemaphores/sem

/*
Variables:
1 isTobacco = isPaper = isMatch = False
2 tobaccoSem = Semaphore(0)
3 paperSem = Semaphore(0)
4 matchSem = Semaphore(0)
Pushers Code:
1 tobacco.wait()
2 mutex.wait()
3 if isPaper:
4 isPaper = False
5 matchSem.signal()
6 elif isMatch:
7 isMatch = False
8 paperSem.signal()
9 else:
10 isTobacco = True
11 mutex.signal()
Smoker with Tobacco:
1 tobaccoSem.wait()
2 makeCigarette()
3 agentSem.signal()
4 smoke()
*/

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/soniakeys/LittleBookOfSemaphores/sem"
)

var ( // Agent semaphores
	agentSem = sem.NewChanSem(1, 1)
	tobacco  = sem.NewChanSem(0, 1)
	paper    = sem.NewChanSem(0, 1)
	match    = sem.NewChanSem(0, 1)
)

var ( // Smoker semaphores
	isTobacco, isPaper, isMatch = false, false, false

	tobaccoSem = sem.NewChanSem(0, 1)
	paperSem   = sem.NewChanSem(0, 1)
	matchSem   = sem.NewChanSem(0, 1)
	mutex      = sem.NewChanSem(1, 1)
)

var provided int64
var wg sync.WaitGroup

const rounds = 1000

func main() {
	start := time.Now()
	wg.Add(rounds)
	go func() { // Agent A code
		for {
			agentSem.Wait()
			if atomic.AddInt64(&provided, 1) > rounds {
				return
			}
			fmt.Println("agent provides tobacco and paper")
			tobacco.Signal()
			paper.Signal()
		}
	}()
	go func() { // Agent B code
		for {
			agentSem.Wait()
			if atomic.AddInt64(&provided, 1) > rounds {
				return
			}
			fmt.Println("agent provides paper and a match")
			paper.Signal()
			match.Signal()
		}
	}()
	go func() { // Agent C code
		for {
			agentSem.Wait()
			if atomic.AddInt64(&provided, 1) > rounds {
				return
			}
			fmt.Println("agent provides tobacco and a match")
			tobacco.Signal()
			match.Signal()
		}
	}()
	go func() { // Pusher A
		for {
			tobacco.Wait()
			mutex.Wait()
			if isPaper {
				isPaper = false
				matchSem.Signal()
			} else if isMatch {
				isMatch = false
				paperSem.Signal()
			} else {
				isTobacco = true
			}
			mutex.Signal()
		}
	}()
	go func() { // Pusher B
		for {
			paper.Wait()
			mutex.Wait()
			if isTobacco {
				isTobacco = false
				matchSem.Signal()
			} else if isMatch {
				isMatch = false
				tobaccoSem.Signal()
			} else {
				isPaper = true
			}
			mutex.Signal()
		}
	}()
	go func() { // Pusher C
		for {
			match.Wait()
			mutex.Wait()
			if isPaper {
				isPaper = false
				tobaccoSem.Signal()
			} else if isTobacco {
				isTobacco = false
				paperSem.Signal()
			} else {
				isMatch = true
			}
			mutex.Signal()
		}
	}()
	go func() { // Smoker with tobacco
		for {
			tobaccoSem.Wait()
			makeCigarette("tobacco")
			agentSem.Signal()
			smoke("tobacco")
		}
	}()
	go func() { // Smoker with paper
		for {
			paperSem.Wait()
			makeCigarette("paper")
			agentSem.Signal()
			smoke("paper")
		}
	}()
	go func() { // Smoker with matches
		for {
			matchSem.Wait()
			makeCigarette("matches")
			agentSem.Signal()
			smoke("matches")
		}
	}()
	wg.Wait()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(time.Since(start))
}

func makeCigarette(s string) {
	fmt.Println("smoker with", s, "makes cigarette")
}

func smoke(s string) {
	fmt.Println("smoker with", s, "smokes")
	wg.Done()
}
