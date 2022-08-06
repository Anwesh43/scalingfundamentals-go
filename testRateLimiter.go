package main 

import (
	"net/http"
	"time"
	"fmt"
	"io"
)

func main() {
	for i := 0; i < 30; i++ {
		res, err := http.Get("http://localhost:8090/hello")
		if err == nil  {
			b, er := io.ReadAll(res.Body)
			if er == nil {
				fmt.Println(string(b))
			} else {
				fmt.Println("error reading body", er)
			}
			res.Body.Close()
		} else {
			fmt.Println("Error",err)
		}
		time.Sleep(time.Second)
	}
}