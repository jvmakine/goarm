package class

import "github.com/jvmakine/goarm/classfile"

type Class struct {
	file *classfile.Classfile
}

func NewClass(from *classfile.Classfile) *Class {
	return &Class{from}
}

func (c *Class) AccessFlags() map[AccessFlag]bool {
	result := map[AccessFlag]bool{}
	for _, flag := range AllAccessFlags {
		if c.file.AccessFlags&uint16(flag) != 0 {
			result[flag] = true
		}
	}
	return result
}

func (c *Class) SetAccessFlags(flags map[AccessFlag]bool) {
	var result uint16
	for flag := range flags {
		result = result | uint16(flag)
	}
	c.file.AccessFlags = result
}
