package main

import (
	"fmt"
	"go-rest/api"
	"go-rest/driver"
	"go-rest/repository"

	"github.com/gowok/gowok"
)

func init() {
	driver.PrepareAll()
	repository.NewRepository()
}

func main() {
	go api.Run()
	go gowok.StartPProf()
	gowok.GracefulStop(func() {
		fmt.Println()
		fmt.Println("Stopping...")
	})
}
