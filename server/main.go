package main

import (
	"fmt"

	"git.bytecode.nl/bytecode/genesis/internal/server"

	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/config"
)

func main() {
	c, err := config.LoadConfig()
	fmt.Println(c)
	fmt.Println(err)

	s, err := server.New(server.Requirements{
		Debug: c.IsDevMode,
		Port:  c.ServerPort,
	})
	if err != nil {
		panic(err)
	}

	err = s.Start()
	if err != nil {
		panic(err)
	}
}
