package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config, args ...string) error {
	fmt.Println("turning off...")
	os.Exit(0)
	return nil
}
