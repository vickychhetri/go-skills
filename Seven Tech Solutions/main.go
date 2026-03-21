package main

import (
	"fmt"
	"time"
)

func main() {
	var timeout <-chan time.Time // nil = no timeout

	enableTimeout := true

	if enableTimeout {
		timeout = time.After(2 * time.Second)
	}

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("work finished")
	case <-timeout:
		fmt.Println("timed out")
	}
}

// If enableTimeout = false → timeout case is disabled

// If true → timeout works normally

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("NIL Channel use")
// 	var ch chan int

// 	for i := 0; i < 10; i++ {
// 		select {
// 		case ch <- i:
// 			fmt.Println("sent- channel working: ", i)
// 			os.Exit(0)
// 		default:
// 			fmt.Println("default ")
// 		}

// 		if i == 5 {
// 			ch = make(chan int, 1)
// 		}
// 	}

// }
