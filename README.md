# acmdl
CLI for ACM Digital Library

[![asciicast](https://asciinema.org/a/3R7o6iqnEKisi6FLnCusXKbbR.svg)](https://asciinema.org/a/3R7o6iqnEKisi6FLnCusXKbbR)

## Installation
### Dependencies
`acmdl` is implemented in Golang and using `dep` to manage dependencies.
If you want build from source code, you have to install Golang to your machine and then install `dep`.
If you don't know how to install Golang, see [here](https://golang.org/doc/install).
If you don't know how to install `dep`, see [here](https://github.com/golang/dep).

After installing Golang and `dep`. Following commands install `acmdl` to your machine.
```
$ git clone https://github.com/d2verb/acmdl
$ cd acmdl
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
Open issues and pull requests are welcome.
