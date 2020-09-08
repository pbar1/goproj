package main

import "github.com/pbar1/goproj/internal/pkg/cli"

var version string

func main() {
	cli.Execute(version)
}
