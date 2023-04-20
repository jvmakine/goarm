package class

import (
	"github.com/jvmakine/goarm/classfile"
)

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

func (c *Class) ThisClass() *ClassInfo {
	return &ClassInfo{c.file, c.file.ThisClass}
}

func (c *Class) SuperClass() *ClassInfo {
	return &ClassInfo{c.file, c.file.SuperClass}
}

func (c *Class) Interfaces() []*ClassInfo {
	result := make([]*ClassInfo, len(c.file.Interfaces))
	for i, ci := range c.file.Interfaces {
		result[i] = &ClassInfo{c.file, ci}
	}
	return result
}

func (c *Class) SetInterfaces(to []*ClassInfo) {
	indices := make([]uint16, len(to))
	for i, ci := range to {
		validateFilesEqual(c.file, ci.file)
		indices[i] = ci.index
	}
	c.file.Interfaces = indices
}

func validateFilesEqual(f1, f2 *classfile.Classfile) {
	if f1 != f2 {
		panic("can not combine values from different files")
	}
}

func (c *Class) Constants() *Constants {
	return &Constants{c.file}
}
