package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...) // #nosec G204
	command.Env = make([]string, 0, 50)
	envmap := make(map[string]string, 50)

	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	for _, v := range os.Environ() {
		e := strings.SplitN(v, "=", 2)
		if len(e) == 2 {
			envmap[e[0]] = e[1]
		}
	}
	for k, v := range env {
		if v.NeedRemove {
			delete(envmap, k)
		} else {
			envmap[k] = v.Value
		}
	}
	for k, v := range envmap {
		command.Env = append(command.Env, k+"="+v)
	}

	if err := command.Start(); err != nil {
		log.Printf("cmd.Start: %v", err)
		return -1
	}
	if err := command.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok { //nolint:errorlint
			returnCode = exiterr.ExitCode()
		} else {
			log.Printf("cmd.Wait: %v", err)
			return -1
		}
	}

	return returnCode
}
