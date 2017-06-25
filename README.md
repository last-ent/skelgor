# Skelgor

Simple tool to create a project skeleton.

It consists of
```
<Project-Dir>
- Makefile
- main.go
- common/
```

## Makefile
Consists of commands such as:

* `test`
* `build`
* `build.docker`
* `run`
* `run.docker`
* `lint`

## main.go
Consists of a simple "Hello, world."

## common/
Consists of a few helper functions to get started.

* `helpers.go` // Empty file.
* `loggers.go` // Consists of a few simple methods to log errors.
* `test_helpers.go` // Simple formatters to help with Test Pass & Fail formatting.

## How to use: 

```shell

# Install the tool
$ go get github.com/last-ent/skelgor

# Use it to create a project
$ skelgor /path/to/project-dir
```
