package wait

import (
	"context"
	"github.com/spf13/cobra"
	"gitlab.com/ptxmac/cobradig"
	"gitlab.com/ptxmac/cobradig/demos/complex/lib/waiter"
	"go.uber.org/dig"
	"time"
)

func AddCommands(parent *cobra.Command, cc *dig.Container) error {
	c := &cobra.Command{
		Use:  "wait",
		RunE: cobradig.Invoke(cc, cmd.wait),
	}
	parent.AddCommand(c)
	return nil
}

type cmd struct {
	dig.In

	Waiter waiter.Waiter
}

func (c cmd) wait(ctx context.Context) error {
	return c.Waiter.Wait(ctx, 5*time.Second)
}
