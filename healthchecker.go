package main

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
)

func main() {
	wait := 30
	for true {
		fmt.Println("Sending healthcheck")
		resp, err := http.Get("http://<SERVER_IP:PORT>")
		if err != nil {
			fmt.Println("Failed to send healthcheck")
			fmt.Println("Retrying after " + strconv.Itoa(wait) + " seconds")
			time.Sleep(time.Duration(wait) * time.Second)
			continue
		} else {
			fmt.Println(resp)
			fmt.Println("Retrying after " + strconv.Itoa(wait) + " seconds")
			time.Sleep(time.Duration(wait) * time.Second)
		}
	}
}
