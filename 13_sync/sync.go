package sync

import "sync"

type Counter struct {
	mu    sync.Mutex // a Mutex is a mutual exclusion lock
	value int
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
