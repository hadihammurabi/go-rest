package main

import "go-rest/cmd"

func main() {
	cmd := cmd.New()

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
