<div id="top" align="center">
    <br>
    <img src="images/nearly.svg" style="height: 80px;"></img>
    <br><br>
    <p>A utility that allows you to toggle system immutability.</p>
    <br>
</div>

> This is still WIP (work in progress).

## Build dependencies

* `golang`

## Building and installing

```sh
make nearly; sudo make install
```

## Authors
`core/chattr.go` is based on https://github.com/snapcore/snapd/blob/master/osutil/chattr.go (I've marked the changes I've made in this file), and is licensed under the GPL-3.0 license.

Everything else in this repository has been generated by `cobra-cli` or written by me, rs2009, and is also licensed under the GPL-3.0 license.

## Usage
```
A utility that allows you to toggle system immutability.

Usage:
  nearly [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  enter       Toggle immutability.
  help        Help about any command
  run         Toggle immutability.
  version     Prints version information.

Flags:
  -h, --help   help for nearly

Use "nearly [command] --help" for more information about a command.
```
