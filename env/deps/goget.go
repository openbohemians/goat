package deps

import (
	. "goat/spec"
	"goat/exec"
	"fmt"
)

func GoGet(depdir string, dep *Dependency) error {
	fmt.Println("go", "get", dep.Location)
	return exec.PipedCmd("go", "get", dep.Location)
}
