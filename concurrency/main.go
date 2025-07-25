package main

import (
	"context"
	"fmt"
	"iter"
	"maps"
	"sync"
	"time"
)

type svc struct {
	items map[string]string
}

func (this *svc) Items(m *sync.Mutex) iter.Seq2[string, string] {
	// fmt.Println("Items locked mutex")
	// m.Lock()
	// defer func() {
	// 	m.Unlock()
	// 	fmt.Println("Items unlocked mutex")
	// }()
	return maps.All(this.items)
}

type collector struct {
	svc *svc
}

func main() {
	// bootstrapping

	var m sync.Mutex = sync.Mutex{}

	var s *svc = &svc{items: map[string]string{}}
	var c *collector = &collector{svc: s}

	var ctx context.Context
	var x context.CancelFunc
	ctx, x = context.WithCancel(context.Background())

	go func() {
		fmt.Println("Started polling...")
		// initial update
		var i int
		for i = range 10 {
			c.svc.items[fmt.Sprintf("item:%d", i)] = "some-new-item"
		}

		var ticker *time.Ticker = time.NewTicker(1 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context Finished... Ending polling")
				return
			case <-ticker.C:
				m.Lock()
				c.svc.items = map[string]string{}
				for i = range 10 {
					c.svc.items[fmt.Sprintf("item:%d", i)] = "some-new-item"
				}
				m.Unlock()
			}
		}
	}()

	// the equivalent of the server serving/listening for http
	go func() {
		fmt.Println("Started listening...")
		var ticker *time.Ticker = time.NewTicker(1 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context Finished... Ending listening")
				return
			case <-ticker.C:
				var k string
				var v string
				var items iter.Seq2[string, string] = s.Items(&m)
				fmt.Println("Reading from Items")
				for k, v = range items {
					fmt.Printf("Poller added key: %s with value: %s\n", k, v)
				}
				// fmt.Printf("key item:0, value: %s\n", s.items["item:0"])
				fmt.Println("--------------------------------------------------")
			}
		}
	}()

	fmt.Println("Program running... 5min sleep...")
	time.Sleep(5 * time.Minute)
	fmt.Println("Program ending... 5min elapsed...")
	x()

	return
}
