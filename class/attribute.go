package class

import "github.com/jvmakine/goasm/classfile"

type Attribute struct {
	file      *classfile.Classfile
	attribute *classfile.AttributeInfo
}

func (a *Attribute) Name() *String {
	return &String{a.file, a.attribute.AttributeNameIndex}
}

func (a *Attribute) SetName(to *String) {
	validateFilesEqual(a.file, to.file)
	a.attribute.AttributeNameIndex = to.index
}

func (a *Attribute) Info() []byte {
	return a.Info()
}

func (a *Attribute) SetInfo(to []byte) {
	a.attribute.Info = to
}
