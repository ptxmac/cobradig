package main

import (
	"log"

	"go.uber.org/dig"

	"go.ptx.dk/cobradig/demos/complex/cmd"
	"go.ptx.dk/cobradig/demos/complex/lib/waiter"
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
