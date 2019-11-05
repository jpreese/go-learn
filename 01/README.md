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

NOTE: These are not _required_ to use Go, and some are purely preference. These are our recommended settings, but you are welcome to change them to better suite your preferences.

## Introducing Concepts (10 minutes)

**Module**
A collection of packages that are _versioned together_

**Package**
A collection of code files that provide _functionality_

## Problem (30 minutes)

Write a program that prints `Banana` to the console and prove that it works with a unit test.

### Getting Started

1. Run `go mod init banana` in the root of your project. If you want your code accessible outside of this project, the name of your module should be a URL. For example `github.com/account/banana`.

1. Create a `.go` file (e.g. `banana.go`) and add a `func` called `main()`

1. To build your program run: `go build`. This will create an executable which can then be run.

Alternatively, you can run: `go run main.go` to run your application without creating a binary and running it manually.

### Testing

1. The `Go` toolchain expects tests to appended with `_test.go`. In other words, if you have a file called `banana.go`, you should have a test in the same directory named `banana_test.go`.

1. `Go` has a package purpose-built for testing called `testing`. At the top of your test file, be sure to `import "testing"`

1. `Go` expects all tests to start with the word `Test`. The name of your test should start with `Test...` and take a single parameter `(t *testing.T)`

1. To report a test failure, you can use `t.Error()` or `t.Errorf()`

1. To run your tests, you can run: `go test`.

## Review and Takeaways (10 minutes)

TBD

