package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	apiv2 "github.com/daopmdean/rate-limiter-demo/api-v2"
)

func main() {
	testApi()
}

func testApi() {
	defer fmt.Println("Exit Server")

	server := apiv2.Open()

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := server.Read(context.Background())
			if err != nil {
				fmt.Printf("%v Got Error from read: %v\n", time.Now().Format("15:04:05"), err)
				return
			}

			fmt.Printf("%v Get Data: %v\n", time.Now().Format("15:04:05"), data)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := server.Resolve(context.Background())
			if err != nil {
				fmt.Printf("%v Got Error from resolve: %v\n", time.Now().Format("15:04:05"), err)
				return
			}

			fmt.Printf("%v Resolved\n", time.Now().Format("15:04:05"))
		}()
	}

	wg.Wait()
}
