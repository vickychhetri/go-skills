package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"sync/atomic"
	"time"
)

/*
pprof works because importing net/http/pprof runs its init() function, which automatically registers profiling HTTP
handlers on Go’s default HTTP mux. When http.ListenAndServe is started with a nil handler, it uses this default mux,
making the pprof endpoints available without extra configuration.

One-liner to remember 🧠
Blank import → init() → handlers registered → DefaultServeMux → ListenAndServe(nil)

We analyze pprof profiles using go tool pprof. The top command shows hot functions, list maps CPU usage to source code,
and the web UI provides flame graphs and call graphs for deeper analysis.
*/
func main() {
	var counter int64
	var wg sync.WaitGroup

	// Start pprof server
	go func() {
		fmt.Println("pprof running at http://localhost:6060/debug/pprof/")
		http.ListenAndServe("localhost:6060", nil)
	}()

	start := time.Now()

	total := 1_000_000 // 1 million (safe)
	for i := 0; i < total; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()

	fmt.Println("counter:", counter)
	fmt.Println("time taken:", time.Since(start))

	// Keep program alive for pprof analysis
	select {}
}
