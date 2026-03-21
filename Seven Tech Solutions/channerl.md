<!-- Q. What are the different types of channels in Go, and where can we use them?
Ans. 
There are several channel types in Go: 

unbuffered channels (synchronous, send blocks until receive), 

buffered channels (asynchronous up to capacity),

directional channels (chan<- send-only, <-chan receive-only), 

nil channels (never ready, useful to disable cases in select), 
and closed channels (reads yield zero value immediately; writes panic). 

Each has predictable behavior with select,
 timeouts, and goroutine coordination.

Use cases: unbuffered for tight synchronization or handoff; buffered for decoupling producer/consumer 
rates or burst handling; directional for API clarity and safety; nil channels to dynamically enable/disable
 select branches; closed channels to broadcast completion or cancellation. 
 
 Common patterns: pipelines, worker pools, 
 fan-in/fan-out, and implementing timeouts/cancellation without shared mutexes. -->


<!-- 
 1. Unbuffered Channel (Synchronous)
Behavior
No capacity
send blocks until a receiver is ready
receive blocks until a sender is ready
Used for strict synchronization (handoff)

Example

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10 // blocks until main receives
	}()

	val := <-ch // blocks until goroutine sends
	fmt.Println(val)
} -->

goroutine send ---> blocks ---> main receives ---> continues


Use Cases

Goroutine coordination
Ensuring order of execution
Request-response patterns




<!-- 2. Buffered Channel (Asynchronous)

Behavior
Has capacity (queue)
send blocks only when buffer is full
receive blocks only when buffer is empty -->

<!-- package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1 // ok
	ch <- 2 // ok
	// ch <- 3 // would block (buffer full)

	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
}

Producer → buffer → Consumer

Use Cases

Producer-consumer problems
Rate decoupling
Burst handling

 -->




<!-- 3. Directional Channels

Behavior

Restrict channel usage:
chan<- T → send-only
<-chan T → receive-only
Improves type safety & API clarity

package main

import "fmt"

func sender(ch chan<- int) {
	ch <- 100
}

func receiver(ch <-chan int) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int)

	go sender(ch)
	receiver(ch)
}


Use Cases

Public APIs
Prevent misuse of channels
Clear intent in large systems -->



<!-- 4. Nil Channel

Behavior
A channel with zero value (var ch chan int)
Send/receive blocks forever
Useful in select to disable cases dynamically

Example (dynamic select control)
package main

import "fmt"

func main() {
	var ch chan int // nil

	select {
	case ch <- 1:
		fmt.Println("sent")
	default:
		fmt.Println("blocked because channel is nil")
	}
} -->


<!-- var ch chan int

if condition {
	ch = make(chan int)
}

select {
case ch <- 10:
	// only runs if ch != nil
default:
} 

Use Cases
Enable/disable select cases dynamically
State-based flow control

-->