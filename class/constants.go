package class

import "github.com/jvmakine/goasm/classfile"

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
		if ci.Tag == classfile.CONSTANT_Class {
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
		Tag:  classfile.CONSTANT_Class,
		Info: nameIndex,
	})
	return &ClassInfo{c.file, uint16(len(c.file.ConstantPool))}
}
