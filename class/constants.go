package class

import (
	"fmt"

	"github.com/jvmakine/goasm/classfile"
)

type Constants struct {
	file *classfile.Classfile
}

func (c *Constants) Strings() []*String {
	var result []*String
	for i, ci := range c.file.ConstantPool {
		if ci.Tag == classfile.CONSTANT_Utf8 {
			result = append(result, &String{c.file, uint16(i + 1)})
		}
	}
	return result
}

func (c *Constants) NewString(text string) *String {
	textBytes := []byte(text)
	lengthBytes := make([]byte, 2)
	classfile.Order.PutUint16(lengthBytes, uint16(len(textBytes)))
	c.file.ConstantPool = append(c.file.ConstantPool, &classfile.ConstantInfo{
		Tag:  classfile.CONSTANT_Utf8,
		Info: append(lengthBytes, textBytes...),
	})
	return &String{c.file, uint16(len(c.file.ConstantPool))}
}

func (c *Constants) ClassInfos() []*ClassInfo {
	var result []*ClassInfo
	for i, ci := range c.file.ConstantPool {
		if ci.Tag == uint8(classfile.CONSTANT_Class) {
			result = append(result, &ClassInfo{c.file, uint16(i + 1)})
		}
	}
	return result
}

func (c *Constants) NewClassInfo(name *String) *ClassInfo {
	validateFilesEqual(name.file, c.file)

	nameIndex := make([]byte, 2)
	classfile.Order.PutUint16(nameIndex, name.index)
	c.file.ConstantPool = append(c.file.ConstantPool, &classfile.ConstantInfo{
		Tag:  uint8(classfile.CONSTANT_Class),
		Info: nameIndex,
	})
	return &ClassInfo{c.file, uint16(len(c.file.ConstantPool))}
}

type ClassInfo struct {
	file  *classfile.Classfile
	index uint16
}

func (ci *ClassInfo) Name() *String {
	classinfo := ci.file.ConstantPool[ci.index-1]
	if classinfo.Tag != uint8(classfile.CONSTANT_Class) {
		panic(fmt.Errorf("incorrect class info constant type: %d", classinfo.Tag))
	}
	nameIndex := classfile.Order.Uint16(classinfo.Info)
	return &String{ci.file, nameIndex}
}

type String struct {
	file  *classfile.Classfile
	index uint16
}

func (s *String) Text() string {
	entry := s.file.ConstantPool[s.index-1]
	if entry.Tag != classfile.CONSTANT_Utf8 {
		panic(fmt.Errorf("incorrect string constant type: %d", entry.Tag))
	}
	length := classfile.Order.Uint16(entry.Info)
	if len(entry.Info) != int(length)+2 {
		panic(fmt.Errorf("incorrect string length: %d, expected %d", len(entry.Info)-2, length))
	}
	return string(entry.Info[2:])
}

func (s *String) SetText(to string) {
	entry := s.file.ConstantPool[s.index-1]
	if entry.Tag != classfile.CONSTANT_Utf8 {
		panic(fmt.Errorf("incorrect string constant type: %d", entry.Tag))
	}
	bytes := []byte(to)

	lengthBytes := make([]byte, 2)
	classfile.Order.PutUint16(lengthBytes, uint16(len(bytes)))

	entry.Info = append(lengthBytes, bytes...)
}
