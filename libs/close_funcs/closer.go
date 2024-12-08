package close_funcs

import (
	"context"
	"fmt"
	"sync"
)

type CloseFunc func()

type Closer struct {
	mu    sync.Mutex
	funcs []CloseFunc
}

func New() *Closer {
	return &Closer{}
}

func (c *Closer) Add(fn CloseFunc) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs = append(c.funcs, fn)
}

func (c *Closer) Close(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	done := make(chan struct{}, 1)
	go func() {
		for _, fn := range c.funcs {
			fn()
		}

		done <- struct{}{}
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("shutdown cancelled: %v", ctx.Err())
	}
}
