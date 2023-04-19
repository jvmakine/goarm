package class

import (
	"github.com/jvmakine/goarm/classfile"
)

type ClassInfo struct {
	file      *classfile.Classfile
	nameIndex uint16
}

func (ci *ClassInfo) Name() *String {
	return &String{ci.file, ci.nameIndex}
}
