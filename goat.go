package main

import (
	"errors"
	"fmt"
	"github.com/mediocregopher/goat/env"
	"github.com/mediocregopher/goat/exec"
	"os"
	"syscall"
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

  deps [install]
      Read the .go.yaml file for this project and install specified dependencies
      in the .go folder. Recursively download dependencies wherever a .go.yaml
      file is encountered.

  deps uninstall
      Remove all previously installed dependencies.

  deps help
      Show this dialog.

All other commands are passed through to the go binary on your system.
Try '%s help' for its available commands
`, os.Args[0], os.Args[0])
	os.Exit(0)
}

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		fatal(err)
	}

	projroot, err := env.FindProjRoot(cwd)
	var genv *env.GoatEnv
	if err == nil {
		genv, err = env.NewGoatEnv(projroot)
		if err != nil {
			fatal(err)
		}

		if err = genv.PrependToGoPath(); err != nil {
			fatal(err)
		}

		if err = genv.Setup(); err != nil {
			fatal(err)
		}
	}

	args := os.Args[1:]
	if len(args) < 1 { printGhelp()	}

  args = append(args, "") // ensure there is an args[1]

	if args[0] == "deps" {
		switch args[1] {
		case "install":
			InstallDepsCmd(genv)
		case "uninstall", "remove":
			UninstallDepsCmd(genv)
		case "help":
			printGhelp()
		default:
			InstallDepsCmd(genv)
		}
	} else {
		if actualgo, ok := ActualGo(); ok {
			exec.PipedCmd(actualgo, args...)
		} else {
			newargs := make([]string, len(args)+1)
			copy(newargs[1:], args)
			newargs[0] = "go"
			exec.PipedCmd("/usr/bin/env", newargs...)
		}
	}
}

func InstallDepsCmd(genv *env.GoatEnv) {
	if genv != nil {
		err := genv.FetchDependencies(genv.AbsDepDir())
		if err != nil {
			fatal(err)
		}
	} else {
		fatal(errors.New(".go.yaml file not found on current path"))
	}
}

// TODO: Maybe ask or pause a moment before actual removal?
func UninstallDepsCmd(genv *env.GoatEnv) {
	if genv != nil {
		err := genv.RemoveDependencies(genv.AbsDepDir())
		if err != nil {
			fatal(err)
		}
	} else {
		fatal(errors.New(".go.yaml file not found on current path"))
	}
}

// ActualGo returns the GOAT_ACTUALGO environment variable contents, and whether
// or not the variable was actually set
func ActualGo() (string, bool) {
	return syscall.Getenv("GOAT_ACTUALGO")
}
