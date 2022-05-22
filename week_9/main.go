package main

import (
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func main() {
	eg := &errgroup.Group{}
	eg.Go(func() error {
		time.Sleep(1 * time.Second)
		RunClient()
		return nil
	})

	eg.Go(func() error {
		Listen()
		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Printf("failed: %v", err)
	}
	log.Println("end")
}
