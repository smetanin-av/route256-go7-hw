package worker_pool

import (
	"context"
	"fmt"
	"sync"
)

// Result результат работы пула
type Result[TIn any, TOut any] struct {
	Inp TIn
	Out TOut
	Err error
}

type PoolFn[TIn any, TOut any] func(ctx context.Context, inp TIn) (TOut, error)

// Process - обработка списка элементов в пуле с помощью указанной функции,
// возвращается канал с результатами и ошибка
func Process[TIn any, TOut any](
	ctx context.Context,
	poolSize int,
	items []TIn,
	poolFn PoolFn[TIn, TOut],
) <-chan Result[TIn, TOut] {
	inputs := sliceToChannel[TIn](items)
	close(inputs)

	results := make(chan Result[TIn, TOut], len(items))
	wg := new(sync.WaitGroup)

	pool := simplePool[TIn, TOut]{inputs, results, poolFn}
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pool.worker(ctx)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func sliceToChannel[TItem any](items []TItem) chan TItem {
	channel := make(chan TItem, len(items))
	for _, item := range items {
		channel <- item
	}
	return channel
}

type simplePool[TIn any, TOut any] struct {
	inputs  chan TIn
	results chan Result[TIn, TOut]
	poolFn  PoolFn[TIn, TOut]
}

func (p *simplePool[TIn, TOut]) worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case inp, ok := <-p.inputs:
			if !ok {
				return
			}

			res := Result[TIn, TOut]{Inp: inp}
			func() {
				defer func() {
					if msg := recover(); msg != nil {
						res.Err = fmt.Errorf("panic: %v", msg)
					}
				}()
				res.Out, res.Err = p.poolFn(ctx, inp)
			}()
			p.results <- res
		}
	}
}
