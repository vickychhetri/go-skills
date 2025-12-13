package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*

📌 Rule of thumb
Use atomic → counters, flags, stats
Use mutex → complex shared state

In Go, the sync/atomic package is used for low-level, lock-free concurrency control.
It lets multiple goroutines safely read/write shared variables without using mutexes.

1. What is the atomic package?
sync/atomic provides atomic memory primitives like:
Atomic load
Atomic store
Atomic add
Atomic compare-and-swap (CAS)
These operations are:
✅ Thread-safe
✅ Lock-free
✅ Very fast
❌ Limited to simple operations only
📌 Atomic = the operation happens completely or not at all (no partial state visible to other goroutines).


2. Why atomic? (Problem it solves)
❌ Without atomic (Race condition)

var counter int64
go func() {
	counter++ // NOT safe
}()

Multiple goroutines can:
Read same value
Increment it
Write back → lost updates

With atomic
atomic.AddInt64(&counter, 1)

Now:
Increment is safe
No race condition
No mutex overhead


3. Common atomic functions

Function					Purpose
atomic.LoadInt64			Safely read value
atomic.StoreInt64			Safely write value
atomic.AddInt64				Add value atomically
atomic.CompareAndSwapInt64	CAS operation
atomic.SwapInt64			Replace value atomically
*/

func main() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 10000000000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)

		}()
	}

	wg.Wait()
	fmt.Println("counter: ", counter)
}

/*
Atomic Load & Store Example
var status int32

// Writer
atomic.StoreInt32(&status, 1)

// Reader
if atomic.LoadInt32(&status) == 1 {
	fmt.Println("Service is running")
}

Useful when many goroutines read, few write.



Compare-And-Swap (CAS) – Most Important
Mental Model
Think of CAS as:
“I’ll change this value only if nobody else touched it first.”

What is CAS?
Update value only if it matches expected value.
func CompareAndSwapInt64(addr *int64, old, new int64) bool

var initialized int32
func initOnce() {
	if atomic.CompareAndSwapInt32(&initialized, 0, 1) {
		fmt.Println("Initialized only once")
	}
}
*/

/*
Real-World Use Cases

1️. Request Counter (API / Microservices)
var requests int64

func handler() {
	atomic.AddInt64(&requests, 1)
}


Used in:
Metrics
Monitoring
Rate limiting





2. Feature Flags
var featureEnabled int32
func enableFeature() {
	atomic.StoreInt32(&featureEnabled, 1)
}

func isEnabled() bool {
	return atomic.LoadInt32(&featureEnabled) == 1
}



3. Graceful Shutdown Flag
var shuttingDown int32

func shutdown() {
	atomic.StoreInt32(&shuttingDown, 1)
}

func worker() {
	for {
		if atomic.LoadInt32(&shuttingDown) == 1 {
			return
		}
	}
}


4. Rate Limiter Counter (IP-wise basic logic)
var hits int64

func allowRequest() bool {
	if atomic.AddInt64(&hits, 1) > 100 {
		return false
	}
	return true
}

5. Stats Collection (Prometheus-like)
var success int64
var failure int64

atomic.AddInt64(&success, 1)
atomic.AddInt64(&failure, 1)


sync/atomic provides low-level, lock-free primitives for safe concurrent access to shared variables.
It is faster than mutex but limited to simple operations like counters, flags, and CAS-based state changes.


Atomic Package Key Points (Remember)

✔ Lock-free
✔ Faster than mutex
✔ Used for counters, flags, stats
✔ CAS is core concept
✔ Hard to debug misuse
✔ Avoid complex logic



*/
