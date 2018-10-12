package main

//source: github.com/soniakeys/LittleBookOfSemaphores/sem

/*
Variables and Definitions :
1 def left(i): return i
2 def right(i): return (i + 1) % 5
1 forks = [Semaphore(1) for i in range(5)]
Solution:
1 def get_forks(i):
2 footman.wait()
3 fork[right(i)].wait()
4 fork[left(i)].wait()
5
6 def put_forks(i):
7 fork[right(i)].signal()
8 fork[left(i)].signal()
9 footman.signal()
*/

import (
	"fmt"
	"sync"
	"time"

	"github.com/soniakeys/LittleBookOfSemaphores/sem"
)

var (
	footman = sem.NewChanSem(4, 4)
	fork    [5]sem.ChanSem
)

func init() {
	for i := range fork {
		fork[i] = sem.NewChanSem(1, 1)
	}
}

var wg sync.WaitGroup

const nBites = 1

func ph(n int) {
	for b := 1; b <= nBites; b++ {
		think(n)
		get_forks(n)
		eat(n, b)
		put_forks(n)
	}
	wg.Done()
}

func think(n int) {
	fmt.Println("phil", n, "thinking")
	//time.Sleep(time.Duration(rand.Intn(1e8)))
}

func eat(n, b int) {
	fmt.Println("phil", n, "eats bite #", b)
	//time.Sleep(time.Duration(rand.Intn(1e8)))
}

func get_forks(i int) {
	//fmt.Println("philosopher", i, "wants to sit and eat")
	footman.Wait()
	//fmt.Println("philosopher", i, "seated, looking for forks")
	fork[right(i)].Wait()
	fmt.Println("phil", i, "picked up right fork")
	fork[left(i)].Wait()
	fmt.Println("phil", i, "picked up left fork")
}
func put_forks(i int) {
	fork[right(i)].Signal()
	fork[left(i)].Signal()
	footman.Signal()
	fmt.Println("phil", i, "full, returns forks")
}

func right(i int) int { return i }
func left(i int) int  { return (i + 1) % 5 }

func main() {
	start := time.Now()
	wg.Add(1000)
	for j := 0; j < 200; j++ {
		for i := 0; i < 5; i++ {
			go ph(i)
		}
	}
	wg.Wait()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(time.Since(start))
}
