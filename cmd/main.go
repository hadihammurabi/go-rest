package main

import (
	"context"
	"go-rest/api"
	"go-rest/driver"
	"go-rest/entity"
	"go-rest/repository"
	"go-rest/service"
	"os"

	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
)

func init() {
	driver.Init()
	repository.Init()
	service.Init()
}

func main() {
	command := "run"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "run":
		cmdRun()
	case "seed":
		cmdSeed()
	}

}

func cmdRun() {
	go api.Run()
	go gowok.StartPProf()
	gowok.GracefulStop(func() {
		println()
		println("Stopping...")
	})
}

func cmdSeed() {
	_, err := ioc.MustGet(service.UserService{}).Create(context.Background(), &entity.User{
		Email:    "root",
		Password: "123123",
	})
	if err != nil {
		panic(err)
	}

	println("done")
}
