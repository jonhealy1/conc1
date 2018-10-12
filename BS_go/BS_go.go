package main

/*
Variables:
1 customers = 0
2 mutex = Semaphore(1)
3 customer = Semaphore(0)
4 barber = Semaphore(0)

Customer Code:
1 mutex.wait()
2 if customers == n+1:
3 mutex.signal()
4 balk()
5 customers += 1
6 mutex.signal()
7
8 customer.signal()
9 barber.wait()
10 getHairCut()
11
12 mutex.wait()
13 customers -= 1
14 mutex.signal()

Barber Code:
1 customer.wait()
2 barber.signal()
3 cutHair()
*/

import (
	"fmt"
	"sync"
	"time"

	"github.com/soniakeys/LittleBookOfSemaphores/sem"
)

var (
	n            = 4000
	customers    = 0
	mutex        = sem.NewChanSem(1, 1)
	customer     = sem.NewChanSem(0, 1)
	barber       = sem.NewChanSem(0, 1)
	customerDone = sem.NewChanSem(0, 1)
	barberDone   = sem.NewChanSem(0, 1)
)

var wg sync.WaitGroup

const nCust = 6000

func main() {
	start := time.Now()
	wg.Add(nCust)
	go barberFunc()
	for i := 1; i <= nCust; i++ {
		go customerFunc(i)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

func barberFunc() {
	for {
		fmt.Println("barber sleeping")

		customer.Wait()
		barber.Signal()

		// cutHair ()
		fmt.Println("barber cutting hair")

		customerDone.Wait()
		barberDone.Signal()
	}
}

func customerFunc(c int) {

	mutex.Wait()
	fmt.Println("customer ", c, " arrives,", customers,
		"customers")
	if customers == n {
		mutex.Signal()
		fmt.Println("customer ", c, " shop full, leaves")
		wg.Done()
		balk()
	}
	customers++
	fmt.Println("customer ", c, " waits")
	mutex.Signal()

	customer.Signal()
	barber.Wait()

	// getHairCut ()
	fmt.Println("customer ", c, " gets hair cut")

	customerDone.Signal()
	barberDone.Wait()

	mutex.Wait()
	customers--
	mutex.Signal()
	fmt.Println("customer ", c, " leaves with hair cut")
	//time.Sleep(100 * time.Millisecond)

	wg.Done()

}

func balk() {
	select {}
}
