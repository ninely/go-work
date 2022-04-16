package main

import (
	"fmt"
	"time"
)

func main() {
	w := NewWindow()

	request := []int{2, 5, 1, 3, 4, 2, 1}

	for _, x := range request {
		w.Increment(x)

		sum := w.Sum(time.Now())

		fmt.Printf("request: %d sum: %d\n", x, sum)

		time.Sleep(1 * time.Second)
	}
}
