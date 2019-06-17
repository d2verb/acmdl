# acmdl
CLI for ACM Digital Library

[![asciicast](https://asciinema.org/a/piE9eiJ5JFMuy1KxmfssgyvDj.svg)](https://asciinema.org/a/piE9eiJ5JFMuy1KxmfssgyvDj)

## Installation
### Dependencies
`acmdl` is implemented in Golang and using `dep` to manage dependencies.
If you want build from source code, you have to install Golang to your machine and then install `dep`.
If you don't know how to install Golang, see [here](https://golang.org/doc/install).
If you don't know how to install `dep`, see [here](https://github.com/golang/dep).

### From source
After installing Golang and `dep`. You can use following commands to install `acmdl` to your machine.
```
$ go get github.com/d2verb/acmdl
$ cd $GOPATH/src/github.com/d2verb/acmdl
$ make ensure
$ make install
```

## Usage
### Cmdline options
```
NAME:
   acmdl - ACM Digital Library CLI tool

USAGE:
   exec [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --words value, -w value  words
   --year value, -y value   minimum publication year (default: -1)
   --start value, -s value  minimum numbering (default: 1)
   --help, -h               show help
   --version, -v            print the version
```

## Contributions
Open issues and submit pull requests are welcome.
