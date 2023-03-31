package cmd

import (
	"context"
	"go-rest/entity"
	"go-rest/service"

	"github.com/gowok/ioc"
	"github.com/spf13/cobra"
)

func cmdSeed() *cobra.Command {
	return &cobra.Command{
		Use:   "seed",
		Short: "prepare initial data",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := ioc.MustGet(service.UserService{}).Create(context.Background(), &entity.User{
				Email:    "root",
				Password: "123123",
			})
			if err != nil {
				panic(err)
			}

			println("done")
		},
	}
}
