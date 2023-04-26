package main

import (
	"os"

	"github.com/jvmakine/goasm/class"
	"github.com/jvmakine/goasm/classfile"
)

type DependifyCmd struct {
	ClassFile string `arg:"" help:"the .class file to modify"`
	Classpath string `name:"classpath" required:"" help:"classpath to load dependencies from"`
}

type ClassHashes map[string]string

func (r *DependifyCmd) Run() error {
	file, err := os.Open(r.ClassFile)
	if err != nil {
		return err
	}
	classFile, err := classfile.Parse(file)
	if err != nil {
		return err
	}
	clazz := class.NewClass(classFile)

	references := clazz.Constants().ClassInfos()
	//hashes := ClassHashes{}
	for _, info := range references {
		name := info.Name().Text()
		if name != clazz.ThisClass().Name().Text() {
			println(name)
			// hash, err := getHash(name)
			// if err != nil {
			// 	return err
			// }
			// hashes[name] = hash
		}
	}

	return nil
}

func getHash(name string) (string, error) {
	panic("unimplemented")
}
