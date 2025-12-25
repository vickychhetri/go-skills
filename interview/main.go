package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3}
	y := (x)
	y[0] = 100
	fmt.Println(x)

	// 	👉 What will be the output and why?
	// Explain slice internals (pointer, length, capacity).
	//[100 2 3]
	// 	Why this happens
	// In Go, a slice is not the data itself.
	// A slice is a small descriptor that points to an underlying array.

	// ##############################################
	// Q2️⃣
	// What is the difference between:
	// make() and new()
	// nil slice vs empty slice
	// 👉 When would you prefer one over the other in production code?

	p := new(int)
	fmt.Println(p)
	fmt.Println(*p)
	/**
	Allocates zeroed memory for type T
	Returns a pointer: *T
	Does not initialize internal structure

	p := new([]int)
	fmt.Println(p)   // pointer
	fmt.Println(*p)  // nil slice
	Rarely used for slices, maps, channels.
	*/

	/*
		make(T, ...)
			s := make([]int, 3)

			Used only for slices, maps, channels
			Allocates and initializes internal data
			Returns the value, not a pointer
			Example:
			s := make([]int, 3)
			fmt.Println(s) // [0 0 0]
	*/

	/*
			Production rule
		Use make() for slices, maps, and channels.
		Avoid new() unless you explicitly need a pointer to a zero value.
	*/

	/**

		Slice header looks like:

	Data = non-nil pointer (points to a zero-size array)
	Len  = 0
	Cap  = 0
	Nil slice has a nil pointer; empty slice has a non-nil pointer.
	*/

	/*
				Explain interface internals in Go.
					What is a nil interface?
					What happens internally when you assign a concrete type to an interface?
					Why does this code panic?
					var r io.Reader
					fmt.Println(r == nil)

		An interface value is NOT just a pointer.
		Internally, it is two things together.
		Conceptually:
		type iface struct {
		    itab *itab        // type information (what concrete type?)
		    data unsafe.Pointer // pointer to the actual value
		}

		So every interface value contains:
		Type information (what concrete type is stored)
		Data pointer (where the actual value lives)

			What is a nil interface?

			A nil interface means both fields are nil:

			itab = nil
			data = nil

			Example:
			var r io.Reader

			At this point:
				No concrete type
				No data

				So:
				r == nil // true
				This is a true nil interface.

				Key rule (MEMORIZE THIS)

				An interface is nil only if BOTH its type and value are nil.

				itab		data	interface == nil
				nil			nil	✅ true
				non-nil		nil	❌ false
				non-nil		non-nil	❌ false
	*/

	/*
				Round 2: Concurrency & Goroutines (Very Important)
					Q4️⃣

					Explain how goroutines are scheduled.

					1️⃣ How goroutines are scheduled (big picture)
						Go uses a user-space scheduler inside the runtime.
						Key ideas:
						Goroutines are much lighter than OS threads
						Scheduling is cooperative + preemptive (since Go 1.14+)
						The runtime multiplexes many goroutines onto fewer OS threads
						👉 You do not control goroutines directly — the runtime does.

		What is M, P, G model?
		2️⃣ The M-P-G model (memorize this)
		Go’s scheduler is built around three entities:
		🔹 G — Goroutine
		A goroutine

		Contains:
		stack
		instruction pointer
		state (running, waiting, runnable)
		Think:
		“What to run”

		🔹 M — Machine (OS thread)
		An actual OS thread
		Executes Go code or syscalls
		Think:
		“Where code runs”

		🔹 P — Processor (logical scheduler)
		NOT an OS core
		Holds:
		run queue of goroutines
		scheduler state
		Required to run Go code
		Think:
		“Permission to run Go code”
		📊 Relationship
		G (goroutine) ── runs on ──> M (OS thread)
		                    │
		                    needs
		                    ▼
		               P (processor)
		Key rule:
		An M must have a P to execute Go code.



		How is it different from OS threads?
		What happens when a goroutine blocks on I/O?

		Go uses an M-P-G scheduler where many goroutines (G) are multiplexed onto fewer OS threads (M),
		controlled by logical processors (P).
		 Blocking I/O does not block the entire thread — the runtime detaches the P
		 and keeps running other goroutines.
	*/

	/*
			Q6️⃣

		When would you use:

		sync.Mutex

		sync.RWMutex

		atomic package

		channels instead of mutex

		👉 Give real production use cases.

			Start with Mutex
			Switch to RWMutex only if reads dominate
			Use atomic only when profiling proves it matters
			Use channels when modeling workflows, not data protection


			Quick Decision 				Table

			Situation					Best Tool
			Simple shared state			sync.Mutex
			Many readers, few writers	sync.RWMutex
			Counters / flags			atomic
			Work distribution			channels
			Event pipelines				channels
			High-performance hot path	atomic or Mutex
			Complex shared logic		Mutex
	*/

	xy := 10
	defer fmt.Println(xy)
	xy = 20
}
