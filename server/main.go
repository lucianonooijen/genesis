package main

import (
	"fmt"
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/config"
)

func main() {
	c, err := config.LoadConfig()
	fmt.Println(c)
	fmt.Println(err)
	panic("Not implemented")
}
