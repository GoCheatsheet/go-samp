// Demo a wait group synchronization of go routines.

package main

import (
	"fmt"  //
	"time" //
	"sync" //
)

var wg sync.WaitGroup

type gdata struct {
	id            string
	loops, sltime int
}

var runData = []gdata{gdata{"1", 6, 2},
	gdata{"2", 5, 1},
	gdata{"3", 8, 3},
}

func called(cd gdata) {
	for i := 0; i < cd.loops; i++ {
		wnum := i + 1
		fmt.Printf("id %s, waitnum: %d\n", cd.id, wnum)
		err := time.Sleep(int64(cd.sltime) * 1e9)
		if err != nil {
			// ??? 
		}
	}
	fmt.Println(cd.id, "is done")
	wg.Done()
}

func main() {
	fmt.Println("Start...")
	//
	for _, curgd := range runData {
		wg.Add(1)
		go called(curgd)
	}
	//
	fmt.Println("Starting main wait")
	wg.Wait()
	fmt.Println("End...")
}
