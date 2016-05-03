package deps

import (
	"fmt"
	"goat/exec"
	. "goat/common"
)

func GoGet(depdir string, dep *Dependency) error {
	fmt.Println("go", "get", dep.Location)
	return exec.PipedCmd("go", "get", dep.Location)
}
