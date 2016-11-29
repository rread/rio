package main

import (
	"os"
	"os/exec"
)

// Shell runs command
func Shell(cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}
