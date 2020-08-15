package flags

import (
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"time"
)

type Flags struct {
	Verbose bool
	Timeout time.Duration
}

func Add(root *cobra.Command, cc *dig.Container) error {
	f := Flags{
		Timeout: 10 * time.Second,
	}
	root.PersistentFlags().BoolVarP(&f.Verbose, "verbose", "v", f.Verbose, "")
	root.PersistentFlags().DurationVarP(&f.Timeout, "timeout", "t", f.Timeout, "")
	return cc.Provide(func() Flags {
		return f
	})
}
