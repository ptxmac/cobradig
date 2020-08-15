package cobradig

import (
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

type Hook struct {
	// Pre func () error
	Post func() error
}

type hookContainer struct {
	dig.In
	Hooks []Hook `group:"hooks"`
}

func Invoke(c *dig.Container, f interface{}) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var hooks []Hook
		if err := multierr.Combine(
			c.Provide(func() []string { return args }),
			c.Invoke(func(h hookContainer) {
				hooks = h.Hooks
			}),
		); err != nil {
			return err
		}

		err := c.Invoke(f)

		for _, h := range hooks {
			err = multierr.Append(err, h.Post())
		}
		return err
	}
}
