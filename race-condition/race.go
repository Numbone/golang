package race_condition

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	count int
	mu    *sync.Mutex //race condition start
}

func (c *counter) increment() {
	c.mu.Lock() // blocked when two goroutine refer one point bite
	c.count++
	c.mu.Unlock()
}

func (c *counter) value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
func mutexExp() {
	c := counter{
		mu: new(sync.Mutex),
	}
	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(c.count)
}
