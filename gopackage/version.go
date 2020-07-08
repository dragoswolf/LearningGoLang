package goversion

import (
	"fmt"
	"runtime"
)

// Version returns the version of golang through a print on screen.
func Version() {
	fmt.Println(runtime.Version())
}

/* 	In order to make a custom go package, we must make the go file with the name of the package as the first line.
After that we make a go mod file with "go mod init <packageName>" and we will use the same package name as we have
in the first line of the source file.

In order to use it, you must import the package in the other source code. If you want to use it from another directory
(that's not cmd) you must specify the whole path to the package itself.
*/
