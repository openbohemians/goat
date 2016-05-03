# Know what you need? ... It's a *goat*!

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

There are two problems that goat aims to solve:

* `go get` does not allow for specifying versions of a library.

* `go get` does not have an easy of way of sandboxing your project's development
  environment. You can either muck up your global environment with dependencies
  or mess with your `GOPATH` everytime you want to develop for that project.
  Others that want to work on your project will have to do the same.

* Other dependency managers are strange and have weird command line arguments
  that I don't feel like learning.

# The solution

* The root of a goat project has a `.go.yaml` file which specifies a
  dependency's location, name, and version control reference if applicable. It
  is formatted using super-simple yaml objects, each having at most four fields.

* All dependencies are placed in a `.deps` directory at the root of your
  project.  goat will automatically look for a `.go.yaml` in your current
  working directory or one of its parents, and call that the project root. For
  the rest of the command's duration your GOPATH will have `<projroot>/.deps`
  prepended to it. This has many useful properties, most notably that your
  dependencies are sandboxed inside your code, but are still usable exactly as
  they would have been if they were global.

* Goat is a wrapper around the go command line utility. It adds one new command,
  all other commands are passed straight through to the normal go binary. This
  command is `goat deps`, and it retrieves all dependencies listed in your
  `.go.yaml` and puts them into a folder called `vendor` in your project. If any
  of those dependencies have `.go.yaml` files then those are processed and put
  in your project's `vendor` folder as well (this is done recursively).

# Installation

To use goat you can either get a pre-compiled binary or build it yourself. Once
you get the binary I recommend renaming aliasing it as `go` (`alias go=goat`),
so that `goat` gets used whenever you use the `go` utility. Don't worry, unless
you are in a directory tree with a `.go.yaml` file or use one of goat's special
commands nothing will be different.

By default, goat uses the `vendor` directory to store all depenendencies. This
has become the standard location for all Go project as of Go 1.6. However, if
you are the rebellious type, you can change this by adding a `depdir:` entry to
the top-level of your `.go.yaml` file (e.g. `depdir: .deps`). Keep in mind
though, that is a non-standard location and if you or someone else has to
fallback to the original `go` command complications might arise.

## Pre-built

Check the releases tab on github, the latest release will have pre-compiled
binaries for various systems, choose the one that applies to you.

## Build it yourself

To build goat yourself make sure you have `go` installed (go figure).

```bash
git clone https://github.com/openbohemians/goat.git
cd goat
make
```

The binaries will be found in the `bin` directory.

# Usage

See the [tutorial][tutorial] for a basic use case for goat. After that check out
the [.go.yaml file][projfile] for more details on what kind of features goat has
for dependency management. There are also some [special features][special] that
don't really fit in anywhere else that might be useful to know about.

# Authors

* Brian Picciano
* Tom Sawyer

# Copyrights

Goat ASCII Art (c) 1997 ejm, Creative Commons

[tutorial]: /docs/tut.md
[projfile]: /docs/projfile.md
[special]: /docs/special.md
