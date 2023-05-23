# sysinfop

Lightweight system information package

## Usage

Simply add `sysinfop` as the last line of your shell config (`.bashrc` etc.) to have the program
run when your Terminal Emulator starts. The program has no options.

## Installation

### From Source

Firstly clone the repo

```console
$ git clone https://github.com/joshburnsxyz/sysinfop
```

Then build the program using the included `Makefile`

```console
$ make
```

__OR__ use `go build` if you don't have GNU make available

```console
$ go build -o ./sysinfop
```

Now move the binary onto your path either using `make install` or manually.
