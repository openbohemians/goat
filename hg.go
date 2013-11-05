package goat

import (
	"fmt"
	"os"
	"path/filepath"
)

func Hg(genv *GoatEnv, dep *Dependency) error {
	localloc := filepath.Join(genv.ProjRootLib, "src", dep.Path)

	fmt.Println("hg", "clone", dep.Location, localloc)
	err := PipedCmd("hg", "clone", dep.Location, localloc)
	if err != nil {
		return err
	}

	origcwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Chdir(localloc)
	if err != nil {
		return err
	}
	defer os.Chdir(origcwd)

	fmt.Println("hg", "pull")
	err = PipedCmd("hg", "pull")
	if err != nil {
		return err
	}

	if dep.Reference == "" {
		dep.Reference = "tip"
	}
	fmt.Println("hg", "update", "-C", dep.Reference)
	err = PipedCmd("hg", "update", "-C", dep.Reference)

	return err

}
