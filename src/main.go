// BOF
package main

import (
	"runtime/debug"
	"fmt"
)

var version = "0.0.0"

func main() {
    fmt.Println("Hello, World! [" + version + "]")
	info, _ := debug.ReadBuildInfo()
	fmt.Println(info)
}
// EOF