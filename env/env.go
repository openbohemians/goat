package env

import (
	. "goat/spec"
	"io/ioutil"
	"launchpad.net/goyaml"
	"os"
	"path/filepath"
	"syscall"
)

var PROJFILE = ".go.yaml"

// unmarshal takes in some bytes and tries to decode them into a GoatEnv
// structure
func unmarshal(genvraw []byte) (*GoatEnv, error) {
	genv := GoatEnv{}
	if err := goyaml.Unmarshal(genvraw, &genv); err != nil {
		return nil, err
	}
	return &genv, nil
}

type GoatEnv struct {
	// ProjRoot is the absolute path to the project root in the current
	// environment
	ProjRoot string

	// Path is the path that the project will be using for its own internal
	// import statements, and consequently what other projects depending on this
	// one will use as well.
	Path string `yaml:"path"`

	// DepDir is the directory to set as the GOPATH when fetching dependencies.
	// It is relative to the projRoot
	DepDir string `yaml:"depdir"`

	// Dependencies are the dependencies listed in the project's Goatfile
	Dependencies []Dependency `yaml:"deps"`
}

// NewGoatEnv takes in the directory where a goat project file can be found,
// creates the GoatEnv struct based on that file, and returns it
func NewGoatEnv(projroot string) (*GoatEnv, error) {
	projfile := filepath.Join(projroot, PROJFILE)
	b, err := ioutil.ReadFile(projfile)
	if err != nil {
		return nil, err
	}
	genv, err := unmarshal(b)
	if err != nil {
		return nil, err
	}

	genv.ProjRoot = projroot
	if genv.DepDir == "" {
		genv.DepDir = "vendor" //".deps"
	}
	return genv, nil
}

// AbsDepDir is the absolute path to the directory specified by DepDir
func (genv *GoatEnv) AbsDepDir() string {
	return filepath.Join(genv.ProjRoot, genv.DepDir)
}

// AbsProjFile is the absolute path to the goat project file for this
// environment
func (genv *GoatEnv) AbsProjFile() string {
	return filepath.Join(genv.ProjRoot, PROJFILE)
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// Setup makes sure the goat env has all the proper directories created inside
// of it. This includes the vendor/ directory, and if it's specified the Path
// loopback in the vendor/ directory
func (genv *GoatEnv) Setup() error {
	var err error

	depdir := genv.AbsDepDir()
	srcdir := filepath.Join(depdir, "src")

	// Make vendor directory if it doesn't exist
	if !pathExists(depdir) {
		err = os.Mkdir(depdir, 0755)
		if err != nil {
			return err
		}
  }

  // Make src directory if it doesn't exist
	if !pathExists(srcdir) {
		err = os.Mkdir(srcdir, 0755)
		if err != nil {
			return err
		}
	}

	if genv.Path != "" {
    // TODO: sould the loopback path just be "../.."?
		loopbackPath := filepath.Join(depdir, "src", genv.Path)
		if !pathExists(loopbackPath) {
			loopbackDir := filepath.Dir(loopbackPath)
			if err = os.MkdirAll(loopbackDir, 0755); err != nil {
				return err
			} else if err = os.Symlink(genv.ProjRoot, loopbackPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func (genv *GoatEnv) PrependToGoPath() error {
	gopath, _ := syscall.Getenv("GOPATH")
	return syscall.Setenv("GOPATH", genv.AbsDepDir()+":"+gopath)
}

