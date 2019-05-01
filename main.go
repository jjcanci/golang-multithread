package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	log.Println("Start")
	rand.Seed(time.Now().UTC().UnixNano())

	respond := make(chan string, 5)
	var wg sync.WaitGroup

	wg.Add(5)
	go doSomething(respond, &wg, "ns1.nameserver.com")
	go doSomething(respond, &wg, "ns2.nameserver.com")
	go doSomething(respond, &wg, "ns3.nameserver.com")
	go doSomething(respond, &wg, "ns4.nameserver.com")
	go doSomething(respond, &wg, "ns5.nameserver.com")

	log.Println("Wait for response...")
	wg.Wait()
	close(respond)
	log.Println("Ended wait!")

	for queryResp := range respond {
		log.Printf("Got Response: %s\n", queryResp)
	}

	/*select {
	case queryResp := <-respond:
		log.Printf("Sent query:\t\t %s\n", query)
		log.Printf("Got Response:\t\t %s\n", queryResp)

	case <-time.After(2 * time.Second):
		fmt.Printf("A timeout occurred for query:\t\t %s\n", query)
	}*/

	log.Println("End")
}

func doSomething(respond chan<- string, wg *sync.WaitGroup, ns string) {
	defer wg.Done()
	log.Printf("Wait for %s...", ns)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	respond <- fmt.Sprintf("%s responded on %s", ns, time.Now().Format(time.RFC1123))
}
