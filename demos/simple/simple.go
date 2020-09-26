package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"go.uber.org/multierr"

	"go.ptx.dk/cobradig"
)

type thingy string

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c := dig.New()
	if err := multierr.Combine(
		c.Provide(func() thingy {
			return "something"
		}),
	); err != nil {
		return err
	}

	root := &cobra.Command{
		Use:  "test",
		RunE: cobradig.Invoke(c, cmd.run),
	}

	return root.Execute()
}

type cmd struct {
	dig.In

	T thingy
}

func (c cmd) run() error {
	fmt.Println(c.T)
	return nil
}
