package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c := context.Background()
	ctx, cancel := context.WithTimeout(c, time.Second*2) //Change second param value to see the difference.
	defer cancel()
	fbrx := make(chan string)
	dbrx := make(chan string)
	go GetDataFrom("facebook", ctx, fbrx)
	go GetDataFrom("database", ctx, dbrx)
	defer close(fbrx)
	defer close(dbrx)
	for i := 0; i < 2; i++ {
		select {
		case fb := <-fbrx:
			fmt.Println(">>>>>> Data Received From: ", fb)
		case db := <-dbrx:
			fmt.Println(">>>>>> Data Received From: ", db)
		case <-ctx.Done():
			fmt.Println(">>>>> Timeout to process")
		}
	}
	time.Sleep(time.Second * 50) // Simulate the context
}

func GetDataFrom(src string, ctx context.Context, channel chan string) {
	startTime := time.Now()
	ticker := time.NewTicker(time.Second * 1)
	src_to := float64(5)
	if src == "database" {
		src_to = float64(10)
	}
	for _ = range ticker.C {
		fmt.Println("Still Processing From", src)
		if time.Since(startTime).Seconds() >= src_to { // Example: Fetch Data needs xx seconds
			ticker.Stop()
			break
		}
	}
	if ctx.Err() == nil { // ctx.Err() will filled when the ctx is timeout or the ctx canceled
		channel <- src
	}
	return
}
