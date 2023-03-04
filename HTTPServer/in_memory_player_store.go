package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.Mutex{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mutex sync.Mutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
