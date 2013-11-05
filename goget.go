package goat

import (
	"fmt"
)

func GoGet(genv *GoatEnv, dep *Dependency) error {
	fmt.Println("go", "get", dep.Location)
	return PipedCmd("go", "get", dep.Location)
}
