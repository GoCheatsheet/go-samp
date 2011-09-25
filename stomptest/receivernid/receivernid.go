// First gostomp demo

package main

import (
	"fmt" //
	"os"
  "net"
	"runtime"
  "stomp"
	"strings"
	"sync"
	"time"
)

var printMsgs bool = true
var printHdrs bool = true
var wg sync.WaitGroup
var	qname = "/queue/gostomp.srpub"
var	mq = 500
var host = "localhost"
var hap = host + ":"

var incrCtl sync.Mutex
var numRecv int

func recMessages(c *stomp.Conn, q string) {

	var error os.Error

	fmt.Printf("Start for q: %s\n", q)

	// Receive phase
  headers := stomp.Header{"destination": q} // no ID here.  1.1 library should provide
	fmt.Printf("qhdrs: %v\n", headers)
	sc, error := c.Subscribe(headers)
	if error != nil {
		// Handle error properly
		fmt.Printf("sub error: %v\n", error)
	}
	first := true
	firstSub := ""
	for input := range sc {
    inmsg := string(input.Message.Data)
    if printHdrs {
  		fmt.Println("queue:", q, "Next Receive: ", input.Message.Header)
    }
    if printMsgs {
  		fmt.Println("queue:", q, "Next Receive: ", inmsg)
    }

		firstSub = input.Message.Header["subscription"]
		if first {
			if firstSub == "" {
				panic("first subscription header is empty")
			}
			fmt.Println("queue:", q, "FirstSub: ", firstSub)
			first = false
		} else {
			if firstSub != input.Message.Header["subscription"] {
				panic(firstSub + " / " + input.Message.Header["subscription"])
			}
		}

		time.Sleep(1e9 / 100)	// Crudely simulate message processing

		incrCtl.Lock()
		numRecv++
		incrCtl.Unlock()

		if strings.HasPrefix(inmsg, "***EOF***") {
			fmt.Println("queue:", q, "FirstSub:", firstSub, "goteof")
			break
		}
		if !strings.HasPrefix(inmsg, q) {
			fmt.Printf("bad prefix: %v, %v\n", q, inmsg)
			panic("bad prefix ....")
		}
		// Poll for adhoc errors
		select {
			case v := <- c.Stompdata:
				fmt.Printf("frameError: %v\n", v.Message)
				fmt.Printf("frameError: [%v] [%v]\n", q, firstSub) 
			default:
				fmt.Println("Nothing to show")
		}

	}

	// headers["subscription"] = firstSub	// Add ID to unsubscribe
	error = c.Unsubscribe(headers)
	if error != nil {
		// Handle error properly
		fmt.Printf("unsub error: %v\n", error)
	}

	wg.Done()
}

func main() {
	fmt.Println("Start...")

  // create a net.Conn, and pass that into Connect
	nc, error := net.Dial("tcp", hap + os.Getenv("STOMP_PORT"))
	if error != nil {
		// Handle error properly
	}

  // Connect
	ch := stomp.Header{"login": "getter", "passcode": "recv1234"}

	//
	ch["accept-version"] = "1.1"
	ch["host"] = host

	c, error := stomp.Connect(nc, ch)
	if error != nil {
		panic(error)
	}

	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		wg.Add(1)
		go recMessages(c, qname + qn)
	}
	wg.Wait()

	fmt.Printf("Num received: %d\n", numRecv)

  // Disconnect
  nh := stomp.Header{}
	error = c.Disconnect(nh)
	if error != nil {
		fmt.Printf("discerr %v\n", error)
	}

	fmt.Println("done nc.Close()")
	nc.Close()

/*
	fmt.Println("start sleep")
	time.Sleep(1e9 / 10)	// 100 ms
	fmt.Println("end sleep")
*/

	ngor := runtime.Goroutines()
	fmt.Printf("egor: %v\n", ngor)

	select {
		case v := <- c.Stompdata:
			fmt.Printf("frame2: %s\n", v.Message.MsgFrame)
			fmt.Printf("header2: %v\n", v.Message.Header)
			fmt.Printf("data2: %s\n", string(v.Message.Data))
		default:
			fmt.Println("Nothing to show")
	}
/*
	if ngor > 1 {
		panic("too many gor")
	}
*/
	fmt.Println("End... ngor:", mq)
}
