package main

import (
	"fmt"
	"context"
	"time"
)

func SomeContextfullFunction(ctx context.Context) error {
	i := 0
	for i < 1000000 {
		i = i + 1
		select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				fmt.Println("still running...")
		}
	}

	return nil
}

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(100 * time.Millisecond))
	defer cancel()

	err := SomeContextfullFunction(ctx)
	if err != nil {
		fmt.Println("Got error: ", err)
	}
}
