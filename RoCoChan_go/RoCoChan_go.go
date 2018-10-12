package main

import (
	"fmt"
	"sync"
	"time"
)

var l0 = make(chan int)
var l1 = make(chan int)
var u0 = make(chan int)
var u1 = make(chan int)
var wg sync.WaitGroup

func car() {
	for {
		load()
		for i := 0; i < C; i++ {
			<-l0
		}
		for i := 0; i < C; i++ {
			<-l1
		}
		run()
		unload()
		for i := 0; i < C; i++ {
			<-u0
		}
		for i := 0; i < C; i++ {
			<-u1
		}
	}
}

func load()   { fmt.Println("car ready to load") }
func run()    { fmt.Println("car runs") }
func unload() { fmt.Println("car ready to unload") }

func passenger() {
	l0 <- 1
	board()
	l1 <- 1
	u0 <- 1
	unboard()
	u1 <- 1
	wg.Done()
}

func board()   { fmt.Println("passenger boards") }
func unboard() { fmt.Println("passenger unboards") }

const (
	C           = 4
	nPassengers = 12000
)

func main() {
	start := time.Now()
	go car()
	wg.Add(nPassengers)
	for i := 0; i < nPassengers; i++ {
		go passenger()
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
