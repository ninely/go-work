package main

import (
	"sync"
	"time"
)

const windowSize = 5

type window struct {
	buckets map[int64]*numberBucket
	mtx     sync.Mutex
}

type numberBucket struct {
	value int
}

func NewWindow() *window {
	return &window{
		buckets: make(map[int64]*numberBucket),
	}
}

func (w *window) Increment(i int) {
	if i == 0 {
		return
	}
	w.mtx.Lock()
	defer w.mtx.Unlock()

	now := time.Now().Unix()
	if bucket, ok := w.buckets[now]; ok {
		bucket.value += i
	} else {
		w.buckets[now] = &numberBucket{value: i}
	}

	for timestamp := range w.buckets {
		if timestamp <= now-windowSize {
			delete(w.buckets, timestamp)
		}
	}
}

func (w *window) Sum(now time.Time) int {
	sum := 0
	w.mtx.Lock()
	defer w.mtx.Unlock()

	for timestamp, bucket := range w.buckets {
		if timestamp > now.Unix()-windowSize {
			sum += bucket.value
		}
	}
	return sum
}
