# goat

       _))
      /* \     _~
      `;'\\__-' \_     A simple go dependency manager and project isolator
         | )  _ \ \
        / / ``   w w
       w w

Goat handles recursive, versioned, dependency management for go projects in an
unobtrusive way. Goat also allows projects to be located anywhere on the file
system, unattached to a Go workspace. Best of all, to switch to using goat you
won't have to change a single line of code.

# The problem

There are three problems that goat aims to solve:

* `go get` does not allow for specifying versions of a library.

* `go get` does not have an easy of way of sandboxing your project's development
  environment. You can either muck up your global environment with dependencies
  or mess with your `GOPATH` every time you want to develop for that project.
  Others that want to work on your project will have to do the same.

* `go get` is designed to accomodate a single workspace for all projects.
  This makes it an exercise in frustration to place go projects freely 
  in other areas of ones file system.

# The solution

* The root of a go project is given a `.go` directory. All dependencies are placed
  in this directory. The goat command will automatically look for a `.go` in your
  current working directory or one of its parents, and call that the project root.
  For the rest of the command's duration your GOPATH will have `<projroot>/.go`
  prepended to it.  This has many useful properties, most notably that a project's
  dependencies are sandboxed inside the project directory, but are still usable exactly
  as they would have been if they were global.

* The `.go.yaml` configuration file contains a project's dependency information,
  It is used to specify type, location, pathname, and version control reference
  as applicable for each of a project's dependencies. There is  also a `path` field
  which is used to specify the import path to be used by the current project.

* The `goat` tool is a wrapper around the `go` command line utility. It adds a
  few new commands, all other commands are passed straight through to the normal
  go binary. The main command is `goat deps`, and it retrieves all dependencies
  listed in your `.go.yaml` file and puts them into the `.go` folder. If any
  of those dependencies have a `.go.yaml` file of their own then those are
  also processed and put in your project's `.go` folder (this is done recursively).

# Installation

To use goat you can either get a pre-compiled binary or build it yourself. Once
you get the binary I recommend aliasing it as `go` (`alias go=goat`), so that
`goat` gets used whenever you use the `go` utility. Don't worry, unless
you are in a directory tree with a `.go` directory or use one of goat's special
commands nothing will be different.

## Pre-built

Check the releases tab on github, the latest release will have pre-compiled
binaries for various systems, choose the one that applies to you.

## Build it yourself

To build goat yourself make sure you have `go` installed (go figure).

```bash
git clone https://github.com/mediocregopher/goat.git
cd goat
make
```

The binaries will be found in the `bin` directory.

# Usage

See the [tutorial][tutorial] for a basic use case for goat. After that check out
the [.go.yaml file][projfile] for more details on what kind of features goat has
for dependency management. There are also some [special features][special] that
don't really fit in anywhere else that might be useful to know about.

# Copyrights

Goat ASCII Art (c) 1997 ejm, Creative Commons

[tutorial]: /docs/tut.md
[projfile]: /docs/projfile.md
[special]: /docs/special.md
