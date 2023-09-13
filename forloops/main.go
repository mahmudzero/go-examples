package main

import "fmt"

func main() {
	var counter int = -1;
	for counter < 100 {
		counter = counter + 1;
		if counter % 2 == 0 {
			continue
		}

		fmt.Printf("counter %d\n", counter);
	}
}
