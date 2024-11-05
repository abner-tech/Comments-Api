package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func (a *applicationDependences) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//defer will be called when the stack unwinds
		defer func() {
			//recover from panic
			err := recover()
			if err != nil {
				w.Header().Set("Connection", "Close")
				a.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (a *applicationDependences) rateLimiting(next http.Handler) http.Handler {
	// rate limit struct
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time //remove map entries that are stable
	}

	var mu sync.Mutex
	var clients = make(map[string]*client)
	//a gorutine to remove stale entries from the map
	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock() //begin clean-up
			//delete any entry not seen in 3 minutes
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock() //finish cleanup
		}
	}()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//get the ip address
		if a.config.limiter.enabled {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				a.serverErrorResponse(w, r, err)
				return
			}
			mu.Lock() //exclusive access to the map
			//check if ip address already is in map, if not add it
			_, found := clients[ip]
			if !found {
				clients[ip] = &client{limiter: rate.NewLimiter(
					rate.Limit(a.config.limiter.rps),
					a.config.limiter.burst,
				)}
			}

			//update the last seen of the clients
			clients[ip].lastSeen = time.Now()

			//check the rate limit status
			if !clients[ip].limiter.Allow() {
				mu.Unlock() //no longer needs exclusive access to the map
				a.rateLimitExceededResponse(w, r)
				return
			}

			mu.Unlock() //others are free to get exclusive access to the map
		}
		next.ServeHTTP(w, r)

	})
}
