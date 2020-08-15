package main

import (
	"gitlab.com/ptxmac/cobradig/demos/complex/cmd"
	"gitlab.com/ptxmac/cobradig/demos/complex/lib/waiter"
	"go.uber.org/dig"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c := dig.New()
	if err := c.Provide(waiter.New); err != nil {
		return err
	}

	root, err := cmd.RootCmd(c)
	if err != nil {
		return err
	}
	return root.Execute()
}
