package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

type S struct {
	I int
}
type M map[S][]string

func main() {
	var list []int = []int{
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24,
	}

	results := make(chan M, len(list))

	for items := range slices.Chunk(list, 5) {
		var wg sync.WaitGroup
		for _, item := range items {
			wg.Add(1)
			go func() {
				defer wg.Done()
				fmt.Printf("job for item %d started\n", item)
				time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
				results <- M{{I: item}: []string{fmt.Sprintf("item: %d", item)}}
				fmt.Printf("job for item %d finished\n", item)
			}()
		}
		wg.Wait()
	}

	var res []M = make([]M, len(list))
	for i := range len(list) {
		r := <-results
		res[i] = r
		fmt.Printf("iteration: %d, item: %+v\n", i, r)
	}

	fmt.Printf("\n\nres\n%+v\n", res)
}
