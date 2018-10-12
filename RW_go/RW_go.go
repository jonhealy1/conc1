/*
Pseudo Code - Little Book of Semaphores

#Variables:
int readers = 0
mutex = Semaphore(1)
roomEmpty = Semaphore(1)

#Writers:
roomEmpty.wait()
critical section
roomEmpty.signal()

#Readers:
mutex.wait()
readers++
if readers == 1
	roomEmpty.wait()
mutex.signal()
critical section
mutex.wait()
readers—
if readers == 0
	roomEmpty.signal()
mutex.signal()
*/

/* The majority of this code comes from github.com/soniakeys/LittleBookOfSemaphores */

package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/soniakeys/LittleBookOfSemaphores/sem"
)

var (
	readers   = 0
	mutex     = sem.NewChanSem(1, 1)
	roomEmpty = sem.NewChanSem(1, 1)
	b         bytes.Buffer // protected by mutex
)

var wg sync.WaitGroup

func writer(nw int) {
	roomEmpty.Wait()
	fmt.Println("writer", nw, "writes")
	b.WriteString("w")
	roomEmpty.Signal()
	wg.Done()
}

func reader(nr int) {
	mutex.Wait()
	readers++
	if readers == 1 {
		roomEmpty.Wait() // first in locks
	}
	mutex.Signal()

	fmt.Println("reader", nr, "sees", b.Len(), "bytes")

	mutex.Wait()
	readers--
	if readers == 0 {
		roomEmpty.Signal() // last out unlocks
	}
	mutex.Signal()
	wg.Done()
}

const nw = 1000
const nr = 1000

func main() {
	start := time.Now()
	wg.Add(nw + nr)
	for i := 1; i <= 1000; i++ {
		go writer(i)
		go reader(i)
	}
	wg.Wait()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(time.Since(start))
}
