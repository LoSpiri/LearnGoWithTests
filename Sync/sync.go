package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

type Counter2 struct {
	sync.Mutex
	value int
}

func (c *Counter2) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}
