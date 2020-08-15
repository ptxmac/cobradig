package waiter

import (
	"context"
	"time"
)

type Waiter interface {
	Wait(ctx context.Context, duration time.Duration) error
}

type waiter struct {
}

func (w *waiter) Wait(ctx context.Context, duration time.Duration) error {
	select {
	case <-ctx.Done():
		if err := ctx.Err(); err != nil {
			return err
		}
	case <-time.After(duration):
	}
	return nil
}

func New() Waiter {
	return &waiter{}
}
