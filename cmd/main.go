package main

import (
	"github.com/on-ghada/d7024e-tutorial/internal/cli"
	"github.com/on-ghada/d7024e-tutorial/pkg/build"
)

var (
	BuildVersion string = ""
	BuildTime    string = ""
)

func main() {
	build.BuildVersion = BuildVersion
	build.BuildTime = BuildTime
	cli.Execute()
}
