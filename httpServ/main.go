package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//MUX

/*
Lifecycle summary (mental model)
New Server

	↓

ListenAndServe()

	↓

Accept connections

	↓

Handle requests (goroutines)

	↓

Shutdown() or Close()

	↓

Server exits
*/
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/information", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hellow Server: Wipro"))
	})

	/**
	Common fields:
		Addr – address and port to listen on
		Handler – HTTP handler (if nil, http.DefaultServeMux is used)
		ReadTimeout, WriteTimeout, IdleTimeout – connection timeouts
		TLSConfig – TLS settings (optional)
	*/
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Server error %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit

	/*
		Graceful shutdown
			The 10 seconds is a deadline:
				“Give existing requests up to 10 seconds to finish.”
	*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	/*
		Immediate shutdown (force close)
		If you don’t need grace:
			srv.Close()
			Immediately closes all connections
			In-flight requests are aborted
	*/

}

/**
Request handling lifecycle

For each request:

TCP connection accepted
    ↓
HTTP request read
    ↓
Handler.ServeHTTP(w, r)
    ↓
Response written
    ↓
Connection reused or closed
*/

/*

FOR TLSConfig – TLS settings (optional)

“TLS behavior” means how a server and client negotiate, secure, and manage an encrypted connection.
In Go (and generally), it’s the set of rules that decide how HTTPS works under the hood.

Think of TLS as the security handshake + encrypted tunnel.
TLS behavior controls how that tunnel is created and enforced.

Big picture (simple)

When a client connects via HTTPS:

Client ──(TLS handshake)──> Server
          ├─ verify identity (cert)
          ├─ agree on TLS version
          ├─ agree on encryption
          └─ establish encrypted channel


TLS behavior = the rules used in this process
*/
