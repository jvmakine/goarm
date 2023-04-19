package class

import (
	"fmt"

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

func (c *Class) Info() *ClassInfo {
	index := c.file.ThisClass
	classinfo := c.file.ConstantPool[index-1]
	if classinfo.Tag != classfile.CONSTANT_Class {
		panic(fmt.Errorf("incorrect class info constant type: %d", classinfo.Tag))
	}
	nameIndex := classfile.Order.Uint16(classinfo.Info)
	return &ClassInfo{c.file, nameIndex}
}
