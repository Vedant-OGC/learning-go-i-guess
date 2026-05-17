package concurrency

import (
	"fmt"
	"sync"
)

type RWSafeMap struct {
	m   map[string]string
	mux sync.RWMutex
}

func (r *RWSafeMap) Get(key string) string {
	r.mux.RLock() // read lock (multiple readers allowed)
	defer r.mux.RUnlock()
	return r.m[key]
}

func (r *RWSafeMap) Set(key, value string) {
	r.mux.Lock() // write lock (exclusive)
	defer r.mux.Unlock()
	r.m[key] = value
}

func MutexStuff() {
	r := RWSafeMap{m: make(map[string]string)}
	r.Set("status", "operational")
	fmt.Println("status:", r.Get("status"))
}
