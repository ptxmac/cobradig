package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"gitlab.com/ptxmac/cobradig"
	"gitlab.com/ptxmac/cobradig/demos/complex/cmd/echo"
	"gitlab.com/ptxmac/cobradig/demos/complex/cmd/wait"
	"gitlab.com/ptxmac/cobradig/demos/complex/flags"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

func RootCmd(cc *dig.Container) (*cobra.Command, error) {
	root := &cobra.Command{
		Use: "complex",
	}

	err := multierr.Combine(
		flags.Add(root, cc),
		echo.AddCommands(root, cc),
		wait.AddCommands(root, cc),
		cc.Provide(timeoutBuilder),
		cobradig.ProvideContext(cc),
	)

	return root, err
}

func timeoutBuilder(f flags.Flags) cobradig.ContextBuilder {
	return func() (context.Context, context.CancelFunc) {
		return context.WithTimeout(context.Background(), f.Timeout)
	}
}
