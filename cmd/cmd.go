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
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(cmdRun())
	cmd.AddCommand(cmdSeed())

	return cmd
}
