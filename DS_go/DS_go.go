package main

//source: github.com/soniakeys/LittleBookOfSemaphores/sem

/*
Variables:
1 servings = 0
2 mutex = Semaphore(1)
3 emptyPot = Semaphore(0)
4 fullPot = Semaphore(0)

Cook Code:
1 while True:
2 emptyPot.wait()
3 putServingsInPot(M)
4 fullPot.signal()

Savage Code:
while True:
2 mutex.wait()
3 if servings == 0:
4 emptyPot.signal()
5 fullPot.wait()
6 servings = M
7 servings -= 1
8 getServingFromPot()
9 mutex.signal()
10
11 eat()

*/

import (
	"fmt"
	"time"

	"github.com/soniakeys/LittleBookOfSemaphores/sem"
)

var (
	servings = 0
	meals    = 0
	mutex    = sem.NewChanSem(1, 1)
	emptyPot = sem.NewChanSem(0, 1)
	fullPot  = sem.NewChanSem(0, 1)
	start    = time.Now()
)

const M = 4
const nSavages = 6
const nPots = 3000

func savage(n int) {

	for i := 0; i < 3000; i++ {
		mutex.Wait()
		if servings == 0 {
			emptyPot.Signal()
			fullPot.Wait()
			servings = M
		}
		servings--
		getServingFromPot(n)
		mutex.Signal()
		eat(n)
	}
}

func getServingFromPot(n int) {
	fmt.Println("savage ", n, " gets serving from pot")
}

func eat(n int) {
	fmt.Println("savage ", n, " eats")
	fmt.Println(time.Since(start))
}

func main() {
	fmt.Println(nPots, " pot fillings:")

	for i := 1; i <= nSavages; i++ {
		go savage(i)
	}
	for j := 0; j <= 3000; j++ {
		emptyPot.Wait()
		if j == nPots {
			return
		}
		meals++
		putServingsInPot(M)
		fullPot.Signal()
	}
}

func putServingsInPot(m int) {
	fmt.Println("cook puts ", m, " servings in pot", meals)
}
