package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please use: go-envdir path-to-envdir command args... ")
		return
	}
	env, err := ReadDir(os.Args[1])
	if errors.Is(err, ErrOpenDirectory) {
		fmt.Println("Bad directory with env variables")
		os.Exit(1)
	} else if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(RunCmd(os.Args[2:], env))
}
