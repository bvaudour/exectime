// exectime project main.go
package main

import (
	"exectime/timer"
	"fmt"
	"os"
	"os/exec"
)

func printError(e interface{}) {
	fmt.Printf("\033[1;31m%v\033[m\n", e)
}

func launchCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	return cmd.Run()
}

func main() {
	if len(os.Args) < 2 {
		printError("Add at least one argument!")
		return
	}
	f := func() {
		if e := launchCommand(os.Args[1], os.Args[2:]...); e != nil {
			printError(e)
		}
	}
	t := timer.NewFunc(f)
	t.Exec()
	fmt.Println(t)
}
