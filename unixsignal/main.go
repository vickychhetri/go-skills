package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
“Signal handling allows Go applications to gracefully respond to OS-level
 events like termination, interrupts, and reload requests,
 which is critical for production services, containers, and long-running processes.”
*/

/*
Signal Handling in Go
---------------------

This program demonstrates how to handle Unix OS signals in Go.

Why signals matter?
- OS sends signals to control running processes
- Allows graceful shutdown instead of force killing the app
- Essential for Docker, Kubernetes, systemd, CLI tools

Common Signals:
- SIGINT  → Ctrl+C (developer interrupts)
- SIGTERM → Graceful termination (Docker/K8s/systemd)
- SIGHUP  → Reload config / terminal closed
- SIGQUIT → Quit + stack dump
*/

// handleSignal performs actions based on received OS signals
func handleSignal(sig os.Signal, start time.Time, cancel context.CancelFunc) {
	switch sig {

	// Ctrl + C (Local development / CLI tools)
	case syscall.SIGINT:
		duration := time.Since(start)
		fmt.Println("SIGINT received (Ctrl+C)")
		fmt.Println("Execution time:", duration)

		// Graceful cleanup logic
		cancel()

	// Graceful shutdown (Docker, Kubernetes, systemd)
	case syscall.SIGTERM:
		fmt.Println("SIGTERM received (termination request)")

		// Cleanup before exit
		cancel()
		os.Exit(0)

	// Terminal closed or reload configuration
	case syscall.SIGHUP:
		fmt.Println("SIGHUP received (reload configuration)")
		// Example:
		// reloadConfig()
		// reopenLogFiles()

	// Ctrl + \ (force quit with dump)
	case syscall.SIGQUIT:
		fmt.Println("SIGQUIT received (quit + stack dump)")
		// Dump goroutines / debug info
		os.Exit(1)

	default:
		fmt.Println("Unhandled signal:", sig)
	}
}

func main() {
	fmt.Println("🔔 Handling Unix Signals in Go")

	// Print process ID (useful for kill command)
	fmt.Printf("Process ID: %d\n", os.Getpid())

	// Context used for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Channel to receive OS signals
	sigChan := make(chan os.Signal, 1)

	// Notify for specific signals (best practice)
	signal.Notify(
		sigChan,
		syscall.SIGINT,
		syscall.SIGTERM, //docker stop my_container
		syscall.SIGHUP,
		syscall.SIGQUIT,
	)

	start := time.Now()

	// Goroutine listening for OS signals
	go func() {
		for {
			select {
			case sig := <-sigChan:
				handleSignal(sig, start, cancel)

			case <-ctx.Done():
				fmt.Println("Graceful shutdown completed")
				return
			}
		}
	}()

	/*
		Main application logic
		-----------------------
		Simulates a long-running service (API, worker, consumer)
	*/
	fmt.Println("Application is running... Press Ctrl+C to stop")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Main loop stopped")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(3 * time.Second)
		}
	}
}
