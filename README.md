# Peanut Compiler

This is our compiler project for CPSC 323-03 of Fall 2023 with Choi. By the end of the semester, this should be able to compile Rat23F programs.

It's called the Peanut Compiler because both rats and gophers can eat peanuts.

## Current state

Assignment 2 - create a syntax analyzer

## Quick start development

### Windows

1. Download and install Go from https://go.dev/dl/
1. Clone the repository
1. Edit code (the _.go_ files) in src\
1. Use the command `go run .` to run/debug the project.
    - You will likely need to provide a file as an argument, for example `go run . "tests\test.rat"`

### Building an exe

1. In the command prompt, navigate to the _src_ directory.
1. Run `go build .`
1. Observe the creation of _peanut-compiler.exe_

## Resources

Learning Go: https://go.dev/tour/list

Lexical Scanning in Go: https://go.dev/talks/2011/lex.slide#1
