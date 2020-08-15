package cobradig

import (
	"context"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

// ContextBuilder must be provided in the container
type ContextBuilder func() (context.Context, context.CancelFunc)

// ProvideContext this will look for a ContextBuilder and invoke it if a context is needed.
// the cancel function is added as a post hook
func ProvideContext(cc *dig.Container) error {
	var cancelFunc context.CancelFunc
	return multierr.Combine(
		cc.Provide(func() Hook {
			return Hook{
				Post: func() error {
					if cancelFunc != nil {
						cancelFunc()
					}
					return nil
				},
			}
		}, dig.Group("hooks")),
		cc.Provide(func(cb ContextBuilder) (context.Context, error) {
			ctx, cancel := cb()
			cancelFunc = cancel
			return ctx, nil
		}),
	)
}
