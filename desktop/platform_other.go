//go:build !windows

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func hideConsole(cmd *exec.Cmd) {}

func fatal(msg string) { fmt.Fprintln(os.Stderr, msg) }
