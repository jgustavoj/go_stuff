# go_stuff

This repo containts packages from other modules:

- [github.com/jgustavoj/go_greetings](https://github.com/jgustavoj/go_greetings)

Some commands to remember

```go
$ go mod init example/hello
// usually you want to use the repository path example: github.com/username/repo

// Create you .go file where you code will live

$ go mod tidy
// Add new module requirements and sums. Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module.

$ go get -u example.com/mypackage
// update package dependencies in your go.mod file

```
