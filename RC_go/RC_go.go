package main

/*

River Crossing


Variables:
1 barrier = Barrier(4)
2 mutex = Semaphore(1)
3 hackers = 0
4 serfs = 0
5 hackerQueue = Semaphore(0)
6 serfQueue = Semaphore(0)
7 local isCaptain = False

Solution:
1 mutex.wait()
2 hackers += 1
3 if hackers == 4:
4 hackerQueue.signal(4)
5 hackers = 0
6 isCaptain = True
7 elif hackers == 2 and serfs >= 2:
8 hackerQueue.signal(2)
9 serfQueue.signal(2)
10 serfs -= 2
11 hackers = 0
12 isCaptain = True
13 else:
14 mutex.signal() # captain keeps the mutex
15
16 hackerQueue.wait()
17
18 board()
19 barrier.wait()
20
21 if isCaptain:
22 rowBoat()
23 mutex.signal() # captain releases the mutex

*/

/* http://williams.comp.ncat.edu/comp755/concurrentExamples1.pdf
semaphore mutex = 1, linux = 0, ms = 0;int numHackers = 0, numSerfs = 0;
hacker {
	p(mutex);
	numHackers++;
	if( numHackers == 4) {v(linux);
		v(linux); v(linux);numHackers -= 4;
		v(mutex);getOnBoard ();
	} else if(numHackers >= 2 && numSerfs >= 2) {v(linux);
		v(ms);
		v(ms);numHackers -= 2;
		numSerfs -= 2;v(mutex);
		getOnBoard ();
	} else {v(mutex);
		p(linux);
		getOnBoard ();
	}
}
Semaphore boardMutex = 1;
int numBoarded = 0;
getOnBoard() {
	p(boardMutex);
	numBoarded++;
	board();
	if (numBoarded == 4) {
		rowboat();
		numBoarded = 0;
	}
	v(boardMutex);
}

*/

import (
	"log"
	"sync"

	"github.com/soniakeys/LittleBookOfSemaphores/sem"
)

var (
	barrier     = sem.NewBarrier(4)
	mutex       = sem.NewChanSem(1, 1)
	hackers     = 0
	serfs       = 0
	hackerQueue = sem.NewChanSem(0, 1)
	serfQueue   = sem.NewChanSem(0, 1)
)

func hacker() {
	isCaptain := false
	mutex.Wait()
	hackers++
	if hackers == 4 {
		hackerQueue.SignalN(4)
		hackers = 0
		isCaptain = true
	} else if hackers == 2 && serfs >= 2 {
		hackerQueue.SignalN(2)
		serfQueue.SignalN(2)
		serfs -= 2
		hackers = 0
		isCaptain = true
	} else {
		mutex.Signal() // captain keeps the mutex
	}
	hackerQueue.Wait()
	board("hacker")
	barrier.Wait()
	if isCaptain {
		rowBoat("hacker")
		mutex.Signal() // captain releases the mutex
	}
}

func serf() {
	isCaptain := false
	mutex.Wait()
	serfs++
	if serfs == 4 {
		serfQueue.SignalN(4)
		serfs = 0
		isCaptain = true
	} else if serfs == 2 && hackers >= 2 {
		serfQueue.SignalN(2)
		hackerQueue.SignalN(2)
		hackers -= 2
		serfs = 0
		isCaptain = true
	} else {
		mutex.Signal() // captain keeps the mutex
	}
	serfQueue.Wait()
	board("serf")
	barrier.Wait()
	if isCaptain {
		rowBoat("serf")
		mutex.Signal() // captain releases the mutex
	}
}

func board(p string) {
	log.Print(p, " boards")
}

func rowBoat(p string) {
	log.Print(p, " rows")
	wg.Done()
}

const (
	nHackers = 6
	nSerfs   = 6
)

var wg sync.WaitGroup

func main() {
	wg.Add((nHackers + nSerfs) / 4)
	for i := 0; i < nHackers; i++ {
		go hacker()
	}
	for i := 0; i < nSerfs; i++ {
		go serf()
	}
	wg.Wait()
}
