package echo

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/ptxmac/cobradig"
	"go.uber.org/dig"
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
