package worker_pool

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	delay     = time.Millisecond * 2
	errorMsg  = "test error"
	panicMsg  = "panic: test error"
	itemsSize = 1000
)

var (
	errFailed = fmt.Errorf(errorMsg)
)

func TestProcess_Passed(t *testing.T) {
	t.Parallel()

	type testCase[TInput any, TOutput any] struct {
		name     string
		poolSize int
		wantTime time.Duration
	}

	items := make([]int, 0, itemsSize)
	for i := 0; i < itemsSize; i++ {
		items = append(items, i)
	}

	poolFn := func(_ context.Context, inp int) (int, error) {
		time.Sleep(delay)
		return inp * 2, nil
	}

	tests := []testCase[int, int]{
		{
			name:     "pool size as used in listCart",
			poolSize: 5,
			wantTime: time.Millisecond * 500,
		},
		{
			name:     "just one worker",
			poolSize: 1,
			wantTime: time.Millisecond * 2500,
		},
		{
			name:     "pool size equal items size",
			poolSize: len(items),
			wantTime: time.Millisecond * 50,
		},
		{
			name:     "pool size larger then items size",
			poolSize: len(items) * 2,
			wantTime: time.Millisecond * 50,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			start := time.Now()
			results := Process(context.Background(), tt.poolSize, items, poolFn)

			counter := 0
			for it := range results {
				assert.NoError(t, it.Err)
				assert.Equal(t, it.Inp*2, it.Out)
				counter++
			}
			assert.Equal(t, len(items), counter)

			elapsed := time.Since(start)
			assert.Less(t, elapsed, tt.wantTime)
		})
	}
}

func TestProcess_Failed(t *testing.T) {
	t.Parallel()

	items := make([]int, 0, itemsSize)
	for i := 0; i < itemsSize; i++ {
		items = append(items, i)
	}

	t.Run("one error", func(t *testing.T) {
		t.Parallel()

		results := Process(context.Background(), 5, items,
			func(_ context.Context, inp int) (int, error) {
				time.Sleep(delay)
				if inp == 5 {
					return 0, errFailed
				}
				return inp * 2, nil
			})

		counter := 0
		for it := range results {
			if it.Inp == 5 {
				assert.ErrorIs(t, it.Err, errFailed)
			} else {
				assert.NoError(t, it.Err)
				assert.Equal(t, it.Inp*2, it.Out)
			}
			counter++
		}
		assert.Equal(t, len(items), counter)
	})

	t.Run("many errors", func(t *testing.T) {
		t.Parallel()

		results := Process(context.Background(), 5, items,
			func(_ context.Context, inp int) (int, error) {
				time.Sleep(delay)
				if inp%2 == 0 {
					return 0, errFailed
				}
				return inp * 2, nil
			})

		counter := 0
		for it := range results {
			if it.Inp%2 == 0 {
				assert.ErrorIs(t, it.Err, errFailed)
			} else {
				assert.NoError(t, it.Err)
				assert.Equal(t, it.Inp*2, it.Out)
			}
			counter++
		}
		assert.Equal(t, len(items), counter)
	})

	t.Run("all errors", func(t *testing.T) {
		t.Parallel()

		results := Process(context.Background(), 5, items,
			func(_ context.Context, inp int) (int, error) {
				time.Sleep(delay)
				return 0, errFailed
			})

		counter := 0
		for it := range results {
			assert.ErrorIs(t, it.Err, errFailed)
			counter++
		}
		assert.Equal(t, len(items), counter)
	})

	t.Run("one panic", func(t *testing.T) {
		t.Parallel()

		results := Process(context.Background(), 5, items,
			func(_ context.Context, inp int) (int, error) {
				time.Sleep(delay)
				if inp == 5 {
					panic(errorMsg)
				}
				return inp * 2, nil
			})

		counter := 0
		for it := range results {
			if it.Inp == 5 {
				assert.EqualError(t, it.Err, panicMsg)
			} else {
				assert.NoError(t, it.Err)
				assert.Equal(t, it.Inp*2, it.Out)
			}
			counter++
		}
		assert.Equal(t, len(items), counter)
	})

	t.Run("many panics", func(t *testing.T) {
		t.Parallel()

		results := Process(context.Background(), 5, items,
			func(_ context.Context, inp int) (int, error) {
				time.Sleep(delay)
				if inp%2 == 0 {
					panic(errorMsg)
				}
				return inp * 2, nil
			})

		counter := 0
		for it := range results {
			if it.Inp%2 == 0 {
				assert.EqualError(t, it.Err, panicMsg)
			} else {
				assert.NoError(t, it.Err)
				assert.Equal(t, it.Inp*2, it.Out)
			}
			counter++
		}
		assert.Equal(t, len(items), counter)
	})

	t.Run("all panics", func(t *testing.T) {
		t.Parallel()

		results := Process(context.Background(), 5, items,
			func(_ context.Context, inp int) (int, error) {
				time.Sleep(delay)
				panic(errorMsg)
			})

		counter := 0
		for it := range results {
			assert.EqualError(t, it.Err, panicMsg)
			counter++
		}
		assert.Equal(t, len(items), counter)
	})

	t.Run("stop by timeout", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), delay*10)
		defer cancel()

		results := Process(ctx, 5, items,
			func(_ context.Context, inp int) (int, error) {
				time.Sleep(delay)
				return inp * 2, nil
			})

		counter := 0
		for it := range results {
			assert.NoError(t, it.Err)
			assert.Equal(t, it.Inp*2, it.Out)
			counter++
		}
		assert.Less(t, counter, len(items))
	})

}
