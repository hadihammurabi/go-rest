package cmd

import (
	"go-rest/driver"
	"go-rest/repository"
	"go-rest/service"

	"github.com/spf13/cobra"
)

func init() {
	driver.Init()
	repository.Init()
	service.Init()
}

func New() *cobra.Command {
	cr := cmdRun()
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cr.Run(cmd, args)
		},
	}

	cmd.AddCommand(cr)
	cmd.AddCommand(cmdSeed())

	return cmd
}
