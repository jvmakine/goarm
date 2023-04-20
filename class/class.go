package class

import (
	"github.com/jvmakine/goasm/classfile"
)

type Class struct {
	file *classfile.Classfile
}

func NewClass(from *classfile.Classfile) *Class {
	return &Class{from}
}

func (c *Class) AccessFlags() *ClassAccess {
	return &ClassAccess{&c.file.AccessFlags}
}

func (c *Class) ThisClass() *ClassInfo {
	return &ClassInfo{c.file, c.file.ThisClass}
}

func (c *Class) SetThisClass(to *ClassInfo) {
	validateFilesEqual(c.file, to.file)
	c.file.ThisClass = to.index
}

func (c *Class) SuperClass() *ClassInfo {
	return &ClassInfo{c.file, c.file.SuperClass}
}

func (c *Class) SetSuperClass(to *ClassInfo) {
	validateFilesEqual(c.file, to.file)
	c.file.ThisClass = to.index
}

func (c *Class) Interfaces() *Interfaces {
	return &Interfaces{c.file}
}

func (c *Class) Constants() *Constants {
	return &Constants{c.file}
}

func (c *Class) Fields() *Fields {
	return &Fields{c.file}
}

func (c *Class) Methods() *Methods {
	return &Methods{c.file}
}

func (c *Class) Attributes() *Attributes {
	return &Attributes{c.file, c.file.Attributes}
}

func validateFilesEqual(f1, f2 *classfile.Classfile) {
	if f1 != f2 {
		panic("can not combine values from different files")
	}
}

type Interfaces struct {
	file *classfile.Classfile
}

func (c *Interfaces) List() []*ClassInfo {
	result := make([]*ClassInfo, len(c.file.Interfaces))
	for i, ci := range c.file.Interfaces {
		result[i] = &ClassInfo{c.file, ci}
	}
	return result
}

func (c *Interfaces) New(name *String) *ClassInfo {
	validateFilesEqual(c.file, name.file)
	c.file.Interfaces = append(c.file.Interfaces, name.index)
	return &ClassInfo{c.file, name.index}
}
