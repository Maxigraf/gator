package main

import (
	"fmt"

	"github.com/maxigraf/gator/internal/config"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return
	}

	err = cfg.SetUser("maxigraf")

	if err != nil {
		fmt.Printf("Cannot set user: %v\n", err)
	}

	cfg, err = config.Read()

	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
	}

	fmt.Printf("Data: %v\n", cfg)
}
