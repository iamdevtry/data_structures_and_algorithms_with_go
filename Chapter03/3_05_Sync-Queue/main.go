package main

import (
	"fmt"
	"math/rand"
	"time"
)

// constants
const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

// Queue class
type Queue struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

// New method initializes queue
func (queue *Queue) New() {
	queue.message = make(chan int)
	queue.queuePass = make(chan int)
	queue.queueTicket = make(chan int)

	go func() {
		var message int
		for {
			select {
			case message = <-queue.message:
				switch message {
				case messagePassStart:
					queue.waitPass++
				case messagePassEnd:
					queue.playPass = false
				case messageTicketStart:
					queue.waitTicket++
				case messageTicketEnd:
					queue.playTicket = false
				}

				if queue.waitPass > 0 &&
					queue.waitTicket > 0 &&
					!queue.playPass &&
					!queue.playTicket {
					queue.playPass = true
					queue.playTicket = true
					queue.waitTicket--
					queue.waitPass--
					queue.queuePass <- 1
					queue.queueTicket <- 1
				}
			}
		}
	}()
}

// StartTicketIssue starts the ticket issue
func (queue *Queue) StartTicketIssue() {
	queue.message <- messagePassStart
	<-queue.queueTicket
}

// EndTicketIssue ends the ticket issue
func (queue *Queue) EndTicketIssue() {
	queue.message <- messageTicketEnd
}

// StartPass ends the Pass Queue
func (queue *Queue) StartPass() {
	queue.message <- messagePassStart
	<-queue.queuePass
}

// EndPass ends the Pass Queue
func (queue *Queue) EndPass() {
	queue.message <- messagePassEnd
}

// ticketIssue starts and ends the ticket issue
func ticketIssue(queue *Queue) {
	for {
		//Sleep up to 10 seconds
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		queue.StartTicketIssue()
		fmt.Println("Ticket issue starts")
		//Sleep up to 2 seconds
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticked issue ends")
		queue.EndTicketIssue()
	}
}

// passenger method starts and ends the pass Queue
func passenger(queue *Queue) {
	fmt.Println("starting the passenger Queue")
	for {
		fmt.Println("starting the processing")

		//Sleep up to 10 secs
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		queue.StartPass()
		fmt.Println(" Passenger starts")

		//Sleep up to 2 secs
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println(" Passenger ends")
		queue.EndPass()
	}
}

func main() {
	queue := Queue{}
	fmt.Println(queue)

	queue.New()
	fmt.Println(queue)

	for i := 0; i < 10; i++ {
		fmt.Println(i, "passenger in the queue")
		go passenger(&queue)
	}

	//close(queue.queuePass)
	for j := 0; j < 5; j++ {
		fmt.Println(j, "ticket issued in the Queue")
		go ticketIssue(&queue)
	}

	select {}
}
