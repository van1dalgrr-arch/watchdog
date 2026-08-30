package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for {
		resp, _ := http.Get("http://localhost:8080/health")

		if resp.StatusCode == 200 {
			fmt.Println("[OK] your server is healthy")

		} else {
			fmt.Println("[ERROR] your server unhealthy ")
		}

		resp.Body.Close()

		time.Sleep(5 * time.Second)
	}
}
