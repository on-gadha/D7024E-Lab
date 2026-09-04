package main

import (
	"github.com/johankristianss/d7024e-tutorial/internal/cli"
	"github.com/johankristianss/d7024e-tutorial/pkg/build"
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
