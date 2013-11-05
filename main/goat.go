package main

import (
	"errors"
	"fmt"
	"os"
	"github.com/mediocregopher/goat"
)

func fatal(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func printGhelp() {
	fmt.Printf(
		`Goat is a command-line wrapper for go which handles dependency
management in a sane way. Check the goat docs at github.com/mediocregopher/goat
for a more in-depth overview.

Usage:

    %s command [arguments]

The commands are:

    deps    Read the Goatfile for this project and set up dependencies in the
            lib folder. Recursively download dependencies wherever a Goatfile is
            encountered

    ghelp   Show this dialog

All other commands are passed through to the go binary on your system. Try '%s
help' for its available commands
`, os.Args[0], os.Args[0])
	os.Exit(0)
}

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		fatal(err)
	}

	projroot, err := goat.FindGoatfile(cwd)
	var genv *goat.GoatEnv
	if err == nil {
		genv, err = goat.SetupGoatEnv(projroot)
		if err != nil {
			fatal(err)
		} else if err = goat.EnvPrependProj(genv); err != nil {
			fatal(err)
		}

		if err = goat.Setup(genv); err != nil {
			fatal(err)
		}
	}

	args := os.Args[1:]
	if len(args) < 1 {
		printGhelp()
	}

	switch args[0] {
	case "deps":
		if genv != nil {
			err := goat.FetchDependencies(genv)
			if err != nil {
				fatal(err)
			}
		} else {
			fatal(errors.New("Goatfile not found on current path"))
		}
	case "ghelp":
		printGhelp()
	default:
		if actualgo, ok := goat.ActualGo(); ok {
			goat.PipedCmd(actualgo, args...)
		} else {
			newargs := make([]string, len(args)+1)
			copy(newargs[1:], args)
			newargs[0] = "go"
			goat.PipedCmd("/usr/bin/env", newargs...)
		}
	}
}
