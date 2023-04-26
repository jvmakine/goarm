package main

import (
	"github.com/alecthomas/kong"
)

var CLI struct {
	Cat       CatCmd       `cmd:""  help:"Print out details of a classfile"`
	Dependify DependifyCmd `cmd:""  help:"Save the dependency table to a classfile"`
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
