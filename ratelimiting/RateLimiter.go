package ratelimiting

import (
	"time"
	"net/http"
	"fmt"
)
type Request struct {
	curr time.Time
}

type RateLimiterHandler interface {
	func(http.HandlerFunc) (http.HandlerFunc);
}
func timeDifference(t1, t2 time.Time) int {
	return int(t2.UnixMilli() - t1.UnixMilli())
}

func RateLimiter(durationInMillis int, total int) func(http.HandlerFunc) (http.HandlerFunc) {
	requests := make([]Request, 0)
	return func(cb http.HandlerFunc) (http.HandlerFunc) {
		return func(w http.ResponseWriter, r *http.Request) {
			if len(requests) > total {
				firstRequest := requests[0]
				t := time.Now()
				fmt.Println("Difference in time", timeDifference(firstRequest.curr, t))
				if timeDifference(firstRequest.curr, t) > durationInMillis {
					requests = requests[1:]
					requests = append(requests, Request {
						curr: t,
					})
					cb(w, r)
				} else {
					fmt.Fprintf(w, "request not allowed")
				}
			} else {
				requests = append(requests, Request {
					curr : time.Now(),
				})
				cb(w, r)
			}
			fmt.Println("Requests size", len(requests))
		}
	}
}