package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/jvmakine/goasm/classfile"
)

var CLI struct {
	ParseCmd
}

type ParseCmd struct {
	ClassFile string `arg:"" help:"the .class file to parse"`
}

func (r *ParseCmd) Run() error {
	file, err := os.Open(r.ClassFile)
	if err != nil {
		return err
	}
	classFile, err := classfile.Parse(file)
	if err != nil {
		return err
	}
	fmt.Printf("version: %d.%d\n", classFile.MajorVersion, classFile.MinorVersion)
	return nil
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
