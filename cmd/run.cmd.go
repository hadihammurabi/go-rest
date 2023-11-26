package cmd

import (
	"go-rest/api"

	"github.com/gowok/gowok"
	"github.com/spf13/cobra"
)

func cmdRun() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "run app",
		Run: func(cmd *cobra.Command, args []string) {
			go api.Run()
			gowok.GracefulStop(func() {
				println()
				println("Stopping...")
			})
		},
	}
}
