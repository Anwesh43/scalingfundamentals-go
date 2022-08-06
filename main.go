package main 

import (
	"fmt"
	"net/http"
	"scalingdemo/ratelimiting"
)



func main() {
	rl := ratelimiting.RateLimiter(10000, 7)
	http.HandleFunc("/hello", rl(func (w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello")
	}))
	if  err := http.ListenAndServe(":8090", nil); err == nil {
		fmt.Println("started server")
	} else {
		fmt.Println(err)
	}
	
	
}
