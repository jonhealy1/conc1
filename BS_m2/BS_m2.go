package main

import (
	"fmt"
	"sync"
	"time"
)

const shopCap = 4000 // number of customers allowed in shop at the same time
const nCust = 6000   // total number of customers that will come to the shop

var (
	inShop = 0        // number of customers currently in shop
	mutex  sync.Mutex // protects inShop
)

// customer sends his customer number to enter the barber room
var barberRoom = make(chan int)

// barber sends dummy value when he's done cutting
var cutDone = make(chan int)

// counts customers as they leave shop
var wg sync.WaitGroup

func main() {
	start := time.Now()
	wg.Add(nCust)
	go barber()
	for c := 1; c <= nCust; c++ {
		go customer(c)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

func barber() {
	for {
		fmt.Println("barber sleeping")
		c := <-barberRoom
		fmt.Println("barber wakes, cuts customer", c, "hair")
		cutDone <- 1
	}
}

func customer(c int) {

	mutex.Lock()
	fmt.Println("customer", c, "arrives,", inShop,
		"customers")
	if inShop == shopCap {
		mutex.Unlock()
		fmt.Println("customer", c, ",shop full, leaves")
		wg.Done()
		return
	}
	inShop++
	fmt.Println("customer", c, "waits")
	mutex.Unlock()
	barberRoom <- c
	time.Sleep(1e6)
	fmt.Println("customer", c, "getting hair cut")
	<-cutDone
	mutex.Lock()
	inShop--
	mutex.Unlock()
	fmt.Println("customer", c, "leaves with hair cut")
	wg.Done()

}
