package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jvmakine/goasm/class"
	"github.com/jvmakine/goasm/classfile"
)

type DependifyCmd struct {
	ClassFile string `arg:"" help:"the .class file to modify"`
	Classpath string `name:"classpath" required:"" help:"classpath to load dependencies from"`
}

type ClassHashes map[string]string

func (cmd *DependifyCmd) Run() error {
	file, err := os.Open(cmd.ClassFile)
	if err != nil {
		return err
	}
	classFile, err := classfile.Parse(file)
	if err != nil {
		return err
	}
	file.Close()
	clazz := class.NewClass(classFile)

	references := clazz.Constants().ClassInfos()
	hashes := ClassHashes{}
	for _, info := range references {
		name := info.Name().Text()
		if name != clazz.ThisClass().Name().Text() {
			if strings.HasPrefix(name, "java/lang") {
				continue
			}
			hash, err := getHash(name, cmd.Classpath)
			if err != nil {
				return err
			}
			hashes[name] = hash
			println(name + " : " + hash)
		}
	}

	builder := strings.Builder{}
	for name, hash := range hashes {
		builder.WriteString(name)
		builder.WriteString(":")
		builder.WriteString(hash)
		builder.WriteString("\n")
	}
	clazz.Attributes().New(clazz.Constants().NewString("dep_hashes"), []byte(builder.String()))

	file, err = os.OpenFile(cmd.ClassFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := classfile.Write(classFile, file); err != nil {
		return err
	}

	return nil
}

func getHash(name string, dir string) (string, error) {
	f, err := os.Open(filepath.Join(dir, name) + ".class")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return string(hex.EncodeToString(h.Sum(nil))), nil
}
