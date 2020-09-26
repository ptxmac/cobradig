package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"go.uber.org/multierr"

	"go.ptx.dk/cobradig"
	"go.ptx.dk/cobradig/demos/complex/cmd/echo"
	"go.ptx.dk/cobradig/demos/complex/cmd/wait"
	"go.ptx.dk/cobradig/demos/complex/flags"
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
