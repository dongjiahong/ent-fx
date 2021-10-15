package main

import (
	"web/cmd"
)

func main() {
	if err := cmd.Cmd.Execute(); err != nil {
		panic(err)
	}
}
