package class

import (
	"fmt"

	"github.com/jvmakine/goasm/classfile"
)

type ClassInfo struct {
	file  *classfile.Classfile
	index uint16
}

func (ci *ClassInfo) Name() *String {
	classinfo := ci.file.ConstantPool[ci.index-1]
	if classinfo.Tag != classfile.CONSTANT_Class {
		panic(fmt.Errorf("incorrect class info constant type: %d", classinfo.Tag))
	}
	nameIndex := classfile.Order.Uint16(classinfo.Info)
	return &String{ci.file, nameIndex}
}
