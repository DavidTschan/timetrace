package cli

import (
	"github.com/dominikbraun/timetrace/core"
	"github.com/dominikbraun/timetrace/out"

	"github.com/spf13/cobra"
)

func toggleCommand(t *core.Timetrace) *cobra.Command {
	toggle := &cobra.Command{
		Use:   "toggle start / stop last project",
		Short: "toggle start / stop last project",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if err := t.Toggle(); err != nil {
				out.Err("failed to toggle tracking: %s", err.Error())
				return
			}
			out.Success("Toggled")
		},
	}

	return toggle
}
