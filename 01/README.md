# It's time to Go!

## Objectives
1. Set up your environment to program in Go
1. Introduce `Modules` and `Packages`
1. Write your first unit test

## Setting up your environment (10 minutes)

1. [Download Go](https://golang.org/dl/)
1. Install the `Go` extension for VSCode (found in Extensions within VSCode)
1. Run `Go: Install/Update Tools` from the [command pallette](https://code.visualstudio.com/docs/getstarted/userinterface#_command-palette)
1. Set your VS Code settings

```yaml
    "go.useLanguageServer": true,
    "go.autocompleteUnimportedPackages": true,
    "go.toolsEnvVars": {
        "GO111MODULE": "on"
    },
    "[go]": {
        "editor.snippetSuggestions": "none",
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true,
        }
    },
    "gopls": {
        "usePlaceholders": true, // add parameter placeholders when completing a function
        "completeUnimported": true, // autocomplete unimported packages
        "watchFileChanges": true,  // watch file changes outside of the editor
        "deepCompletion": true,     // enable deep completion
    }
```

NOTE: These settings are not _required_ to use Go. They are merely recommended, and you are welcome to change them to better suite your preferences.

## Introducing Concepts (10 minutes)

**Package**
A collection of code files that provide _functionality_

- Should be named after the folder that the file is contained in
- Packages can span multiple files
- Should expose functionality, not data

**Module**
A collection of packages that are _versioned together_

- Typically named after the URL where the repository lives
- Uses Git tags that follow SemVer
- Downloaded to your `module cache` which can be found at `$GOPATH/pkg`
- Can use a _proxy_ to store releases (e.g. [proxy.golang.org](https://proxy.golang.org))
- `sum` database stores hash to verify integrity of the module

## Problem (30 minutes)

Write a function that takes a single `string` argument and prefixes the argument with `Hello, `. For example, if the input is `Banana`, the result should be: `Hello, Banana`.

Print the result to the console, and prove your functionality works with a unit test.

### Getting Started

1. Run `go mod init hello` in the root of your project. If you want your code accessible outside of this project, the name of your module should be a URL. For example `github.com/account/hello`.

1. Create a `.go` file (e.g. `hello.go`) and add a `func` called `main()`

1. To build your program run: `go build`. This will create an executable which can then be run.

Alternatively, you can run: `go run hello.go` to run your application without creating a binary.

### Testing

1. The `Go` toolchain expects tests to appended with `_test.go`. In other words, if you have a file called `hello.go`, you should have a test in the same directory named `hello_test.go`.

1. `Go` has a package purpose-built for testing called `testing`. At the top of your test file, be sure to include `import "testing"` at the top of your file.

1. `Go` expects all tests to start with the word `Test`. The name of your test should start with `Test...` and take a single parameter `(t *testing.T)`

1. To report a test failure, you can use `t.Error()` or `t.Errorf()`

1. To run your tests, you can run: `go test`.

## Review and Discussion (10 minutes)

TBD

## Further Reading

1. [Using Go Modules](https://blog.golang.org/using-go-modules) - Go blog series on using modules
1. [Module Mirror](https://blog.golang.org/module-mirror-launch) - Providing module immutibility and integrity
