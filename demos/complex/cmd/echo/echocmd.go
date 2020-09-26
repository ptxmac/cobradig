package echo

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/dig"

	"go.ptx.dk/cobradig"
)

func AddCommands(parent *cobra.Command, cc *dig.Container) error {
	c := &cobra.Command{
		Use:  "echo",
		RunE: cobradig.Invoke(cc, echo),
	}

	parent.AddCommand(c)
	return nil
}

func echo(args []string) error {
	fmt.Println(args)
	return nil
}
