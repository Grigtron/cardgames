package main

import (
	"fmt"
	"os"
)

func commandExit(args ...string) error {
	fmt.Println("Thanks for playing!")
	os.Exit(0)
	return nil
	
}